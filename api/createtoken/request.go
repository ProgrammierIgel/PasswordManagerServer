package createtoken

type RequestBody struct {
	Password    string `json:"password"`
	AccountName string `json:"accountname"`
}
