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

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/http/internal/interanl"
	"github.com/goexl/http/internal/interanl/constant"
	"github.com/goexl/http/internal/param"
	"github.com/goexl/simaqian"
)

type Client struct {
	*resty.Client

	logger  simaqian.Logger
	proxies map[string]*param.Proxy
	_       gox.CannotCopy
}

func NewClient(params *param.Client) (client *Client) {
	client = new(Client)
	client.Client = resty.New()
	client.logger = params.Logger
	client.proxies = make(map[string]*param.Proxy)

	params.Init(client.Client)
	// 设置动态代理
	if nil != params.Proxy && "" == params.Proxy.Target && 0 == len(params.Proxies) {
		addr := params.Proxy.Addr()
		client.SetProxy(addr)
		params.Logger.Debug("设置通用代理服务器", field.New("proxy", addr))
	} else {
		if nil != params.Proxy {
			target := gox.Ift("" == params.Proxy.Target, constant.TargetDefault, params.Proxy.Target)
			client.proxies[target] = params.Proxy
		}
		for _, proxy := range params.Proxies {
			client.proxies[proxy.Target] = proxy
		}
	}
	// 动态代理
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
	c.logger.Debug("向服务器发送请求", fields...)

	return
}

func (c *Client) setProxy(client *resty.Client, req *resty.Request) (err error) {
	if host, he := c.host(req.URL); nil != he {
		err = he
	} else if hp, hostOk := c.proxies[host]; hostOk {
		addr := hp.Addr()
		client.SetProxy(addr)
		c.logger.Debug("设置代理服务器", field.New("url", req.URL), field.New("Proxy", addr))
	} else if dp, defaultOk := c.proxies[constant.TargetDefault]; defaultOk {
		addr := dp.Addr()
		client.SetProxy(addr)
		c.logger.Debug("设置代理服务器", field.New("url", req.URL), field.New("Proxy", addr))
	}

	return
}

func (c *Client) unsetProxy(client *resty.Client, _ *resty.Response) (err error) {
	client.RemoveProxy()

	return
}

func (c *Client) host(raw string) (host string, err error) {
	if _url, ue := url.Parse(raw); nil != ue {
		err = ue
	} else {
		host = _url.Host
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
