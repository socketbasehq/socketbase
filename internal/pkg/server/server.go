package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type ServerParams struct {
	fx.In
	RouteGroups []RouteGroup `group:"routes"`
}

func NewServer(params ServerParams) *gin.Engine {
	app := gin.Default()

	// Register all route groups
	for _, group := range params.RouteGroups {
		router := app.Group(group.Prefix)
		for _, route := range group.Routes {
			router.Handle(route.Method, route.Path, route.Handler)
		}
	}

	return app
}
