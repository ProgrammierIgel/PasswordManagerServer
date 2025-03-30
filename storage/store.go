package storage

type Store interface {
	AddNewAccount(name string, password string) error
	DeleteAccount(name string, password string) error
	CheckPassword(account string, password string) error
	AddPassword(masterPassword string, account string, passwordName string, passwordToAdd string) error
	DeletePassword(masterPassword string, account string, passwordName string) error
	GetPassword(account string, masterPassword string, passwordName string) (string, error)
	GetAllPasswordNamesOfAccount(account string, masterPassword string) ([]string, error)
	SyncFromFile() error
	SyncToFile() error
	EnableSync(password string) (bool, error)
	DisableSync(password string) (bool, error)
	IsSyncDisabled() bool
}
