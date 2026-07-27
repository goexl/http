package param

import (
	"time"
)

type Timeout struct {
	Total     time.Duration `json:"total,omitempty"`
	Connect   time.Duration `json:"connect,omitempty"`
	Header    time.Duration `json:"header,omitempty"`
	Idle      time.Duration `json:"idle,omitempty"`
	Handshake time.Duration `json:"handshake,omitempty"`
}

func NewTimeout() *Timeout {
	return &Timeout{
		Total:     120 * time.Second,
		Connect:   10 * time.Second,
		Header:    600 * time.Second,
		Idle:      90 * time.Second,
		Handshake: 10 * time.Second,
	}
}
