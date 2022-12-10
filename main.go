package main

import (
	"djan-go/httpx"
)

type MyView struct {
}

func (v *MyView) Get(r *httpx.Request) httpx.Response {
	return &httpx.HTTPResponse{Status: 200, Data: []byte("hello world")}
}

func index(r *httpx.Request) httpx.Response {
	return &httpx.HTTPResponse{Status: 200, Data: []byte("index")}
}

func main() {
	dj := New(nil)
	dj.Path("/index", index)
	dj.Path("/helloworld", &MyView{})
	if err := dj.Run(":8000"); err != nil {
		panic(err)
	}
}
