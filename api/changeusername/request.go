package changeusername

type RequestBody struct {
	PasswordName string `json:"passwordName"`
	NewUsername  string `json:"newUsername"`
	Token        string `json:"token"`
}
