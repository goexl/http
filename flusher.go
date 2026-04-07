package http

import (
	"net/http"
)

// Flusher 避免包冲突
type Flusher = http.Flusher
