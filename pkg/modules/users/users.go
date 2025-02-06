package users

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbase/socketbase/middlewares"
	"github.com/socketbase/socketbase/pkg/server"
	"github.com/socketbase/socketbase/pkg/types"
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
