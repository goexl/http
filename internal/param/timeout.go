package param

import (
	"time"
)

type Timeout struct {
	Connection time.Duration `json:"connection,omitempty"`
	Response   time.Duration `json:"response,omitempty"`
	Idle       time.Duration `json:"idle,omitempty"`
	Handshake  time.Duration `json:"handshake,omitempty"`
}

func NewTimeout() *Timeout {
	return &Timeout{
		Connection: 120 * time.Second,
		Response:   600 * time.Second,
		Idle:       90 * time.Second,
		Handshake:  10 * time.Second,
	}
}
