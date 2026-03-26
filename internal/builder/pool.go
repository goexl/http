package builder

import (
	"time"

	"github.com/goexl/http/internal/param"
)

type Pool struct {
	core   *Client
	params *param.Pool
}

func NewPool(core *Client) *Pool {
	return &Pool{
		core:   core,
		params: param.NewPool(),
	}
}

func (p *Pool) All(pool int) *Pool {
	return p.set(func() {
		p.params.All = pool
	})
}

func (p *Pool) Host(pool int) *Pool {
	return p.set(func() {
		p.params.Host = pool
	})
}

func (p *Pool) Build() (core *Client) {
	p.core.params.Pool = p.params
	core = p.core

	return
}

func (p *Pool) set(callback func()) (pool *Pool) {
	callback()
	pool = p

	return
}
