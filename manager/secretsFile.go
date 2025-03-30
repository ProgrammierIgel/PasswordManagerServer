package manager

type SecretsFile struct {
	MainPasswords map[string]Password          `json:"mainPasswords"`
	Secrets       map[string]map[string]string `json:"secrets"`
}
