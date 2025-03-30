package addnewpassword

type RequestBody struct {
	PasswordName   string `json:"passwordName"`
	Password       string `json:"passwordToAdd"`
	AccountName    string `json:"accountName"`
	MasterPassword string `json:"accountPassword"`
}
