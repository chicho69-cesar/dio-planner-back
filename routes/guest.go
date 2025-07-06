package routes

import (
	"strconv"

	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/types"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/kataras/iris/v12"
)

// Send a invitation to an user
func AddGuest(ctx iris.Context) {
	var guestInput types.GuestInput
	err := ctx.ReadJSON(&guestInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var guest types.GuestResponse

	query := /* sql */`
		INSERT INTO guests (user_id, event_id, status)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, event_id, status;
	`

	queryErr := storage.PostgresDB.QueryRow(
		query,
		guestInput.UserID,
		guestInput.EventID,
		guestInput.Status,
	).Scan(
		&guest.ID,
		&guest.UserID,
		&guest.EventID,
		&guest.Status,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(guest)
}

// Search user to send invitations to them
func SearchGuests(ctx iris.Context) {
	var searchQuery string = ctx.URLParam("query")

	var guests []types.UserOutput

	query := /* sql */`
		SELECT id, name, description, picture
		FROM users
		WHERE name LIKE $1
	`

	rows, queryErr := storage.PostgresDB.Query(
		query,
		"%"+searchQuery+"%",
	)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var guest types.UserOutput
		errRow := rows.Scan(
			&guest.ID,
			&guest.Name,
			&guest.Description,
			&guest.Picture,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		guests = append(guests, guest)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(guests) == 0 {
		ctx.JSON([]types.UserOutput{})
	} else {
		ctx.JSON(guests)
	}
}

// Get all guests of an event
func GetGuests(ctx iris.Context) {
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

	var guests []types.GuestsOutput

	query := /* sql */`
		SELECT guests.id, name, description, picture, status 
		FROM guests, users 
		WHERE event_id = $1 AND user_id=users.id
	`

	rows, queryErr := storage.PostgresDB.Query(query, eventInputID)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var guest types.GuestsOutput
		errRow := rows.Scan(
			&guest.ID,
			&guest.Name,
			&guest.Description,
			&guest.Picture,
			&guest.Status,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		guests = append(guests, guest)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(guests) == 0 {
		ctx.JSON([]types.GuestsOutput{})
	} else {
		ctx.JSON(guests)
	}
}

// Get all invitations of an user
func GetInvitations(ctx iris.Context) {
	params := ctx.Params()
	userID := params.Get("user_id")

	var userInputID int
	userInputID, errConvert := strconv.Atoi(userID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	var invitations []types.InvitationsOutput

	actualStatus := "Pendiente"
	query := /* sql */`
		SELECT guests.id, name, date, description, img 
		FROM guests, events 
		WHERE guests.user_id = $1 AND guests.event_id = events.id AND status= $2
	`

	rows, queryErr := storage.PostgresDB.Query(query, userInputID, actualStatus)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var invitation types.InvitationsOutput
		errRow := rows.Scan(
			&invitation.ID,
			&invitation.Name,
			&invitation.Date,
			&invitation.Description,
			&invitation.Img,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		invitations = append(invitations, invitation)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(invitations) == 0 {
		ctx.JSON([]types.InvitationsOutput{})
	} else {
		ctx.JSON(invitations)
	}
}

// Accept an invitation
func AcceptInvitation(ctx iris.Context) {
	params := ctx.Params()
	guestID := params.Get("guest_id")

	var guestInputID int
	guestInputID, errConvert := strconv.Atoi(guestID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	query := /* sql */`
		UPDATE guests
		SET status = 'Aceptada'
		WHERE id = $1
	`

	statement, errStmt := storage.PostgresDB.Prepare(query)
	if errStmt != nil {
		utils.CreateQueryError(ctx)
		return
	}

	_, errExec := statement.Exec(guestInputID)
	if errExec != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"message": "Invitación aceptada con éxito",
	})
}

// Decline an invitation
func DeclineInvitation(ctx iris.Context) {
	params := ctx.Params()
	guestID := params.Get("guest_id")

	var guestInputID int
	guestInputID, errConvert := strconv.Atoi(guestID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	query := /* sql */`
		UPDATE guests
		SET status = 'Cancelada'
		WHERE id = $1
	`

	statement, errStmt := storage.PostgresDB.Prepare(query)
	if errStmt != nil {
		utils.CreateQueryError(ctx)
		return
	}

	_, errExec := statement.Exec(guestInputID)
	if errExec != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"message": "Invitación rechazada con éxito",
	})
}

// Get my events and the events in where I'm invited
func GetMyEvents(ctx iris.Context) {
	params := ctx.Params()
	userID := params.Get("user_id")

	var userInputID int
	userInputID, errConvert := strconv.Atoi(userID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	var events []types.EventOutput

	actualStatus := "Aceptada"
	query := /* sql */`
		SELECT DISTINCT ON (events.id)
			events.id, events.name, events.date, events.description, 
			events.img, events.location, events.topic, events.user_id 
		FROM guests, events 
		WHERE 
			((events.user_id = $1) OR (guests.user_id = $2 AND guests.status = $3)) 
			AND (guests.event_id = events.id)
	`

	rows, queryErr := storage.PostgresDB.Query(
		query,
		userInputID,
		userInputID,
		actualStatus,
	)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var event types.EventOutput
		errRow := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Date,
			&event.Description,
			&event.Img,
			&event.Location,
			&event.Topic,
			&event.UserID,
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
