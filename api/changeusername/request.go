package changeusername

type RequestBody struct {
	PasswordName   string `json:"passwordName"`
	NewURL         string `json:"newUrl"`
	AccountName    string `json:"accountName"`
	MasterPassword string `json:"accountPassword"`
}
