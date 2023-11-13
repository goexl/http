package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/http/internal/interanl"
	"github.com/goexl/http/internal/param"
)

type Client struct {
	*resty.Client

	proxies []*param.Proxy
	_       gox.CannotCopy
}

func NewClient(params *param.Client) (client *Client) {
	client = new(Client)
	client.Client = resty.New()
	client.proxies = params.Proxies

	params.Init(client.Client)
	// 设置动态代理
	if 0 != len(client.proxies) {
		client.OnBeforeRequest(client.setProxy)
		client.OnAfterResponse(client.unsetProxy)
	}
	// 记录日志
	client.SetPreRequestHook(client.log)

	return
}

func (c *Client) Curl(rsp *resty.Response) (string, error) {
	return c.curl(rsp.Request)
}

func (c *Client) Fields(rsp *resty.Response) (fields gox.Fields[any]) {
	if nil == rsp {
		return
	}

	fields = gox.Fields[any]{
		field.New("url", rsp.Request.URL),
		field.New("status", rsp.StatusCode()),
		field.New("body", string(rsp.Body())),
	}

	return
}

func (c *Client) log(_ *resty.Client, req *http.Request) (err error) {
	fields := gox.Fields[any]{
		field.New("url", req.URL),
	}
	if nil != req.Body {
		if body, re := io.ReadAll(req.Body); nil != re {
			err = re
		} else {
			req.Body = interanl.NopCloser{Reader: bytes.NewBuffer(body)}
			fields = append(fields, field.New[json.RawMessage]("body", body))
		}
	}
	if nil != err {
		return
	}

	for key, value := range req.Header {
		if 1 == len(value) {
			fields = append(fields, field.New(fmt.Sprintf("header.%s", key), value[0]))
		} else if 1 < len(value) {
			fields = append(fields, field.New(fmt.Sprintf("header.%s", key), value))
		}
	}
	// TODO c.logger.Debug("向服务器发送请求", fields...)

	return
}

func (c *Client) setProxy(client *resty.Client, req *resty.Request) (err error) {
	if host, he := c.host(req.URL); nil != he {
		err = he
	} else if addr, settable := c.canSetProxy(host); settable {
		client.SetProxy(addr)
	} else if !settable && client.IsProxySet() { // ! 如果不需要设置代理而客户端又被设置了代理，需要去除代理
		client.RemoveProxy()
	}

	return
}

func (c *Client) unsetProxy(client *resty.Client, _ *resty.Response) (err error) {
	if client.IsProxySet() {
		client.RemoveProxy()
	}

	return
}

func (c *Client) host(raw string) (host string, err error) {
	if link, ue := url.Parse(raw); nil != ue {
		err = ue
	} else {
		host = link.Host
	}

	return
}

func (c *Client) curl(req *resty.Request) (curl string, err error) {
	command := new(strings.Builder)
	command.WriteString("curl")
	command.WriteString("-X")
	command.WriteString(c.bashEscape(req.Method))

	if nil != req.Body {
		if body, re := io.ReadAll(req.RawRequest.Body); nil != re {
			err = re
		} else {
			req.Body = interanl.NopCloser{Reader: bytes.NewBuffer(body)}
			command.WriteString("-d")
			command.WriteString(c.bashEscape(string(body)))
		}
	}
	if nil != err {
		return
	}

	keys := make([]string, 0, len(req.Header))
	for key := range req.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		command.WriteString("-H")
		command.WriteString(c.bashEscape(fmt.Sprintf("%s: %s", key, strings.Join(req.Header[key], " "))))
	}
	command.WriteString(c.bashEscape(req.URL))

	return
}

func (c *Client) bashEscape(from string) string {
	return `'` + strings.Replace(from, `'`, `'\''`, -1) + `'`
}

func (c *Client) canSetProxy(host string) (addr string, settable bool) {
	for _, proxy := range c.proxies {
		if proxy.Targeted(host) && !proxy.Excluded(host) {
			addr = proxy.Addr()
			settable = true
		}
		if settable {
			break
		}
	}

	return
}
