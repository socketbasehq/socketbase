package common

import "github.com/gin-gonic/gin"

type Router struct {
	path    string
	method  string
	handler gin.HandlerFunc
}

func NewRouter(path string, method string, handler gin.HandlerFunc) *Router {
	return &Router{path: path, method: method, handler: handler}
}

func (r *Router) Register(engine *gin.Engine) {
	engine.Handle(r.method, r.path, r.handler)
}
