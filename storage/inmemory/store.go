package inmemory

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/programmierigel/pwmanager/cryptography"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/programmierigel/pwmanager/manager"
	"github.com/programmierigel/pwmanager/tools"
)

type Store struct {
	file                 string
	decryptionPasswords  map[string]manager.Password
	secrets              map[string]map[string]manager.Secret
	syncDisabled         bool
	password             string
	token                map[string]manager.Token
	remoteRequestAddress map[string]map[string]manager.RemoteRequestAddress
}

func New(path string, password string) *Store {

	store := &Store{
		file:                 fmt.Sprintf("%s/secrets.json", path),
		decryptionPasswords:  make(map[string]manager.Password),
		secrets:              make(map[string]map[string]manager.Secret),
		syncDisabled:         false,
		password:             password,
		token:                make(map[string]manager.Token),
		remoteRequestAddress: make(map[string]map[string]manager.RemoteRequestAddress),
	}
	store.SyncFromFile()
	logger.Info("[STORE] New Store was created")
	return store
}

func (s *Store) SyncFromFile() error {
	if s.syncDisabled {
		logger.Critiacal("[STORE] syncronization is disabled")
		return nil
	}

	jsonFile, err := os.Open(s.file)

	if err != nil {
		logger.Critiacal("[STORE] Secrets file was not found")
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var secretsFile manager.SecretsFile

	err = json.Unmarshal(byteValue, &secretsFile)
	if err != nil {
		logger.Critiacal("[STORE] Content of secrets file is not compatible to JSON")
		return err
	}
	s.decryptionPasswords = secretsFile.MainPasswords
	s.secrets = secretsFile.Secrets
	logger.Info("[STORE] System load save from file")
	return nil
}

func (s *Store) SyncToFile() error {
	if s.syncDisabled {
		logger.Critiacal("[STORE] syncronization is disabled")
		return nil
	}

	file := manager.SecretsFile{
		MainPasswords: s.decryptionPasswords,
		Secrets:       s.secrets,
	}

	parsedFile, err := json.Marshal(file)
	if err != nil {
		logger.Critiacal("[STORE] Content of secrets file is not compatible to JSON")
		return err
	}

	err = os.WriteFile(s.file, parsedFile, fs.FileMode(0222))
	if err != nil {
		logger.Critiacal("[STORE] Content cant saved to secrets file")
		return err
	}
	logger.Info("[STORE] System saved to file")
	return nil

}

func (s *Store) AddNewAccount(account string, password string) error {
	salt, err := cryptography.GenerateSalt(15)
	if err != nil {
		return err
	}
	s.SyncFromFile()
	if s.decryptionPasswords[account].PasswordHash != "" {
		logger.Warning(fmt.Sprintf("[STORE] Attempt to create a account that already exists (account name: %s)", account))
		return fmt.Errorf("account already exists")
	}
	updatePasswordStruct := manager.Password{
		PasswordHash: cryptography.EncryptSHA256(password + salt),
		Salt:         salt,
	}

	s.decryptionPasswords[account] = updatePasswordStruct
	logger.Info(fmt.Sprintf("[STORE] New Account was created with name %s", account))
	s.SyncToFile()
	return nil
}

func (s *Store) DeleteAccount(account string, token string) error {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to delete account %s account. Failed due to incorrect token.", account))
		return fmt.Errorf("incorrect token")
	}
	if s.token[token].AccountName != account {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to delete account %s account. Failed due trying to delete other account", account))
		return fmt.Errorf("cant delete other account")
	}
	s.decryptionPasswords = tools.RemovePasswordFromMap(s.decryptionPasswords, account)
	s.secrets = tools.RemoveMapFromMap(s.secrets, account)
	logger.Warning(fmt.Sprintf("[STORE] Account %s was deleted", account))
	s.SyncToFile()
	return nil
}

func (s *Store) AddNewPassword(token string, passwordName string, passwordToAdd string, url string, username string) error {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to add a password (%s). Failed due to incorrect password.", passwordName))
		return fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]

	if tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to add a password (%s) on account %s but account already exists.", passwordName, tokenValue.AccountName))
		return fmt.Errorf("already exists")
	}
	hash, err := cryptography.Encrypt(passwordToAdd, tokenValue.MasterPassword)
	if err != nil {
		logger.Critiacal(fmt.Sprintf("[STORE] Cant hash password: %s", err.Error()))
		return err
	}

	if s.secrets[tokenValue.AccountName] == nil {
		s.secrets[tokenValue.AccountName] = make(map[string]manager.Secret)
	}

	secretsObject := manager.Secret{
		Secret:   hash,
		URL:      url,
		Username: username,
	}

	s.secrets[tokenValue.AccountName][passwordName] = secretsObject
	logger.Info(fmt.Sprintf("[STORE] New password (%s) added on account %s", passwordName, tokenValue.AccountName))
	s.SyncToFile()
	return nil
}

func (s *Store) DeletePassword(token string, passwordName string) error {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to delete password (%s) on account. Failed due to incorrect token.", passwordName))
		return fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]

	if !tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning("[Store] Attempt to deleting non-existing password")
		return fmt.Errorf("passwordname does not exist")
	}

	s.secrets[tokenValue.AccountName] = tools.RemoveStringFromMap(s.secrets[tokenValue.AccountName], passwordName)
	logger.Warning(fmt.Sprintf("[STORE] Deleted %s password on account %s", passwordName, tokenValue.AccountName))
	s.SyncToFile()
	return nil
}

func (s *Store) GetPassword(token string, passwordName string) (string, error) {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to get password %s from account. Failed due to incorrect token to decryption.", passwordName))
		return "", fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]
	if !tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to get password %s from account %s but password doesn't exists on account", passwordName, tokenValue.AccountName))
		return "", fmt.Errorf("password on account not found")
	}
	defer logger.Info(fmt.Sprintf("[STORE] Password to %s on account %s successfully returned", passwordName, tokenValue.AccountName))

	password, err := cryptography.Decrypt(s.secrets[tokenValue.AccountName][passwordName].Secret, tokenValue.MasterPassword)
	if err != nil {
		logger.Critiacal(fmt.Sprintf("[STORE] Cant return password: %s", err.Error()))
		return "", err
	}
	return password, nil
}

func (s *Store) GetURL(token string, passwordName string) (string, error) {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to get url %s from account. Failed due to incorrect token.", passwordName))
		return "", fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]

	if !tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to get url %s from account %s but url doesn't exists on account", passwordName, tokenValue.AccountName))
		return "", fmt.Errorf("url on account not found")
	}
	defer logger.Info(fmt.Sprintf("[STORE] Url to %s on account %s successfully returned", passwordName, tokenValue.AccountName))

	url := s.secrets[tokenValue.AccountName][passwordName].URL
	return url, nil
}

func (s *Store) GetUsername(token string, passwordName string) (string, error) {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to get username %s from account. Failed due to incorrect token.", passwordName))
		return "", fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]

	if !tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to get username %s from account %s but url doesn't exists on account", s.secrets[tokenValue.AccountName][passwordName].Username, tokenValue.AccountName))
		return "", fmt.Errorf("url on account not found")
	}
	defer logger.Info(fmt.Sprintf("[STORE] Username to %s on account %s successfully returned", passwordName, tokenValue.AccountName))

	username := s.secrets[tokenValue.AccountName][passwordName].Username
	return username, nil
}

func (s *Store) GetAllPasswordNamesOfAccount(token string) ([]string, error) {
	if !s.CheckToken(token) {
		logger.Warning("[STORE] Attemt to get all registered passwords from account. Failed due to incorrect token to decryption.")
		return make([]string, 0), fmt.Errorf("incorrect token")
	}

	tokenValue := s.token[token]
	allPasswordNames := make([]string, 0)
	for name := range s.secrets[tokenValue.AccountName] {
		allPasswordNames = append(allPasswordNames, name)
	}
	logger.Info(fmt.Sprintf("[STORE] All PasswordNames of Account %s returned", tokenValue.AccountName))
	return allPasswordNames, nil

}

func (s *Store) DisableSync(password string) (bool, error) {
	if s.password != password {
		return s.syncDisabled, fmt.Errorf("wrong password")
	}
	s.syncDisabled = true
	logger.Critiacal("[STORE] !SYNC IS DISABLED!")
	return s.syncDisabled, nil
}

func (s *Store) EnableSync(password string) (bool, error) {
	if s.password != password {
		return s.syncDisabled, fmt.Errorf("wrong password")
	}
	s.syncDisabled = false
	logger.Info("[STORE] SYNC IS ENABLED")
	return s.syncDisabled, nil
}

func (s *Store) IsSyncDisabled() bool {
	return s.syncDisabled
}

func (s *Store) ChangeUsername(token string, passwordName string, newUsername string) error {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change username to %s. Failed due to incorrect token.", newUsername))
		return fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]

	if !tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change username %s from account %s but username doesn't exists on account", passwordName, tokenValue.AccountName))
		return fmt.Errorf("username on account not found")
	}
	defer logger.Info(fmt.Sprintf("[STORE] URL to %s on account %s successfully changed", passwordName, tokenValue.AccountName))

	currentPasswordStruct := s.secrets[tokenValue.AccountName][passwordName]
	s.secrets[tokenValue.AccountName][passwordName] = manager.Secret{
		URL:      currentPasswordStruct.URL,
		Secret:   currentPasswordStruct.Secret,
		Username: newUsername,
	}
	return nil
}

func (s *Store) ChangeURL(token string, passwordName string, newURL string) error {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change URL to %s. Failed due to incorrect password.", newURL))
		return fmt.Errorf("incorrect token")
	}

	tokenValue := s.token[token]
	if !tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change url %s from account %s but url doesn't exists on account", passwordName, tokenValue.AccountName))
		return fmt.Errorf("username on account not found")
	}
	defer logger.Info(fmt.Sprintf("[STORE] URL to %s on account %s successfully changed", passwordName, tokenValue.AccountName))

	currentPasswordStruct := s.secrets[tokenValue.AccountName][passwordName]
	s.secrets[tokenValue.AccountName][passwordName] = manager.Secret{
		URL:      newURL,
		Secret:   currentPasswordStruct.Secret,
		Username: currentPasswordStruct.Username,
	}
	return nil
}

func (s *Store) ChangePassword(token string, passwordName string, newSecret string) error {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change password to %s. Failed due to incorrect token.", newSecret))
		return fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]

	if !tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change password %s from account %s but url doesn't exists on account", passwordName, tokenValue.AccountName))
		return fmt.Errorf("username on account not found")
	}
	defer logger.Info(fmt.Sprintf("[STORE] URL to %s on account %s successfully changed", passwordName, tokenValue.AccountName))

	hash, err := cryptography.Encrypt(newSecret, tokenValue.MasterPassword)
	if err != nil {
		logger.Critiacal(fmt.Sprintf("[STORE] Cant hash password: %s", err.Error()))
		return err
	}

	currentPasswordStruct := s.secrets[tokenValue.AccountName][passwordName]
	s.secrets[tokenValue.AccountName][passwordName] = manager.Secret{
		URL:      currentPasswordStruct.URL,
		Secret:   hash,
		Username: currentPasswordStruct.Username,
	}
	return nil
}

func (s *Store) ChangePasswordName(token string, passwordName string, newPasswordName string) error {
	if !s.CheckToken(token) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change passwordname %s to %s. Failed due to incorrect token.", passwordName, newPasswordName))
		return fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]

	if tools.IsElementInMap(newPasswordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change passwordname %s from account %s but id already exists", passwordName, tokenValue.AccountName))
		return fmt.Errorf("new passwordname already exists (overwrite)")
	}

	if !tools.IsElementInMap(passwordName, s.secrets[tokenValue.AccountName]) {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to change passwordname %s from account %s but but passwordname doesn't exists on account", passwordName, tokenValue.AccountName))
		return fmt.Errorf("passwordname on account not found")
	}
	defer logger.Info(fmt.Sprintf("[STORE] Passwordname to %s on account %s successfully changed", passwordName, tokenValue.AccountName))

	s.secrets[tokenValue.AccountName][passwordName] = s.secrets[tokenValue.AccountName][newPasswordName]

	s.secrets[tokenValue.AccountName] = tools.RemoveStringFromMap(s.secrets[tokenValue.AccountName], newPasswordName)

	return nil
}

func (s *Store) CreateToken(accountName string, masterpassword string, remoteAddress string) (string, error) {
	currentTime := time.Now()

	err := s.CheckPassword(accountName, masterpassword, remoteAddress)
	if err != nil {
		logger.Warning(fmt.Sprintf("[STORE] Attemt to create token on account %s. Failed due to incorrect password.", accountName))
		return "", err
	}
	newID := uuid.New().String()
	if tools.IsIDInMap(newID, s.token) {
		return "", fmt.Errorf("internal error: cant create new Tokens")
	}

	s.token[newID] = manager.Token{
		Timestamp:      currentTime,
		AccountName:    accountName,
		MasterPassword: masterpassword,
	}
	logger.Info(fmt.Sprintf("[STORE] Successfully token on account %s created.", accountName))
	return newID, nil

}

func (s *Store) CheckToken(token string) bool {
	currentTime := time.Now()
	if !tools.IsIDInMap(token, s.token) {
		logger.Info("[STORE] token to check not registered")
		return false
	}
	oneHourInMillisec := int64(3600000)
	if currentTime.UnixMilli()-s.token[token].Timestamp.UnixMilli() > oneHourInMillisec {
		logger.Info("[STORE] time of token to check is over")
		s.token = tools.RemoveTokenFromMap(token, s.token)
		return false
	}
	return true
}

func (s *Store) DevalueToken(token string) {
	if !s.CheckToken(token) {
		return
	}

	defer logger.Info((fmt.Sprintf("[STORE] Successfully token from account %s devalued", s.token[token].AccountName)))
	s.token = tools.RemoveTokenFromMap(token, s.token)
}

func (s *Store) CheckPassword(account string, password string, remoteAddress string) error {
	currentTime := time.Now()
	s.SyncFromFile()
	if s.remoteRequestAddress[account] == nil {
		s.remoteRequestAddress[account] = make(map[string]manager.RemoteRequestAddress)
	}
	accessess := s.remoteRequestAddress[account][remoteAddress].NumberOfAccessess
	timestamp := s.remoteRequestAddress[account][remoteAddress].LastRequest

	switch {
	case accessess > 3:
		if currentTime.UnixMilli()-timestamp.UnixMilli() < time.Minute.Milliseconds()*5 {
			logger.Warning("[STORE] Too many tries to check password! Must wait!")
			return fmt.Errorf("too many tries! please wait")
		}
	case accessess > 6:
		if currentTime.UnixMilli()-timestamp.UnixMilli() < time.Hour.Milliseconds() {
			logger.Warning("[STORE] Too many tries to check password! Must wait!")
			return fmt.Errorf("too many tries! please wait")
		}
	case accessess > 9:
		if currentTime.UnixMilli()-timestamp.UnixMilli() < time.Hour.Milliseconds()*24 {
			logger.Warning("[STORE] Too many tries to check password! Must wait!")
			return fmt.Errorf("too many tries! please wait")
		}

	}

	salt := s.decryptionPasswords[account].Salt
	if s.decryptionPasswords[account].PasswordHash != cryptography.EncryptSHA256(password+salt) {
		s.remoteRequestAddress[account][remoteAddress] = manager.RemoteRequestAddress{
			LastRequest:       currentTime,
			NumberOfAccessess: accessess + 1,
		}

		return fmt.Errorf("unknown password")
	}

	remainingRequests := make(map[string]manager.RemoteRequestAddress)
	for request, requestStruct := range s.remoteRequestAddress[account] {
		if remoteAddress == request {
			continue
		}
		remainingRequests[request] = requestStruct
	}

	s.remoteRequestAddress[account] = remainingRequests
	return nil
}

func (s *Store) DevalueAllTokens(password string) error {
	if password != s.password {
		logger.Warning("[STORE] Attempt to devalue all tokens. Failed due incorrect password")
		return fmt.Errorf("incorrect password")
	}
	// Path and Password are not needed to create new token object
	newStore := New("", "")
	s.token = newStore.token
	logger.Warning("[STORE] All tokens succesfully devalued!")
	return nil
}

func (s *Store) DevalueAllTokensOfAccount(token string) error {
	if !s.CheckToken(token) {
		logger.Warning("[STORE] Attempt to devalue all tokens of account. Token incorrect")
		return fmt.Errorf("token incorrect")
	}

	accountName := s.token[token].AccountName
	remainingTokens := make(map[string]manager.Token)
	for tokenIterator, tokenValue := range s.token {
		if tokenValue.AccountName == accountName {
			continue
		}
		remainingTokens[tokenIterator] = tokenValue
	}
	s.token = remainingTokens
	logger.Info(fmt.Sprintf("[STORE] Successfully all tokens of account %s devalued", accountName))
	return nil
}

func (s *Store) GetAllActiveTokens(token string) (uint, error) {
	if !s.CheckToken(token) {
		logger.Warning("[STORE] Attemt to get the number of all active tokens %s. Failed due to incorrect token.")
		return 0, fmt.Errorf("incorrect token")
	}
	tokenValue := s.token[token]
	numberOfRegisteredTokens := uint(0)
	for t, v := range s.token {
		if v.AccountName == tokenValue.AccountName {
			if s.CheckToken(t) {
				numberOfRegisteredTokens += 1
			}
		}
		continue
	}
	logger.Info("[STORE] Successuly number of all active tokens returned")
	return numberOfRegisteredTokens, nil
}
