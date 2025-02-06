package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbase/socketbase/middlewares"
	"github.com/socketbase/socketbase/pkg/server"
	"github.com/socketbase/socketbase/pkg/types"
	"go.uber.org/fx"
)

func NewAuthModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{
			Prefix: "/auth",

			Routes: []server.Route{
				{
					Method:  "POST",
					Path:    "/login",
					Handler: handleLogin,
				},
				{
					Method:  "POST",
					Path:    "/register",
					Handler: handleRegister,
				},
				{
					Method:  "GET",
					Path:    "/me",
					Handler: handleGetMe,
					Middlewares: []gin.HandlerFunc{
						middlewares.CheckAuth,
					},
				},
			},
		},
	}
}

var Auth = fx.Annotate(
	NewAuthModule,
	fx.ResultTags(`group:"routes"`),
)
