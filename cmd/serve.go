package cmd

import (
	"github.com/spf13/cobra"

	"demoBlog/pkg/config"
	"demoBlog/pkg/routing"
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

	routing.Init()

  router := routing.GetRouter()

   //set route
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
			"app name": viper.Get("App.Name"),
    })
  })

    routing.Serve()
}