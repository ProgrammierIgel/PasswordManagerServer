package manager

import "time"

type RemoteRequestAddress struct {
	LastRequest       time.Time
	NumberOfAccessess uint
}
