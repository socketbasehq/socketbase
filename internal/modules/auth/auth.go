package auth

import (
	"github.com/socketbase/socketbase/internal/pkg/server"
	"go.uber.org/fx"
)

func NewAuthModule() server.Module {
	return server.Module{
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
			},
		},
	}
}

var Auth = fx.Annotate(
	NewAuthModule,
	fx.ResultTags(`group:"routes"`),
)
