package routes

import (
	"strconv"

	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/types"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/kataras/iris/v12"
)

// Get memories of paginated form
func GetMemories(ctx iris.Context) {
	params := ctx.Params()
	
	eventID := params.Get("event_id")
	page := params.Get("page")

	var eventInputID, pageOffset int
	var errConvert error

	eventInputID, errConvert = strconv.Atoi(eventID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error al recibir el parámetro event_id",
			ctx,
		)
	}

	pageOffset, errConvert = strconv.Atoi(page)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error al recibir el parámetro page",
			ctx,
		)
	}

	var memories []types.MemoryOutput

	query := /* sql */`
		SELECT id, title, picture, event_id
		FROM memories
		WHERE event_id = $1
		ORDER BY id DESC
		LIMIT $2
		OFFSET $3
	`

	rows, queryErr := storage.PostgresDB.Query(
		query, 
		eventInputID,
		20, 
		((pageOffset - 1) * 20),
	)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var memory types.MemoryOutput
		errRow := rows.Scan(
			&memory.ID, 
			&memory.Title,
			&memory.Picture,
			&memory.EventID,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		memories = append(memories, memory)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(memories) == 0 {
		ctx.JSON([]types.EventOutput{})
	} else {
		ctx.JSON(memories)
	}
}

// Get all memories
func GetAllMemories(ctx iris.Context) {
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


	var memories []types.MemoryOutput

	query := /* sql */`
		SELECT id, title, picture, event_id
		FROM memories
		WHERE event_id = $1
		ORDER BY id DESC
	`

	rows, queryErr := storage.PostgresDB.Query(
		query, 
		eventInputID,
	)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var memory types.MemoryOutput
		errRow := rows.Scan(
			&memory.ID, 
			&memory.Title,
			&memory.Picture,
			&memory.EventID,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		memories = append(memories, memory)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(memories) == 0 {
		ctx.JSON([]types.EventOutput{})
	} else {
		ctx.JSON(memories)
	}
}

// Share a memory
func ShareMemory(ctx iris.Context) {
	var memoryInput types.MemoryInput
	err := ctx.ReadJSON(&memoryInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var memory types.MemoryOutput

	query := /* sql */`
		INSERT INTO memories (title, picture, event_id)
		VALUES ($1, $2, $3)
		RETURNING id, title, picture, event_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query, 
		memoryInput.Title,
		memoryInput.Picture,
		memoryInput.EventID,
	).Scan(
		&memory.ID,
		&memory.Title,
		&memory.Picture,
		&memory.EventID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(memory)
}
