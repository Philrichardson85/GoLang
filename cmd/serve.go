package cmd

import (
	"demoBlog/pkg/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
  Use:   "serve",
  Short: "Serve a app on dev server",
  Long:  `Server will be served on host and port defined in the config.yml`,
  Run: func(cmd *cobra.Command, args []string) {
    serve()
  },
}


func serve() {

    bootstrap.Serve()
}