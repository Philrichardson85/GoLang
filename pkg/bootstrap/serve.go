package bootstrap

import (
	"demoBlog/internal/modules/home/routes"
	"demoBlog/pkg/config"
	"demoBlog/pkg/html"
	"demoBlog/pkg/routing"
	"demoBlog/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())
 
	routes.Routes(routing.GetRouter())
	
	routing.Serve()
	
}