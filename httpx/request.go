package httpx

import (
	"net/http"
)

type Request struct {
	*http.Request
}

func NewRequest(r *http.Request) *Request {
	if r == nil {
		panic("nil http.Request")
	}
	req := &Request{Request: r}
	return req
}
