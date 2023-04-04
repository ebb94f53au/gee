package gee

import (
    "net/http"
)

const sub = "-"

type HandlerFunc func(*Context)

type Engine struct {
    router *router
}

func (engine *Engine) Post(path string, handler HandlerFunc) {
    engine.router.addRoute(http.MethodPost, path, handler)
}

func (engine *Engine) Get(path string, handler HandlerFunc) {
    engine.router.addRoute(http.MethodGet, path, handler)
}

func (engine *Engine) ServeHTTP(respWriter http.ResponseWriter, request *http.Request) {
    engine.router.handle(NewContext(respWriter, request))
}

func New() *Engine {
    return &Engine{
        router: newRouter(),
    }
}

func (engine *Engine) Run(addr string) (err error) {
    return http.ListenAndServe(addr, engine)
}

func (engine *Engine) getRouterKey(request *http.Request) string {
    return request.Method + sub + request.URL.Path
}
