package builder

import (
	"time"

	"github.com/goexl/http/internal/core"
	"github.com/goexl/http/internal/interanl"
	"github.com/goexl/http/internal/param"
)

type Client struct {
	params *param.Client
}

func NewClient() *Client {
	return &Client{
		params: param.NewClient(),
	}
}

func (c *Client) Timeout(timeout time.Duration) (builder *Client) {
	c.params.Timeout = timeout
	builder = c

	return
}

func (c *Client) Payload(payload bool) (builder *Client) {
	c.params.Payload = payload
	builder = c

	return
}

func (c *Client) Queries(queries map[string]string) (builder *Client) {
	c.params.Queries = queries
	builder = c

	return
}

func (c *Client) Query(key string, value string) (builder *Client) {
	c.params.Queries[key] = value
	builder = c

	return
}

func (c *Client) Forms(forms map[string]string) (builder *Client) {
	c.params.Forms = forms
	builder = c

	return
}

func (c *Client) Form(key string, value string) (builder *Client) {
	c.params.Forms[key] = value
	builder = c

	return
}

func (c *Client) Headers(headers map[string]string) (builder *Client) {
	c.params.Headers = headers
	builder = c

	return
}

func (c *Client) Header(key string, value string) (builder *Client) {
	c.params.Headers[key] = value
	builder = c

	return
}

func (c *Client) Logger(logger interanl.Logger) (builder *Client) {
	c.params.Logger = logger
	builder = c

	return
}

func (c *Client) Auth() *Auth {
	return NewAuth(c)
}

func (c *Client) Proxy() *Proxy {
	return NewProxy(c)
}

func (c *Client) Certificate() *Certificate {
	return NewCertificate(c)
}

func (c *Client) Cookie() *Cookie {
	return NewCookie(c)
}

func (c *Client) Build() *core.Client {
	return core.NewClient(c.params)
}
