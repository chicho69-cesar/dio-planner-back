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
		event.Get("/get-user-events/{user_id}", routes.GetEventsByUser)
		event.Get("/get-query-events", routes.GetEventsByQuery)
		event.Patch("/update/{event_id}", routes.UpdateEvent)
		event.Delete("/delete/{event_id}", routes.DeleteEvent)
	}

	grade := app.Party("/grade")
	{
		grade.Post("/add-grade", routes.AddGrade)
		grade.Get("/get-grades/{event_id}", routes.GetGrades)
	}

	guest := app.Party("/guest")
	{
		guest.Post("/add-guest", routes.AddGuest)
		guest.Get("/search-guests", routes.SearchGuests)
		guest.Get("/get-guests/{event_id}", routes.GetGuests)
		guest.Get("/get-invitations/{user_id}", routes.GetInvitations)
		guest.Patch("/accept-invitation/{guest_id}", routes.AcceptInvitation)
		guest.Patch("/decline-invitation/{guest_id}", routes.DeclineInvitation)
		guest.Get("/get-my-events/{user_id}", routes.GetMyEvents)
	}

	memory := app.Party("/memory")
	{
		memory.Get("/get-memories/{event_id}/{page}", routes.GetMemories)
		memory.Post("/share-memory", routes.ShareMemory)
	}

	purchase := app.Party("/purchase")
	{
		purchase.Post("/add-purchase", routes.AddPurchase)
		purchase.Get("/get-purchases/{event_id}", routes.GetPurchases)
		purchase.Patch("/update-purchase/{purchase_id}", routes.UpdatePurchase)
		purchase.Delete("/delete-purchase/{purchase_id}", routes.DeletePurchase)
	}

	todo := app.Party("/todo")
	{
		todo.Post("/add-todo", routes.AddTodo)
		todo.Get("/get-todos/{event_id}", routes.GetTodos)
		todo.Patch("/update-todo/{todo_id}", routes.UpdateTodo)
		todo.Delete("/delete-todo/{todo_id}", routes.DeleteTodo)
	}

	user := app.Party("/user")
	{
		user.Post("/register", routes.Register)
		user.Post("/login", routes.Login)
		user.Post("/facebook", routes.FacebookLoginOrSignUp)
		user.Post("/google", routes.GoogleLoginOrSignUp)
		user.Post("/apple", routes.AppleLoginOrSignUp)
		user.Patch("/update/{user_id}", routes.UpdateUser)
	}

	app.Listen(":4000")
}
