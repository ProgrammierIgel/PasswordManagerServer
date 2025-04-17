package changeurl

type RequestBody struct {
	PasswordName   string `json:"passwordName"`
	NewUsername    string `json:"newUsername"`
	AccountName    string `json:"accountName"`
	MasterPassword string `json:"accountPassword"`
}
