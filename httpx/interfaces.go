package httpx

import (
	"io"
	"net/http"
)

type Response interface {
	StatusCode() int
	Body() io.ReadWriter
	Headers() http.Header
}
