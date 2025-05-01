package deleteaccount

type RequestBody struct {
	Token       string `json:"token"`
	AccountName string `json:"accountname"`
}
