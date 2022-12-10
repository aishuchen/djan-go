package shortcuts

import (
	"djan-go/httpx"
	"net/http"
)

func MethodNotAllowed(_ *httpx.Request) httpx.Response {
	return &httpx.HTTPResponse{Status: http.StatusMethodNotAllowed}
}

func NotFound(_ *httpx.Request) httpx.Response {
	return &httpx.HTTPResponse{Status: http.StatusNotFound, Data: []byte("404 Not Found")}
}
