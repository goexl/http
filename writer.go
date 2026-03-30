package http

import (
	"net/http"
)

// ResponseWriter 避免包冲突
type ResponseWriter = http.ResponseWriter
