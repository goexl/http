package builder

import (
	"github.com/goexl/http/internal/param"
)

type Proxy struct {
	core   *Client
	params *param.Proxy
}

func NewProxy(core *Client) *Proxy {
	return &Proxy{
		core:   core,
		params: param.NewProxy(),
	}
}

func (p *Proxy) Basic(username string, password string) (auth *Proxy) {
	p.params.Username = username
	p.params.Password = password
	auth = p

	return
}

func (p *Proxy) Scheme(scheme string) (auth *Proxy) {
	p.params.Scheme = scheme
	auth = p

	return
}

func (p *Proxy) Target(target string) (auth *Proxy) {
	p.params.Target = target
	auth = p

	return
}

func (p *Proxy) Exclude(exclude string) (auth *Proxy) {
	p.params.Exclude = exclude
	auth = p

	return
}

func (p *Proxy) Host(host string) (auth *Proxy) {
	p.params.Host = host
	auth = p

	return
}

func (p *Proxy) Build() (core *Client) {
	p.core.params.Proxies = append(p.core.params.Proxies, &param.Proxy{
		Host:     p.params.Host,
		Scheme:   p.params.Scheme,
		Target:   p.params.Target,
		Username: p.params.Username,
		Password: p.params.Password,
	})
	core = p.core

	return
}
