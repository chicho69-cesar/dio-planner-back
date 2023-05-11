package routes

import (
	"strings"

	"github.com/chicho69-cesar/dio-planner-back/models"
	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

/* ***** NOTE: Funciones para los endpoints de usuarios ***** */

func Register(ctx iris.Context) {
	var userInput RegisterUserInput
	err := ctx.ReadJSON(&userInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var newUser models.User
	userExists, userExistsErr := getAndHandleUserExists(&newUser, userInput.Email)
	if userExistsErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	if userExists == true {
		utils.CreateEmailAlreadyRegistered(ctx)
		return
	}

	hashedPassword, hashErr := hashAndSaltPassword(userInput.Password)
	if hashErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	newUser = models.User{
		Name:   		 userInput.Name,
		Email:       strings.ToLower(userInput.Email),
		Password:    hashedPassword,
		Description: "",
		Picture: 		 "https://dio-planner.s3.us-east-2.amazonaws.com/no-image.jpg",
		SocialLogin: false,
	}

	storage.DB.Create(&newUser)

	ctx.JSON(iris.Map{
		"ID":              newUser.ID,
		"name":       		 newUser.Name,
		"email":           newUser.Email,
		"description":     newUser.Description,
		"picture":         newUser.Picture,
	})
}

func Login(ctx iris.Context) {
	var userInput LoginUserInput
	err := ctx.ReadJSON(&userInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var existingUser models.User
	userExists, userExistsErr := getAndHandleUserExists(&existingUser, userInput.Email)
	if userExistsErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	errorMsg := "Invalid email or password."
	
	if userExists == false {
		utils.CreateError(iris.StatusUnauthorized, "Credentials Error", errorMsg, ctx)
		return
	}

	if existingUser.SocialLogin == true {
		utils.CreateError(iris.StatusUnauthorized, "Credentials Error", "Social Login Account", ctx)
		return
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userInput.Password))
	if passwordErr != nil {
		utils.CreateError(iris.StatusUnauthorized, "Credentials Error", errorMsg, ctx)
		return
	}

	ctx.JSON(iris.Map{
		"ID":              existingUser.ID,
		"name":       existingUser.Name,
		"email":           existingUser.Email,
		"description":     existingUser.Description,
		"picture":         existingUser.Picture,
	})
}

/* ***** NOTE: Funciones para funcionalidades extras ***** */

func getAndHandleUserExists(user *models.User, email string) (exists bool, err error) {
	userExistsQuery := storage.DB.Where("email = ?", strings.ToLower(email)).Limit(1).Find(&user)

	if userExistsQuery.Error != nil {
		return false, userExistsQuery.Error
	}

	userExists := userExistsQuery.RowsAffected > 0

	if userExists == true {
		return true, nil
	}

	return false, nil
}

func hashAndSaltPassword(password string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

/* ***** NOTE: Tipos para los endpoints ***** */

type RegisterUserInput struct {
	Name 			string 		`json:"name" validate:"required,max=256"`
	Email     string  	`json:"email" validate:"required,max=256,email"`
	Password  string  	`json:"password" validate:"required,min=8,max=256"`
}

type LoginUserInput struct {
	Email    	string 	`json:"email" validate:"required,email"`
	Password 	string 	`json:"password" validate:"required"`
}
