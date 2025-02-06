package socketbase

import (
	"github.com/socketbasehq/socketbase/pkg/pkg/server"
	"github.com/socketbasehq/socketbase/pkg/pkg/types"
)

func NewSocketbaseModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{
			Prefix: "/app",
			Routes: []server.Route{
				{
					Method:  "GET",
					Path:    ":id",
					Handler: handleAppRoute,
				},
			},
		},
	}
}
