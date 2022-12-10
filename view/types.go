package view

import "djan-go/httpx"

type HandleFunc func(req *httpx.Request) httpx.Response

type View struct {
	Get    HandleFunc
	Post   HandleFunc
	Put    HandleFunc
	Patch  HandleFunc
	Delete HandleFunc
}

type GetView interface {
	Get(r *httpx.Request) httpx.Response
}

type PostView interface {
	Post(r *httpx.Request) httpx.Response
}

type PutView interface {
	Put(r *httpx.Request) httpx.Response
}

type PatchView interface {
	Patch(r *httpx.Request) httpx.Response
}

type DeleteView interface {
	Delete(r *httpx.Request) httpx.Response
}

type HeadView interface {
	Head(r *httpx.Request) httpx.Response
}

type OptionsView interface {
	Options(r *httpx.Request) httpx.Response
}
