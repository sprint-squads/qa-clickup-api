package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sprint-squads/qa-clickup-api/internal/api"
	"os"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	Long:  "Starts a http server and serves the configured api",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := api.NewServer()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		server.Start()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
