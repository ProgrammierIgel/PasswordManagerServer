package storage

type Store interface {
	AddNewAccount(name string, password string) error
	DeleteAccount(name string, token string) error
	CheckPassword(account string, password string) error
	AddNewPassword(token string, passwordName string, passwordToAdd string, url string, username string) error
	DeletePassword(token string, passwordName string) error
	GetPassword(token string, passwordName string) (string, error)
	GetUsername(token string, passwordName string) (string, error)
	GetAllPasswordNamesOfAccount(token string) ([]string, error)
	SyncFromFile() error
	SyncToFile() error
	EnableSync(password string) (bool, error)
	DisableSync(password string) (bool, error)
	IsSyncDisabled() bool
	GetURL(token string, passwordName string) (string, error)
	ChangeUsername(token string, passwordName string, newUsername string) error
	ChangeURL(token string, passwordName string, newURL string) error
	ChangePassword(token string, passwordName string, newSecret string) error
	ChangePasswordName(token string, passwordName string, newPasswordName string) error
	DevalueToken(token string)
	CheckToken(token string) bool
	CreateToken(accountName string, masterpassword string) (string, error)
	DevalueAllTokens(password string) error
	DevalueAllTokensOfAccount(token string) error
}
