package types

import (
	"github.com/socketbase/socketbase/pkg/server"
	"go.uber.org/fx"
)

type Module struct {
	fx.Out

	Routes server.RouteGroup `group:"routes"`
}
