package http

import (
	"github.com/goexl/http/internal/builder"
	"github.com/goexl/http/internal/core"
)

var _ = New

// Client 客户端
type Client = core.Client

// New 创建客户端
func New() *builder.Client {
	return builder.NewClient()
}
