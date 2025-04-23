package addnewpassword

type RequestBody struct {
	PasswordName string   `json:"passwordName"`
	Password     Password `json:"passwordToAdd"`
	Token        string   `json:"token"`
}

type Password struct {
	Password string `json:"password"`
	URL      string `json:"url"`
	Username string `json:"username"`
}
