package bootstrap

import (
	"xzl_blog/pkg/config"
	"xzl_blog/pkg/database"
	"xzl_blog/pkg/html"
	"xzl_blog/pkg/routing"
	"xzl_blog/pkg/sessions"
	"xzl_blog/pkg/static"
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
