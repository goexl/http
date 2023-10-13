package core

import (
	"github.com/go-resty/resty/v2"
)

type Request struct {
	*resty.Request
}
