package changepasswordname

type RequestBody struct {
	PasswordName    string `json:"passwordName"`
	NewPasswordName string `json:"newPasswordName"`
	Token           string `json:"token"`
}
