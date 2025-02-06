package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbase/socketbase/middlewares"
	"github.com/socketbase/socketbase/pkg/server"
	"github.com/socketbase/socketbase/pkg/types"
	"go.uber.org/fx"
)

func NewAppsModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{
			Prefix:      "/apps",
			Middlewares: []gin.HandlerFunc{middlewares.CheckAuth},
			Routes: []server.Route{
				{
					Method:  "GET",
					Path:    "/",
					Handler: handleListApps,
				},
				{
					Method:  "POST",
					Path:    "/",
					Handler: handleCreateApp,
				},
				{
					Method:  "GET",
					Path:    "/:id",
					Handler: handleGetApp,
				},
			},
		},
	}
}

var Apps = fx.Annotate(
	NewAppsModule,
	fx.ResultTags(`group:"routes"`),
)
