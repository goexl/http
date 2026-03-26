package builder

import (
	"time"

	"github.com/goexl/http/internal/param"
)

type Timeout struct {
	core   *Client
	params *param.Timeout
}

func NewTimeout(core *Client) *Timeout {
	return &Timeout{
		core:   core,
		params: param.NewTimeout(),
	}
}

func (t *Timeout) Connection(timeout time.Duration) *Timeout {
	return t.set(func() {
		t.params.Connection = timeout
	})
}

func (t *Timeout) Idle(timeout time.Duration) *Timeout {
	return t.set(func() {
		t.params.Idle = timeout
	})
}

func (t *Timeout) Handshake(timeout time.Duration) *Timeout {
	return t.set(func() {
		t.params.Handshake = timeout
	})
}

func (t *Timeout) Build() (core *Client) {
	t.core.params.Timeout = t.params
	core = t.core

	return
}

func (t *Timeout) set(callback func()) (timeout *Timeout) {
	callback()
	timeout = t

	return
}
