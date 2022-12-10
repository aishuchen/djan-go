package main

import (
	"djan-go/conf/settings"
	"djan-go/httpx"
	"djan-go/middleware"
	"djan-go/urlx"
	"djan-go/urlx/shortcuts"
	"djan-go/view"
	"io"
	"net/http"
)

type DjangGo struct {
	Settings *settings.Settings

	rw          http.ResponseWriter
	router      *urlx.Router
	middlewares []middleware.Middleware
}

func New(settings *settings.Settings) *DjangGo {
	dj := &DjangGo{
		router:   urlx.NewRouter(),
		Settings: settings,
	}
	return dj
}

func (dj *DjangGo) Path(path string, handler interface{}) {
	dj.router.Register(path, handler)
}

func (dj *DjangGo) Run(addr string) error {
	return http.ListenAndServe(addr, dj)
}

func (dj *DjangGo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := httpx.NewRequest(r)
	dj.rw = w
	dj.serve(req)
}

func (dj *DjangGo) init() {
	if dj.Settings != nil {
		dj.middlewares = dj.Settings.Middlewares
	}
}

func (dj *DjangGo) serve(req *httpx.Request) {

	// middleware
	for _, mw := range dj.middlewares {
		mw.ProcessRequest(req)
	}

	// View or Handler
	handler := dj.getHandler(req.URL.Path, req.Method)

	// Handle response
	resp := handler(req)
	for _, mw := range dj.middlewares {
		mw.ProcessResponse(resp)
	}

	// Return Response
	body, err := io.ReadAll(resp.Body())
	if err != nil {
		panic(err)
	}
	dj.rw.WriteHeader(resp.StatusCode())
	dj.rw.Write(body)
}

func (dj *DjangGo) getHandler(path, method string) view.HandleFunc {
	handlerIface := dj.router.Find(path)
	if handlerIface == nil {
		return shortcuts.NotFound
	}

	// got HandlerFunc
	if handler, ok := handlerIface.(func(*httpx.Request) httpx.Response); ok {
		return handler
	}
	if handler, ok := handlerIface.(view.HandleFunc); ok {
		return handler
	}
	// got View
	switch method {
	case http.MethodGet:
		if handler, ok := handlerIface.(view.GetView); ok {
			return handler.Get
		}
	case http.MethodPost:
		if handler, ok := handlerIface.(view.PostView); ok {
			return handler.Post
		}
	case http.MethodPut:
		if handler, ok := handlerIface.(view.PutView); ok {
			return handler.Put
		}
	case http.MethodPatch:
		if handler, ok := handlerIface.(view.PatchView); ok {
			return handler.Patch
		}
	case http.MethodDelete:
		if handler, ok := handlerIface.(view.DeleteView); ok {
			return handler.Delete
		}
	case http.MethodHead:
		if handler, ok := handlerIface.(view.HeadView); ok {
			return handler.Head
		}
	case http.MethodOptions:
		if handler, ok := handlerIface.(view.OptionsView); ok {
			return handler.Options
		}
	case http.MethodTrace:
		// ...
		return shortcuts.MethodNotAllowed
	}
	return shortcuts.MethodNotAllowed
}
