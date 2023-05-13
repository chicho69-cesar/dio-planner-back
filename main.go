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
	storage.InitializePostgresDB()

	app := iris.Default()
	app.Validator = validator.New()

	test := app.Party("/test")
	{
		test.Get("/", routes.HelloWorld)
	}

	event := app.Party("/event")
	{
		event.Post("/create", routes.CreateEvent)
		event.Get("/get/{id}", routes.GetEventByID)
		event.Get("/get-events/{page}", routes.GetEvents)
	}

	// grade := app.Party("/grade")
	{

	}

	// guest := app.Party("/guest")
	{

	}

	// memory := app.Party("/memory")
	{

	}

	// purchase := app.Party("/purchase")
	{

	}

	// todo := app.Party("/todo")
	{

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
