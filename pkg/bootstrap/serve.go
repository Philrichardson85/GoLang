package bootstrap

import (
	"demoBlog/pkg/config"
	"demoBlog/pkg/database"
	"demoBlog/pkg/html"
	"demoBlog/pkg/routing"
	"demoBlog/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()
 	
	routing.Serve()
	
}