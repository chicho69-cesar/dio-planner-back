package routes

import (
	"strconv"
	"time"

	"github.com/chicho69-cesar/dio-planner-back/models"
	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/kataras/iris/v12"
)

// Add a event
func CreateEvent(ctx iris.Context) {
	var eventQuery eventQuery
	err := ctx.ReadJSON(&eventQuery)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var event models.Event

	query := `
		INSERT INTO events (name, date, description, img, location, user_id) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, date, description, img, location, user_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query, 
		eventQuery.Name, 
		eventQuery.Date, 
		eventQuery.Description, 
		eventQuery.Img, 
		eventQuery.Location,
		eventQuery.UserID,
	).Scan(
		&event.ID, &event.Name, &event.Date, &event.Description, 
		&event.Img, &event.Location, &event.UserID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"ID":  				 event.ID,
		"name": 			 event.Name,
		"date": 			 event.Date,
		"description": event.Description,
		"img": 				 event.Img,
		"location": 	 event.Location,
		"user_id": 		 event.UserID,
	})
}

// Get a event by id
func GetEventByID(ctx iris.Context) {
	params := ctx.Params()
	id := params.Get("id")

	var event models.Event

	query := `
		SELECT id, name, date, description, img, location, user_id
		FROM events WHERE id = $1
	`

	queryErr := storage.PostgresDB.
		QueryRow(query, id).
		Scan(
			&event.ID, &event.Name, &event.Date, &event.Description, 
			&event.Img, &event.Location, &event.UserID,
		)

	if queryErr != nil {
		utils.CreateError(
			iris.StatusNotFound,
			"Elemento no encontrado",
			"Error al encontrar el evento con el id especificado",
			ctx,
		)
		return
	}

	ctx.JSON(iris.Map{
		"ID":  				 event.ID,
		"name": 			 event.Name,
		"date": 			 event.Date,
		"description": event.Description,
		"img": 				 event.Img,
		"location": 	 event.Location,
		"user_id": 		 event.UserID,
	})
}

// Get all events by pagination
func GetEvents(ctx iris.Context) {
	params := ctx.Params()
	page := params.Get("page")

	var pageOffset int
	pageOffset, errConvert := strconv.Atoi(page)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error al recibir el parámetro page",
			ctx,
		)
	}

	var events []EventOutput

	query := `
		SELECT id, name, date, description, img, location, user_id
		FROM events
		LIMIT $1
		OFFSET $2
	`

	rows, queryErr := storage.PostgresDB.Query(query, 20, ((pageOffset - 1) * 20))
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var event EventOutput
		errRow := rows.Scan(
			&event.ID, &event.Name, &event.Date, &event.Description, 
			&event.Img, &event.Location, &event.UserID,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		events = append(events, event)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(events) == 0 {
		ctx.JSON([]EventOutput{})
	} else {
		ctx.JSON(events)
	}
}

// Get all events by user
func GetEventsByUser(ctx iris.Context) {
	params := ctx.Params()
	userID := params.Get("user_id")

	var user int
	user, errConvert := strconv.Atoi(userID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	var events []EventOutput

	query := `
		SELECT id, name, date, description, img, location, user_id
		FROM events
		WHERE user_id = $1
	`

	rows, queryErr := storage.PostgresDB.Query(query, user)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var event EventOutput
		errRow := rows.Scan(
			&event.ID, &event.Name, &event.Date, &event.Description, 
			&event.Img, &event.Location, &event.UserID,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		events = append(events, event)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(events) == 0 {
		ctx.JSON([]EventOutput{})
	} else {
		ctx.JSON(events)
	}
}

// Get all events by query
func GetEventsByQuery(ctx iris.Context) {
	var eventQuery EventQuery
	err := ctx.ReadJSON(&eventQuery)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var events []EventOutput

	query := `
		SELECT id, name, date, description, img, location, user_id
		FROM events
		WHERE name LIKE $1
		OR description LIKE $2
		OR location LIKE $3
	`

	rows, queryErr := storage.PostgresDB.Query(
		query, 
		"%" + eventQuery.Name + "%",
		"%" + eventQuery.Description + "%",
		"%" + eventQuery.Location + "%",
	)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var event EventOutput
		errRow := rows.Scan(
			&event.ID, &event.Name, &event.Date, &event.Description, 
			&event.Img, &event.Location, &event.UserID,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		events = append(events, event)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(events) == 0 {
		ctx.JSON([]EventOutput{})
	} else {
		ctx.JSON(events)
	}
}

// Update a event
func UpdateEvent(ctx iris.Context) {
	// 
}

// Delete a event
func DeleteEvent(ctx iris.Context) {
	// 
}

type eventQuery struct {
	Name        string 		 	`json:"name"`
	Date 				time.Time  	`json:"date"`
	Description string 		 	`json:"description"`
	Img 				string 		 	`json:"img"`
	Location 		string 		 	`json:"location"`
	UserID      int    		 	`json:"user_id"`
}

type EventOutput struct {
	ID          int    		 	`json:"id" gorm:"primaryKey"`
	Name        string 		 	`json:"name"`
	Date 				time.Time  	`json:"date"`
	Description string 		 	`json:"description"`
	Img 				string 		 	`json:"img"`
	Location 		string 		 	`json:"location"`
	UserID      int    		 	`json:"user_id"`
}

type EventQuery struct {
	Name        string 		 	`json:"name"`
	Description string 		 	`json:"description"`
	Location 		string 		 	`json:"location"`
}
