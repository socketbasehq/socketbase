package users

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbase/socketbase/middlewares"
	"github.com/socketbase/socketbase/pkg/server"
	"github.com/socketbase/socketbase/pkg/types"

	"go.uber.org/fx"
)

func NewUsersModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{
			Prefix: "/users",
			Middlewares: []gin.HandlerFunc{
				middlewares.CheckAuth,
			},
			Routes: []server.Route{},
		},
	}
}

var Users = fx.Annotate(
	NewUsersModule,
	fx.ResultTags(`group:"routes"`),
)
