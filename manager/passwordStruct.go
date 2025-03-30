package manager

type Password struct {
	PasswordHash string `json:"hash"`
	Salt         string `json:"salt"`
}
