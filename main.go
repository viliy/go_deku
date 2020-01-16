package main

import (
	"deku/controllers"
	"deku/repositories"
	"deku/services"
	"deku/sources"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	mvc.Configure(app.Party("/posts"), posts)

	app.Run(
		iris.Addr(":9001"),
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

func posts(app *mvc.Application) {
	repo := repositories.NewPostRepository(sources.Posts)
	postService := services.NewPostService(repo)
	app.Register(postService)

	app.Handle(new(controllers.PostsController))
}
