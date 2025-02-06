package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbasehq/socketbase/pkg/middlewares"
	"github.com/socketbasehq/socketbase/pkg/pkg/server"
	"github.com/socketbasehq/socketbase/pkg/pkg/types"
)

func NewAppsModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{
			Prefix:      "/api/apps",
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
