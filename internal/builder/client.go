package builder

import (
	"time"

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

func (c *Client) Query(key string, value string) (builder *Client) {
	c.params.Queries[key] = value
	builder = c

	return
}

func (c *Client) Form(key string, value string) (builder *Client) {
	c.params.Forms[key] = value
	builder = c

	return
}

func (c *Client) Header(key string, value string) (builder *Client) {
	c.params.Headers[key] = value
	builder = c

	return
}
