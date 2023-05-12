package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

func CreateError(statusCode int, title string, detail string, ctx iris.Context) {
	ctx.StopWithProblem(
		statusCode, 
		iris.NewProblem().
			Title(title).
			Detail(detail),
	)
}

func CreateInternalServerError(ctx iris.Context) {
	CreateError(
		iris.StatusInternalServerError,
		"Internal Server Error",
		"Internal Server Error",
		ctx,
	)
}

func CreateEmailAlreadyRegistered(ctx iris.Context) {
	CreateError(
		iris.StatusConflict,
		"Conflict",
		"El Email ya existe en la base de datos.",
		ctx,
	)
}

func CreateQueryError(ctx iris.Context) {
	CreateError(
		iris.StatusBadRequest,
		"Bad Request",
		"Error al realizar la operación.",
		ctx,
	)
}

func HandleValidationErrors(err error, ctx iris.Context) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		validationErrors := wrapValidationErrors(errs)

		fmt.Println("validationErrors", validationErrors)
		ctx.StopWithProblem(
			iris.StatusBadRequest,
			iris.NewProblem().
				Title("Validation error").
				Detail("Uno o mas campos fallaron en las validaciones.").
				Key("errors", validationErrors),
		)

		return
	}

	fmt.Print(err)

	CreateInternalServerError(ctx)
}

func wrapValidationErrors(errs validator.ValidationErrors) []validationError {
	validationErrors := make([]validationError, 0, len(errs))

	for _, validationErr := range errs {
		validationErrors = append(validationErrors, validationError{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Kind:      validationErr.Kind().String(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})
	}

	return validationErrors
}

type validationError struct {
	ActualTag 	string 	`json:"tag"`
	Namespace 	string 	`json:"namespace"`
	Kind      	string 	`json:"kind"`
	Type      	string 	`json:"type"`
	Value     	string 	`json:"value"`
	Param     	string 	`json:"param"`
}
