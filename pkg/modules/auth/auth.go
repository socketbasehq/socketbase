package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbasehq/socketbase/middlewares"
	"github.com/socketbasehq/socketbase/pkg/server"
	"github.com/socketbasehq/socketbase/pkg/types"
)

func NewAuthModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{
			Prefix: "/api/auth",

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
