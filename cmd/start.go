package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/socketbasehq/socketbase/pkg/config"
	"github.com/socketbasehq/socketbase/pkg/db"
	"github.com/socketbasehq/socketbase/pkg/modules"
	"github.com/socketbasehq/socketbase/pkg/modules/auth"
	"github.com/socketbasehq/socketbase/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			modules.Modules,
			fx.Provide(config.NewConfig),
			fx.Provide(server.NewServer),
			fx.Invoke(func(app *gin.Engine, cfg *config.Config) {
				dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
				db.ConnectDB(dns)
				auth.CreateAdminUser()
				go app.Run(cfg.Server.Host + ":" + cfg.Server.Port)
			}),
		)

		app.Run()
	},
}
