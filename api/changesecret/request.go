package changesecret

type RequestBody struct {
	PasswordName   string `json:"passwordName"`
	NewSecret        string `json:"newSecret"`
	AccountName    string `json:"accountName"`
	MasterPassword string `json:"accountPassword"`
}
