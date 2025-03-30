package deletepassword

type RequestBody struct {
	PasswordName   string `json:"passwordName"`
	AccountName    string `json:"accountName"`
	MasterPassword string `json:"accountPassword"`
}
