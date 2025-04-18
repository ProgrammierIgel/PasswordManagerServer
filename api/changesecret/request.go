package changesecret

type RequestBody struct {
	PasswordName string `json:"passwordName"`
	NewSecret    string `json:"newSecret"`
	Token        string `json:"token"`
}
