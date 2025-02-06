package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbasehq/socketbase/pkg/middlewares"
	"github.com/socketbasehq/socketbase/pkg/pkg/server"
	"github.com/socketbasehq/socketbase/pkg/pkg/types"
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
