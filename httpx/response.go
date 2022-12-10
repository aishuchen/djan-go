package httpx

import (
	"bytes"
	"io"
	"net/http"
)

type HTTPResponse struct {
	Status  int
	Data    []byte
	headers http.Header
}

func (hr *HTTPResponse) Body() io.ReadWriter {
	r := bytes.NewBuffer(hr.Data)
	return r
}

func (hr *HTTPResponse) StatusCode() int {
	if hr.Status == 0 {
		hr.Status = 200
	}
	return hr.Status
}

func (hr *HTTPResponse) Headers() http.Header {
	if hr.headers == nil {
		hr.headers = make(http.Header)
	}
	return hr.headers
}
