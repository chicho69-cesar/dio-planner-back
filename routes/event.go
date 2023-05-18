package routes

import (
	"strconv"

	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/types"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/kataras/iris/v12"
)

// Add a event
func CreateEvent(ctx iris.Context) {
	var eventInput types.EventInput
	err := ctx.ReadJSON(&eventInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var event types.EventOutput

	query := `
		INSERT INTO events (name, date, description, img, location, user_id) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, date, description, img, location, user_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query,
		eventInput.Name,
		eventInput.Date,
		eventInput.Description,
		eventInput.Img,
		eventInput.Location,
		eventInput.UserID,
	).Scan(
		&event.ID, &event.Name, &event.Date, &event.Description,
		&event.Img, &event.Location, &event.UserID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(event)
}

// Get a event by id
func GetEventByID(ctx iris.Context) {
	params := ctx.Params()
	id := params.Get("id")

	var event types.EventOutput

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

	ctx.JSON(event)
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

	var events []types.EventOutput

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
		var event types.EventOutput
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
		ctx.JSON([]types.EventOutput{})
	} else {
		ctx.JSON(events)
	}
}

func GetTopEvents(ctx iris.Context) {
	var events []types.TopEventOutput

	query := `
		SELECT AVG(grade), events.id, events.name, events.date, events.description, events.img, events.location, events.user_id, events.accessibility
		FROM grades, events 
		WHERE grades.event_id = events.id AND events.accessibility = $1
		GROUP BY events.id 
		ORDER BY 1 DESC;
	`

	var accessibility string = "publico"

	rows, queryErr := storage.PostgresDB.Query(query, accessibility)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var event types.TopEventOutput
		errRow := rows.Scan(
			&event.AVG,
			&event.ID,
			&event.Name,
			&event.Date,
			&event.Description,
			&event.Img,
			&event.Location,
			&event.UserID,
			&event.Accessibility,
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
		ctx.JSON([]types.TopEventOutput{})
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

	var events []types.EventOutput

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
		var event types.EventOutput
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
		ctx.JSON([]types.EventOutput{})
	} else {
		ctx.JSON(events)
	}
}

// Get all events by query
func GetEventsByQuery(ctx iris.Context) {
	var queryName string = ctx.URLParam("name")
	var queryLocation string = ctx.URLParam("location")

	var events []types.SearchEventOutput

	query := `
		SELECT id, name, date, description, img, location, user_id, accessibility
		FROM events
		WHERE (name LIKE $1 OR location LIKE $2) AND accessibility = $3
	`

	var search1 string = "%" + queryName + "%"
	var search2 string = "%" + queryLocation + "%"
	var accessibility string = "publico"

	rows, queryErr := storage.PostgresDB.Query(
		query,
		search1,
		search2,
		accessibility,
	)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var event types.SearchEventOutput
		errRow := rows.Scan(
			&event.ID, &event.Name, &event.Date, &event.Description,
			&event.Img, &event.Location, &event.UserID, &event.Accessibility,
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
		ctx.JSON([]types.SearchEventOutput{})
	} else {
		ctx.JSON(events)
	}
}

// Update a event
func UpdateEvent(ctx iris.Context) {
	params := ctx.Params()
	eventID := params.Get("event_id")

	var eventInputID int
	eventInputID, errConvert := strconv.Atoi(eventID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	var eventInput types.EventInput
	err := ctx.ReadJSON(&eventInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var event types.EventOutput

	query := `
		UPDATE events
		SET name = $1, date = $2, description = $3, img = $4, location = $5, user_id = $6
		WHERE id = $7
		RETURNING id, name, date, description, img, location, user_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query,
		eventInput.Name,
		eventInput.Date,
		eventInput.Description,
		eventInput.Img,
		eventInput.Location,
		eventInput.UserID,
		eventInputID,
	).Scan(
		&event.ID, &event.Name, &event.Date, &event.Description,
		&event.Img, &event.Location, &event.UserID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(event)
}

// Delete a event
func DeleteEvent(ctx iris.Context) {
	params := ctx.Params()
	eventID := params.Get("event_id")

	var eventInputID int
	eventInputID, errConvert := strconv.Atoi(eventID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	query := `
		DELETE FROM events
		WHERE id = $1
	`

	_, queryErr := storage.PostgresDB.Exec(query, eventInputID)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"message": "Evento eliminado con éxito",
	})
}
