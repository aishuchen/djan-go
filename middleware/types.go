package middleware

import "djan-go/httpx"

type Middleware interface {
	ProcessRequest(req *httpx.Request)
	ProcessResponse(resp httpx.Response)
}
