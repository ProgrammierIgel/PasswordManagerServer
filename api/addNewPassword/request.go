package addnewpassword

type RequestBody struct {
	PasswordName   string   `json:"passwordName"`
	Password       Password `json:"passwordToAdd"`
	AccountName    string   `json:"accountName"`
	MasterPassword string   `json:"accountPassword"`
}

type Password struct {
	Password string `json:"password"`
	URL      string `json:"url"`
	Username string `json:"username"`
}
