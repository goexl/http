package param

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	// 超时
	Timeout time.Duration `json:"timeout,omitempty"`
	// 代理列表
	Proxies []*Proxy `json:"proxies,omitempty"`
	// 授权配置
	Auth *Auth `json:"auth,omitempty"`
	// Body数据传输控制
	Payload bool `json:"payload,omitempty"`
	// 秘钥配置
	Certificate *Certificate `json:"certificate,omitempty"`
	// 通用的查询参数
	Queries map[string]string `json:"queries,omitempty"`
	// 表单参数
	Forms map[string]string `json:"forms,omitempty"`
	// 通用头信息
	Headers map[string]string `json:"headers,omitempty"`
	// 信息
	Cookies []*http.Cookie `json:"cookies,omitempty"`
}

func NewClient() *Client {
	return &Client{
		Payload: true,
		Proxies: make([]*Proxy, 0),
		Queries: make(map[string]string),
		Forms:   make(map[string]string),
		Headers: make(map[string]string),
		Cookies: make([]*http.Cookie, 0),
	}
}

func (c *Client) Init(client *resty.Client) {
	client.SetTimeout(c.Timeout)
	client.SetAllowGetMethodPayload(c.Payload)
	client.SetHeaders(c.Headers)
	client.SetQueryParams(c.Queries)
	client.SetFormData(c.Forms)
	client.SetCookies(c.Cookies)
	c.auth(client)
	c.certificate(client)
}

func (c *Client) certificate(client *resty.Client) {
	if nil == c.Certificate {
		return
	}

	if c.Certificate.Skip {
		// nolint:gosec
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	} else {
		c.root(client)
	}
}

func (c *Client) root(client *resty.Client) {
	if "" != c.Certificate.Root {
		client.SetRootCertificate(c.Certificate.Root)
	}
	if 0 != len(c.Certificate.Clients) {
		certificates := make([]tls.Certificate, 0, len(c.Certificate.Clients))
		for _, c := range c.Certificate.Clients {
			cert, ce := tls.LoadX509KeyPair(c.Public, c.Private)
			if nil != ce {
				continue
			}
			certificates = append(certificates, cert)
		}
		client.SetCertificates(certificates...)
	}
}

func (c *Client) auth(client *resty.Client) {
	if nil == c.Auth {
		return
	}

	if "" != c.Auth.Username && "" != c.Auth.Password {
		client.SetBasicAuth(c.Auth.Username, c.Auth.Password)
	} else if "" != c.Auth.Scheme {
		client.SetAuthToken(c.Auth.Token)
		client.SetAuthScheme(c.Auth.Scheme)
	}
}
