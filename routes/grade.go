package routes

import (
	"strconv"

	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/types"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/kataras/iris/v12"
)

// Create a grade for an event with some user
func AddGrade(ctx iris.Context) {
	var gradeInput types.GradeInput
	err := ctx.ReadJSON(&gradeInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var grade types.GradeOutput

	query := /* sql */`
		INSERT INTO grades (opinion, grade, event_id, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, opinion, grade, event_id, user_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query, 
		gradeInput.Opinion,
		gradeInput.Grade,
		gradeInput.EventID,
		gradeInput.UserID,
	).Scan(
		&grade.ID,
		&grade.Opinion,
		&grade.Grade,
		&grade.EventID,
		&grade.UserID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	query = /* sql */`
		SELECT name FROM users WHERE id = $1
	`

	queryErr = storage.PostgresDB.
		QueryRow(query, grade.UserID).
		Scan(&grade.User)

	ctx.JSON(grade)
}

// Get all grades for an event
func GetGrades(ctx iris.Context) {
	params := ctx.Params()
	eventID := params.Get("event_id")

	var gradeInputID int
	gradeInputID, errConvert := strconv.Atoi(eventID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error al recibir el parámetro event_id",
			ctx,
		)
	}

	var grades []types.GradeOutput

	query := /* sql */`
		SELECT grades.id, opinion, grade, users.name, event_id, user_id
		FROM grades, users
		WHERE event_id = $1 AND users.id = user_id
		ORDER BY grades.id DESC
	`

	rows, queryErr := storage.PostgresDB.Query(query, gradeInputID)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var grade types.GradeOutput
		errRow := rows.Scan(
			&grade.ID, 
			&grade.Opinion,
			&grade.Grade,
			&grade.User,
			&grade.EventID,
			&grade.UserID,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		grades = append(grades, grade)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(grades) == 0 {
		ctx.JSON([]types.GradeOutput{})
	} else {
		ctx.JSON(grades)
	}
}
