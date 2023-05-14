package routes

import (
	"strconv"

	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/types"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/kataras/iris/v12"
)

// Add a todo to an Event
func AddTodo(ctx iris.Context) {
	var todoInput types.TodoInput
	err := ctx.ReadJSON(&todoInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var todo types.TodoOutput

	query := `
		INSERT INTO todos (text, date, complete, event_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, text, date, complete, event_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query, 
		todoInput.Text, 
		todoInput.Date, 
		todoInput.Complete, 
		todoInput.EventID,
	).Scan(
		&todo.ID, 
		&todo.Text,
		&todo.Date,
		&todo.Complete,
		&todo.EventID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(todo)
}

// Get todos of an Event
func GetTodos(ctx iris.Context) {
	params := ctx.Params()
	eventID := params.Get("event_id")

	var eventInputID int
	eventInputID, errConvert := strconv.Atoi(eventID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error al recibir el parámetro event_id",
			ctx,
		)
	}

	var todos []types.TodoOutput

	query := `
		SELECT id, text, date, complete, event_id
		FROM todos
		WHERE event_id = $1
		ORDER BY date, id DESC
	`

	rows, queryErr := storage.PostgresDB.Query(query, eventInputID)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo types.TodoOutput
		errRow := rows.Scan(
			&todo.ID, 
			&todo.Text,
			&todo.Date,
			&todo.Complete,
			&todo.EventID,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		todos = append(todos, todo)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(todos) == 0 {
		ctx.JSON([]types.TodoOutput{})
	} else {
		ctx.JSON(todos)
	}
}

// Update a todo
func UpdateTodo(ctx iris.Context) {
	params := ctx.Params()
	todoID := params.Get("todo_id")

	var todoInputID int
	todoInputID, errConvert := strconv.Atoi(todoID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	var todoInput types.TodoInput
	err := ctx.ReadJSON(&todoInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var todo types.TodoOutput

	query := `
		UPDATE todos
		SET text = $1, date = $2, complete = $3
		WHERE id = $4
		RETURNING id, text, date, complete, event_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query, 
		todoInput.Text, 
		todoInput.Date, 
		todoInput.Complete, 
		todoInputID,
	).Scan(
		&todo.ID, 
		&todo.Text,
		&todo.Date,
		&todo.Complete,
		&todo.EventID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(todo)
}

// Delete a todo
func DeleteTodo(ctx iris.Context) {
	params := ctx.Params()
	todoID := params.Get("todo_id")

	var todoInputID int
	todoInputID, errConvert := strconv.Atoi(todoID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	query := `
		DELETE FROM todos
		WHERE id = $1
	`

	_, queryErr := storage.PostgresDB.Exec(query, todoInputID)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"message": "Evento eliminado con éxito",
	})
}
