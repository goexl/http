package builder

import (
	"github.com/goexl/http/internal/param"
)

type Certificate struct {
	core   *Client
	params *param.Certificate
}

func NewCertificate(core *Client) *Certificate {
	return &Certificate{
		core:   core,
		params: param.NewCertificate(),
	}
}

func (c *Certificate) Skip(skip bool) (auth *Certificate) {
	c.params.Skip = skip
	auth = c

	return
}

func (c *Certificate) Root(root string) (auth *Certificate) {
	c.params.Root = root
	auth = c

	return
}

func (c *Certificate) Client(public string, private string) (auth *Certificate) {
	c.params.Clients = append(c.params.Clients, &param.ClientCertificate{
		Public:  public,
		Private: private,
	})
	auth = c

	return
}

func (c *Certificate) Build() (core *Client) {
	core.params.Certificate = c.params
	core = c.core

	return
}
