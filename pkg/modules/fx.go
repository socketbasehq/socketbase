package modules

import (
	"github.com/socketbasehq/socketbase/pkg/pkg/modules/apps"
	"github.com/socketbasehq/socketbase/pkg/pkg/modules/auth"
	"github.com/socketbasehq/socketbase/pkg/pkg/modules/socketbase"
	"github.com/socketbasehq/socketbase/pkg/pkg/modules/users"
	"github.com/socketbasehq/socketbase/pkg/pkg/modules/ws"

	"go.uber.org/fx"
)

var Modules = fx.Provide(
	ws.NewWSModule,
	auth.NewAuthModule,
	users.NewUsersModule,
	apps.NewAppsModule,
	socketbase.NewSocketbaseModule,
)
