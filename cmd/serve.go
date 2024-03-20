package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"demoBlog/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

	config.Set()

	configs := config.Get()

  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
			"app name": viper.Get("App.Name"),
    })
  })
  r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}