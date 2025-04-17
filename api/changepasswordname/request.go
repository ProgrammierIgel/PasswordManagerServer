package changepasswordname

type RequestBody struct {
	PasswordName    string `json:"passwordName"`
	NewPasswordName string `json:"newPasswordName"`
	AccountName     string `json:"accountName"`
	MasterPassword  string `json:"accountPassword"`
}
