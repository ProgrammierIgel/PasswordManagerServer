package addnewpassword

type RequestBody struct {
	PasswordName   string `json:"passwordName"`
	Password       string `json:"passwordToAdd"`
	AccountName    string `json:"accountName"`
	MasterPassword string `json:"accountPassword"`
	URL            string `json:"url"`
	Username       string `json:"username"`
}
