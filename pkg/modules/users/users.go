package users

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbasehq/socketbase/middlewares"
	"github.com/socketbasehq/socketbase/pkg/server"
	"github.com/socketbasehq/socketbase/pkg/types"
)

func NewUsersModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{
			Prefix: "/api/users",
			Middlewares: []gin.HandlerFunc{
				middlewares.CheckAuth,
			},
			Routes: []server.Route{},
		},
	}
}
