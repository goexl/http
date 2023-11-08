package param

import (
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"
)

type Proxy struct {
	Host     string `json:"host,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Target   string `json:"target,omitempty"`
	Exclude  string `json:"exclude,omitempty"`
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
