package storage

type Store interface {
	AddNewAccount(name string, password string) error
	DeleteAccount(name string, password string) error
	CheckPassword(account string, password string) error
	AddNewPassword(masterPassword string, account string, passwordName string, passwordToAdd string, url string, username string) error
	DeletePassword(masterPassword string, account string, passwordName string) error
	GetPassword(account string, masterPassword string, passwordName string) (string, error)
	GetUsername(account string, masterPassword string, passwordName string) (string, error)
	GetAllPasswordNamesOfAccount(account string, masterPassword string) ([]string, error)
	SyncFromFile() error
	SyncToFile() error
	EnableSync(password string) (bool, error)
	DisableSync(password string) (bool, error)
	IsSyncDisabled() bool
	GetURL(account string, masterPassword string, passwordName string) (string, error)
	ChangeUsername(account string, masterPassword string, passwordName string, newUsername string) error
	ChangeURL(account string, masterPassword string, passwordName string, newURL string) error
	ChangePassword(account string, masterPassword string, passwordName string, newSecret string) error
	ChangePasswordName(account string, masterPassword string, passwordName string, newPasswordName string) error
}
