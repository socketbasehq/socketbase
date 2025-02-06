package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Route struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

type RouteGroup struct {
	Prefix      string
	Routes      []Route
	Middlewares []gin.HandlerFunc
}

type ServerParams struct {
	fx.In
	RouteGroups []RouteGroup `group:"routes"`
}

func NewServer(params ServerParams) *gin.Engine {
	app := gin.Default()

	for _, group := range params.RouteGroups {
		router := app.Group(group.Prefix)

		for _, route := range group.Routes {
			middlewares := []gin.HandlerFunc{}

			if len(group.Middlewares) > 0 {
				middlewares = append(middlewares, group.Middlewares...)
			}
			if len(route.Middlewares) > 0 {
				middlewares = append(middlewares, route.Middlewares...)
			}

			if len(middlewares) > 0 {
				router.Handle(route.Method, route.Path, append(middlewares, route.Handler)...)
			} else {
				router.Handle(route.Method, route.Path, route.Handler)
			}
		}
	}

	return app
}
