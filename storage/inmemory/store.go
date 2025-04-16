package inmemory

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/programmierigel/pwmanager/cryptography"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/programmierigel/pwmanager/manager"
	"github.com/programmierigel/pwmanager/tools"
)

type Store struct {
	file                string
	decryptionPasswords map[string]manager.Password
	secrets             map[string]map[string]manager.Secret
	syncDisabled        bool
	password            string
}

func New(path string, password string) *Store {

	store := &Store{
		file:                path + "/secrets.json",
		decryptionPasswords: make(map[string]manager.Password),
		secrets:             make(map[string]map[string]manager.Secret),
		syncDisabled:        false,
		password:            password,
	}
	store.SyncFromFile()
	logger.Info("New Store was created")
	return store
}

func (s *Store) SyncFromFile() error {
	if s.syncDisabled {
		logger.Critiacal("syncronization is disabled")
		return nil
	}

	jsonFile, err := os.Open(s.file)

	if err != nil {
		logger.Critiacal("Secrets file was not found")
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var secretsFile manager.SecretsFile

	err = json.Unmarshal(byteValue, &secretsFile)
	if err != nil {
		logger.Critiacal("Content of secrets file is not compatible to JSON")
		return err
	}
	s.decryptionPasswords = secretsFile.MainPasswords
	s.secrets = secretsFile.Secrets
	logger.Info("System load save from file")
	return nil
}

func (s *Store) SyncToFile() error {
	if s.syncDisabled {
		logger.Critiacal("syncronization is disabled")
		return nil
	}

	file := manager.SecretsFile{
		MainPasswords: s.decryptionPasswords,
		Secrets:       s.secrets,
	}

	parsedFile, err := json.Marshal(file)
	if err != nil {
		logger.Critiacal("Content of secrets file is not compatible to JSON")
		return err
	}

	err = os.WriteFile(s.file, parsedFile, fs.FileMode(0222))
	if err != nil {
		logger.Critiacal("Content cant saved to secrets file")
		return err
	}
	logger.Info("System saved to file")
	return nil

}

func (s *Store) AddNewAccount(account string, password string) error {
	salt, err := cryptography.GenerateSalt(15)
	if err != nil {
		return err
	}
	s.SyncFromFile()
	if s.decryptionPasswords[account].PasswordHash != "" {
		logger.Warning(fmt.Sprintf("Attempt to create a account that already exists (account name: %s)", account))
		return fmt.Errorf("account already exists")
	}
	updatePasswordStruct := manager.Password{
		PasswordHash: cryptography.EncryptSHA256(password + salt),
		Salt:         salt,
	}

	s.decryptionPasswords[account] = updatePasswordStruct
	logger.Info(fmt.Sprintf("New Account was created with name %s", account))
	s.SyncToFile()
	return nil
}

func (s *Store) DeleteAccount(account string, password string) error {
	err := s.CheckPassword(account, password)
	if err != nil {
		logger.Warning(fmt.Sprintf("Attemt to delete account %s account. Failed due to incorrect password.", account))
		return err
	}
	s.decryptionPasswords = tools.RemovePasswordFromMap(s.decryptionPasswords, account)
	s.secrets = tools.RemoveMapFromMap(s.secrets, account)
	logger.Warning(fmt.Sprintf("Account %s was deleted", account))
	s.SyncToFile()
	return nil
}

func (s *Store) CheckPassword(account string, password string) error {
	s.SyncFromFile()
	salt := s.decryptionPasswords[account].Salt
	if s.decryptionPasswords[account].PasswordHash != cryptography.EncryptSHA256(password+salt) {
		return fmt.Errorf("unknown password")
	}
	return nil
}

func (s *Store) AddNewPassword(masterPassword string, account string, passwordName string, passwordToAdd string, url string) error {
	err := s.CheckPassword(account, masterPassword)
	if err != nil {
		logger.Warning(fmt.Sprintf("Attemt to add a password (%s) on account %s. Failed due to incorrect password.", passwordName, account))
		return err
	}

	if tools.IsElementInMap(passwordName, s.secrets[account]) {
		logger.Warning(fmt.Sprintf("Attemt to add a password (%s) on account %s but account already exists.", passwordName, account))
		return fmt.Errorf("already exists")
	}
	hash, err := cryptography.Encrypt(passwordToAdd, masterPassword)
	if err != nil {
		logger.Critiacal(fmt.Sprintf("Cant hash password: %s", err.Error()))
		return err
	}

	if s.secrets[account] == nil {
		s.secrets[account] = make(map[string]manager.Secret)
	}

	secretsObject := manager.Secret{
		Secret: hash,
		URL:    url,
	}

	s.secrets[account][passwordName] = secretsObject
	logger.Info(fmt.Sprintf("New password (%s) added on account %s", passwordName, account))
	s.SyncToFile()
	return nil
}

func (s *Store) DeletePassword(masterPassword string, account string, passwordName string) error {
	err := s.CheckPassword(account, masterPassword)
	if err != nil {
		logger.Warning(fmt.Sprintf("Attemt to delete password (%s) on account %s . Failed due to incorrect password.", passwordName, account))
		return err
	}

	s.secrets[account] = tools.RemoveStringFromMap(s.secrets[account], passwordName)
	logger.Warning(fmt.Sprintf("Deleted %s password on account %s", passwordName, account))
	s.SyncToFile()
	return nil
}

func (s *Store) GetPassword(account string, masterPassword string, passwordName string) (string, error) {
	err := s.CheckPassword(account, masterPassword)
	if err != nil {
		logger.Warning(fmt.Sprintf("Attemt to get password %s from account %s. Failed due to incorrect password to decryption.", passwordName, account))
		return "", err
	}

	if !tools.IsElementInMap(passwordName, s.secrets[account]) {
		logger.Warning(fmt.Sprintf("Attemt to get password %s from account %s but password doesn't exists on account", passwordName, account))
		return "", fmt.Errorf("password on account not found")
	}
	defer logger.Info(fmt.Sprintf("Password to %s on account %s successfully returned", passwordName, account))

	password, err := cryptography.Decrypt(s.secrets[account][passwordName].Secret, masterPassword)
	if err != nil {
		logger.Critiacal(fmt.Sprintf("Cant return password: %s", err.Error()))
		return "", err
	}
	return password, nil
}

func (s *Store) GetURL(account string, masterPassword string, passwordName string) (string, error) {
	err := s.CheckPassword(account, masterPassword)
	if err != nil {
		logger.Warning(fmt.Sprintf("Attemt to get url %s from account %s. Failed due to incorrect password.", passwordName, account))
		return "", err
	}

	if tools.IsElementInMap(passwordName, s.secrets[account]) {
		logger.Warning(fmt.Sprintf("Attemt to get url %s from account %s but url doesn't exists on account", passwordName, account))
		return "", fmt.Errorf("url on account not found")
	}
	defer logger.Info(fmt.Sprintf("Url to %s on account %s successfully returned", passwordName, account))

	url := s.secrets[account][passwordName].Secret
	return url, nil
}

func (s *Store) GetAllPasswordNamesOfAccount(account string, masterPassword string) ([]string, error) {
	err := s.CheckPassword(account, masterPassword)
	if err != nil {
		logger.Warning(fmt.Sprintf("Attemt to get all registered passwords from account %s. Failed due to incorrect password to decryption.", account))
		return make([]string, 0), err
	}
	allPasswordNames := make([]string, 0)
	for name := range s.secrets[account] {
		allPasswordNames = append(allPasswordNames, name)
	}
	logger.Debug(fmt.Sprintf("All PasswordNames of Account %s returned", account))
	return allPasswordNames, nil

}

func (s *Store) DisableSync(password string) (bool, error) {
	if s.password != password {
		return s.syncDisabled, fmt.Errorf("wrong password")
	}
	s.syncDisabled = true
	logger.Critiacal("!SYNC IS DISABLED!")
	return s.syncDisabled, nil
}

func (s *Store) EnableSync(password string) (bool, error) {
	if s.password != password {
		return s.syncDisabled, fmt.Errorf("wrong password")
	}
	s.syncDisabled = false
	logger.Info("SYNC IS ENABLED")
	return s.syncDisabled, nil
}

func (s *Store) IsSyncDisabled() bool {
	return s.syncDisabled
}
