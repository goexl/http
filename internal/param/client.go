package param

import (
	"crypto/tls"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	// 超时
	Timeout *Timeout `json:"timeout,omitempty"`
	// 连接池
	Pool *Pool `json:"pool,omitempty"`
	// 代理列表
	Proxies []*Proxy `json:"proxies,omitempty"`
	// 授权配置
	Auth *Auth `json:"auth,omitempty"`
	// Body数据传输控制
	Payload bool `json:"payload,omitempty"`
	// 警告消息开关
	Warning bool `json:"warning,omitempty"`
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
		Timeout: NewTimeout(),
		Pool:    NewPool(),
		Payload: true,
		Proxies: make([]*Proxy, 0),
		Queries: make(map[string]string),
		Forms:   make(map[string]string),
		Headers: make(map[string]string),
		Cookies: make([]*http.Cookie, 0),
	}
}

func (c *Client) Init(client *resty.Client) {
	client.SetTimeout(c.Timeout.Connection) // 启用连接池和长连接
	client.SetCloseConnection(false)        // 不关闭连接
	client.SetTransport(&http.Transport{    // 设置连接池配置
		MaxIdleConns:        c.Pool.All,     // 最大空闲连接数
		MaxIdleConnsPerHost: c.Pool.Host,    // 每个机器最大空闲连接数
		IdleConnTimeout:     c.Timeout.Idle, // 空闲连接超时时间

		TLSHandshakeTimeout: c.Timeout.Handshake, // 握手超时
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify:     true,  // 不验证证书
			SessionTicketsDisabled: false, // 启用复用功能
		},

		ForceAttemptHTTP2: true,  // 启用
		DisableKeepAlives: false, // 保持长链接
	})

	client.SetAllowGetMethodPayload(c.Payload)
	client.SetHeaders(c.Headers)
	client.SetQueryParams(c.Queries)
	client.SetFormData(c.Forms)
	client.SetCookies(c.Cookies)
	client.SetDisableWarn(!c.Warning)
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
