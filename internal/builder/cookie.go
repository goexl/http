package builder

import (
	"net/http"
	"time"

	"github.com/goexl/http/internal/param"
)

type Cookie struct {
	core   *Client
	params *param.Cookie
}

func NewCookie(core *Client) *Cookie {
	return &Cookie{
		core:   core,
		params: param.NewCookie(),
	}
}

func (a *Cookie) Name(name string) (cookie *Cookie) {
	a.params.Name = name
	cookie = a

	return
}

func (a *Cookie) Value(value string) (cookie *Cookie) {
	a.params.Value = value
	cookie = a

	return
}

func (a *Cookie) Path(path string) (cookie *Cookie) {
	a.params.Path = path
	cookie = a

	return
}

func (a *Cookie) Domain(domain string) (cookie *Cookie) {
	a.params.Domain = domain
	cookie = a

	return
}

func (a *Cookie) Expires(expires time.Time) (cookie *Cookie) {
	a.params.Expires = expires
	cookie = a

	return
}

func (a *Cookie) Age(age int) (cookie *Cookie) {
	a.params.Age = age
	cookie = a

	return
}

func (a *Cookie) Secure(secure bool) (cookie *Cookie) {
	a.params.Secure = secure
	cookie = a

	return
}

func (a *Cookie) Build() (core *Client) {
	a.core.params.Cookies = append(a.core.params.Cookies, &http.Cookie{
		Name:     a.params.Name,
		Value:    a.params.Value,
		Path:     a.params.Path,
		Domain:   a.params.Domain,
		Expires:  a.params.Expires,
		MaxAge:   a.params.Age,
		Secure:   a.params.Secure,
		HttpOnly: a.params.Httponly,
		SameSite: a.params.SameSite,
		Raw:      a.params.Raw,
		Unparsed: a.params.Unparsed,
	})
	core = a.core

	return
}
