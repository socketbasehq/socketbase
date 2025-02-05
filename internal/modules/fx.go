package modules

import (
	"github.com/socketbase/socketbase/internal/modules/auth"
	"github.com/socketbase/socketbase/internal/modules/users"
	"github.com/socketbase/socketbase/internal/modules/ws"

	"go.uber.org/fx"
)

var Modules = fx.Provide(
	auth.NewAuthModule,
	users.NewUsersModule,
	ws.NewWSModule,
)
