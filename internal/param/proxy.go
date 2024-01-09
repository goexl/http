package param

import (
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/goexl/gox"
)

type Proxy struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Target   string `json:"target,omitempty"`
	Exclude  string `json:"exclude,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewProxy() *Proxy {
	return new(Proxy)
}

func (p *Proxy) Uri() (uri string) {
	addr := gox.StringBuilder(p.Host)
	if 0 != p.Port {
		addr.Append(p.Port)
	}

	if "" != p.Username && "" != p.Password {
		uri = fmt.Sprintf(
			"%s://%s:%s@%s",
			p.Scheme,
			url.QueryEscape(p.Username), url.QueryEscape(p.Password),
			addr.String(),
		)
	} else {
		uri = fmt.Sprintf("%s://%s", p.Scheme, addr.String())
	}

	return
}

func (p *Proxy) Targeted(host string) bool {
	return "" == p.Target || p.match(p.Target, host)
}

func (p *Proxy) Excluded(host string) bool {
	return "" != p.Exclude && p.match(p.Exclude, host)
}

func (p *Proxy) match(target string, host string) (matched bool) {
	if host == target {
		matched = true
	} else if strings.Contains(target, host) {
		matched = true
	} else if mm, me := path.Match(target, host); nil == me && mm {
		matched = true
	} else if rm, re := regexp.MatchString(target, host); nil == re && rm {
		matched = true
	}

	return
}
