package changeurl

type RequestBody struct {
	PasswordName string `json:"passwordName"`
	NewURL       string `json:"newUrl"`
	Token        string `json:"token"`
}
