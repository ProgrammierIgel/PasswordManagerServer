package manager

import "time"

type Token struct {
	Timestamp      time.Time
	AccountName    string
	MasterPassword string
}
