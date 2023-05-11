package main

import (
	"github.com/chicho69-cesar/dio-planner-back/routes"
	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()
	storage.InitializeDB()

	app := iris.Default()
	app.Validator = validator.New()

	test := app.Party("/test")
	{
		test.Get("/", routes.HelloWorld)
	}

	user := app.Party("/user") 
	{
		user.Post("/register", routes.Register)
		user.Post("/login", routes.Login)
		user.Post("/facebook", routes.FacebookLoginOrSignUp)
		user.Post("/google", routes.GoogleLoginOrSignUp)
		user.Post("/apple", routes.AppleLoginOrSignUp)
	}

	app.Listen(":4000")
}
