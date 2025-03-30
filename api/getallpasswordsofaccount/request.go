package getallpasswordsofaccount

type RequestBody struct {
	AccountName    string `json:"accountName"`
	MasterPassword string `json:"accountPassword"`
}
