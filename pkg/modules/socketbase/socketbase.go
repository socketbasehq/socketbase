package socketbase

import (
	"github.com/socketbasehq/socketbase/pkg/pkg/server"
	"github.com/socketbasehq/socketbase/pkg/pkg/types"
)

func NewSocketbaseModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{

			Routes: []server.Route{
				{
					Method:  "GET",
					Path:    "/app/:id",
					Handler: handleAppRoute,
				},
				{
					Method:  "POST",
					Path:    "/apps/:id/events",
					Handler: handleAppEvent,
				},
			},
		},
	}
}
