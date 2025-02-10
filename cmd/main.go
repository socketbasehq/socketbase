package main

import (
	"github.com/socketbasehq/socketbase/pkg/pkg/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "socketbase",
	Short: "socketbase - an open-source, self-hosted alternative to Pusher Channels",
	Long: `Socketbase enables real-time bidirectional communication in your applications.
	
Built with Go, it provides a reliable and scalable solution for developers who want 
to maintain control over their real-time infrastructure.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	config.LoadEnvs()
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.Execute()
}
