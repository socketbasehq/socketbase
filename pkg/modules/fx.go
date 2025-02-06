package modules

import (
	"github.com/socketbase/socketbase/pkg/modules/apps"
	"github.com/socketbase/socketbase/pkg/modules/auth"
	"github.com/socketbase/socketbase/pkg/modules/users"
	"github.com/socketbase/socketbase/pkg/modules/ws"

	"go.uber.org/fx"
)

var Modules = fx.Provide(
	ws.NewWSModule,
	auth.NewAuthModule,
	users.NewUsersModule,
	apps.NewAppsModule,
)
