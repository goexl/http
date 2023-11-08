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

func (p *Proxy) Basic(username string, password string) (proxy *Proxy) {
	p.params.Username = username
	p.params.Password = password
	proxy = p

	return
}

func (p *Proxy) Scheme(scheme string) (proxy *Proxy) {
	p.params.Scheme = scheme
	proxy = p

	return
}

func (p *Proxy) Target(target string) (proxy *Proxy) {
	p.params.Target = target
	proxy = p

	return
}

func (p *Proxy) Exclude(exclude string) (proxy *Proxy) {
	p.params.Exclude = exclude
	proxy = p

	return
}

func (p *Proxy) Host(host string) (proxy *Proxy) {
	p.params.Host = host
	proxy = p

	return
}

func (p *Proxy) Build() (core *Client) {
	p.core.params.Proxies = append(p.core.params.Proxies, &param.Proxy{
		Host:     p.params.Host,
		Scheme:   p.params.Scheme,
		Target:   p.params.Target,
		Exclude:  p.params.Exclude,
		Username: p.params.Username,
		Password: p.params.Password,
	})
	core = p.core

	return
}
