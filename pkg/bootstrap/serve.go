package bootstrap

import (
	"demoBlog/internal/modules/home/routes"
	"demoBlog/pkg/config"
	"demoBlog/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routes.Routes(routing.GetRouter())
	
	routing.Serve()
}