package main

import (
	"github.com/socketbasehq/socketbase/migrate"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Migrate()
	},
}
