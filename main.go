package main

import (
	"deku/controllers"
	"deku/repositories"
	"deku/services"
	"deku/sources"
	"github.com/kataras/iris"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Use(recover2.New())
	app.Logger().SetLevel("debug")
	mvc.Configure(app.Party("/posts"), posts)

	app.Run(iris.Addr(":9001"))
}

func posts(app *mvc.Application) {
	repo := repositories.NewPostRepository(sources.Posts)
	service := services.NewPostService(repo)
	app.Register(service)

	app.Handle(new(controllers.PostsController))
}
