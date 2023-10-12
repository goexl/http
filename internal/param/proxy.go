package param

import (
	"fmt"
	"net/url"
)

type Proxy struct {
	Host     string `json:"host,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Target   string `json:"target,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewProxy() *Proxy {
	return new(Proxy)
}

func (p *Proxy) Addr() (addr string) {
	if "" != p.Username && "" != p.Password {
		addr = fmt.Sprintf(
			"%s://%s:%s@%s",
			p.Scheme,
			url.QueryEscape(p.Username), url.QueryEscape(p.Password),
			p.Host,
		)
	} else {
		addr = fmt.Sprintf("%s://%s", p.Scheme, p.Host)
	}

	return
}
