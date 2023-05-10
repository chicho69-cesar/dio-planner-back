package routes

import "github.com/kataras/iris/v12"

func HelloWorld(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "Hello World!",
	})
}
