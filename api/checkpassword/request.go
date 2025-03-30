package checkpassword

type RequestBody struct {
	AccountName    string `json:"accountName"`
	MasterPassword string `json:"accountPassword"`
}
