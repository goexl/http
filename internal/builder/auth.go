package builder

import (
	"github.com/goexl/http/internal/param"
)

type Auth struct {
	core   *Client
	params *param.Auth
}

func NewAuth(core *Client) *Auth {
	return &Auth{
		core:   core,
		params: param.NewAuth(),
	}
}

func (a *Auth) Basic(username string, password string) (auth *Auth) {
	a.params.Username = username
	a.params.Password = password
	auth = a

	return
}

func (a *Auth) Scheme(scheme string) (auth *Auth) {
	a.params.Scheme = scheme
	auth = a

	return
}

func (a *Auth) Token(token string) (auth *Auth) {
	a.params.Token = token
	auth = a

	return
}

func (a *Auth) Build() (core *Client) {
	a.core.params.Auth = a.params
	core = a.core

	return
}
