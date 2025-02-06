package socketbase

import (
	"github.com/socketbase/socketbase/pkg/server"
	"github.com/socketbase/socketbase/pkg/types"
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
