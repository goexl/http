package param

import (
	"net/http"
	"time"
)

type Cookie struct {
	Name  string
	Value string

	Path    string
	Domain  string
	Expires time.Time

	Age      int
	Secure   bool
	Httponly bool
	SameSite http.SameSite
	Raw      string
	Unparsed []string
}

func NewCookie() *Cookie {
	return new(Cookie)
}
