package main

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbasehq/socketbase/pkg/pkg/db"
	"github.com/socketbasehq/socketbase/pkg/pkg/modules"
	"github.com/socketbasehq/socketbase/pkg/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		db.ConnectDB()

		app := fx.New(
			modules.Modules,
			fx.Provide(server.NewServer),
			fx.Invoke(func(app *gin.Engine) {
				go app.Run(":3000")
			}),
		)

		app.Run()
	},
}
