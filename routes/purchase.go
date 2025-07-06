package routes

import (
	"strconv"

	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/types"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/kataras/iris/v12"
)

// Add a purchase to an Event
func AddPurchase(ctx iris.Context) {
	var purchaseInput types.PurchaseInput
	err := ctx.ReadJSON(&purchaseInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var purchase types.PurchaseOutput

	query := /* sql */`
		INSERT INTO purchases (title, price, event_id)
		VALUES ($1, $2, $3)
		RETURNING id, title, price, event_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query, 
		purchaseInput.Title, 
		purchaseInput.Price, 
		purchaseInput.EventID,
	).Scan(
		&purchase.ID, 
		&purchase.Title,
		&purchase.Price,
		&purchase.EventID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(purchase)
}

// Get purchases of an Event
func GetPurchases(ctx iris.Context) {
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

	var purchases []types.PurchaseOutput

	query := /* sql */`
		SELECT id, title, price, event_id
		FROM purchases
		WHERE event_id = $1
		ORDER BY id DESC
	`

	rows, queryErr := storage.PostgresDB.Query(query, eventInputID)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var purchase types.PurchaseOutput
		errRow := rows.Scan(
			&purchase.ID, 
			&purchase.Title,
			&purchase.Price,
			&purchase.EventID,
		)

		if errRow != nil {
			utils.CreateQueryError(ctx)
			return
		}

		purchases = append(purchases, purchase)
	}

	if errRead := rows.Err(); errRead != nil {
		utils.CreateQueryError(ctx)
		return
	}

	if len(purchases) == 0 {
		ctx.JSON([]types.TodoOutput{})
	} else {
		ctx.JSON(purchases)
	}
}

// Update a purchase
func UpdatePurchase(ctx iris.Context) {
	params := ctx.Params()
	purchaseID := params.Get("purchase_id")

	var purchaseInputID int
	purchaseInputID, errConvert := strconv.Atoi(purchaseID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	var purchaseInput types.PurchaseInput
	err := ctx.ReadJSON(&purchaseInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var purchase types.PurchaseOutput

	query := /* sql */`
		UPDATE purchases
		SET title = $1, price = $2
		WHERE id = $3
		RETURNING id, title, price, event_id
	`

	queryErr := storage.PostgresDB.QueryRow(
		query, 
		purchaseInput.Title, 
		purchaseInput.Price, 
		purchaseInputID,
	).Scan(
		&purchase.ID, 
		&purchase.Title,
		&purchase.Price,
		&purchase.EventID,
	)

	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(purchase)
}

// Delete a purchase
func DeletePurchase(ctx iris.Context) {
	params := ctx.Params()
	purchaseID := params.Get("purchase_id")

	var purchaseInputID int
	purchaseInputID, errConvert := strconv.Atoi(purchaseID)
	if errConvert != nil {
		utils.CreateError(
			iris.StatusBadRequest,
			"Error",
			"Error ID invalido",
			ctx,
		)
	}

	query := /* sql */`
		DELETE FROM purchases
		WHERE id = $1
	`

	_, queryErr := storage.PostgresDB.Exec(query, purchaseInputID)
	if queryErr != nil {
		utils.CreateQueryError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"message": "Evento eliminado con éxito",
	})
}
