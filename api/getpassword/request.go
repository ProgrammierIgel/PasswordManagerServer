package getpassword

type RequestBody struct {
	PasswordName string `json:"passwordName"`
	Token        string `json:"token"`
}
