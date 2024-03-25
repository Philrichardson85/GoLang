package bootstrap

import (
	"demoBlog/pkg/config"
	"demoBlog/pkg/database"
	"demoBlog/pkg/html"
	"demoBlog/pkg/routing"
	"demoBlog/pkg/static"
	"demoBlog/pkg/sessions"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()
 	
	routing.Serve()
	
}