package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/chicho69-cesar/dio-planner-back/models"
	"github.com/chicho69-cesar/dio-planner-back/storage"
	"github.com/chicho69-cesar/dio-planner-back/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

// Register a user with email and password
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

// Login with user and password
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

// Sign in and Sign up with Facebook
func FacebookLoginOrSignUp(ctx iris.Context) {
	var userInput FacebookOrGoogleUserInput
	err := ctx.ReadJSON(&userInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	endpoint := "https://graph.facebook.com/me?fields=id,name,email&access_token=" + userInput.AccessToken
	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	res, facebookErr := client.Do(req)
	if facebookErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Panic(bodyErr)
		utils.CreateInternalServerError(ctx)
		return
	}

	var facebookBody FacebookUserRes
	json.Unmarshal(body, &facebookBody)

	if facebookBody.Email != "" {
		var user models.User
		userExists, userExistsErr := getAndHandleUserExists(&user, facebookBody.Email)

		if userExistsErr != nil {
			utils.CreateInternalServerError(ctx)
			return
		}

		if userExists == false {
			nameArr := strings.SplitN(facebookBody.Name, " ", 2)
			user = models.User{
				Name: 			 		nameArr[0] + " " + nameArr[1], 
				Email: 					facebookBody.Email, 
				Description: 		"",
				Picture: 				"https://dio-planner.s3.us-east-2.amazonaws.com/no-image.jpg",
				SocialLogin: 		true, 
				SocialProvider: "Facebook",
			}
			storage.DB.Create(&user)

			ctx.JSON(iris.Map{
				"ID":              user.ID,
				"name":       		 user.Name,
				"email":           user.Email,
				"description":     user.Description,
				"picture":         user.Picture,
			})

			return
		}

		if user.SocialLogin == true && user.SocialProvider == "Facebook" {
			ctx.JSON(iris.Map{
				"ID":              user.ID,
				"name":       		 user.Name,
				"email":           user.Email,
				"description":     user.Description,
				"picture":         user.Picture,
			})

			return
		}

		utils.CreateEmailAlreadyRegistered(ctx)

		return
	}
}

// Sign in and Sign up with Google
func GoogleLoginOrSignUp(ctx iris.Context) {
	var userInput FacebookOrGoogleUserInput
	err := ctx.ReadJSON(&userInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	endpoint := "https://www.googleapis.com/userinfo/v2/me"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	header := "Bearer " + userInput.AccessToken
	req.Header.Set("Authorization", header)
	res, googleErr := client.Do(req)
	if googleErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Panic(bodyErr)
		utils.CreateInternalServerError(ctx)
		return
	}

	var googleBody GoogleUserRes
	json.Unmarshal(body, &googleBody)

	if googleBody.Email != "" {
		var user models.User
		userExists, userExistsErr := getAndHandleUserExists(&user, googleBody.Email)

		if userExistsErr != nil {
			utils.CreateInternalServerError(ctx)
			return
		}

		if userExists == false {
			user = models.User{
				Name:  					googleBody.GivenName + " " + googleBody.FamilyName, 
				Email: 					googleBody.Email, 
				Description: 		"",
				Picture: 				"https://dio-planner.s3.us-east-2.amazonaws.com/no-image.jpg",
				SocialLogin: 		true, 
				SocialProvider: "Google",
			}
			storage.DB.Create(&user)

			ctx.JSON(iris.Map{
				"ID":              user.ID,
				"name":       		 user.Name,
				"email":           user.Email,
				"description":     user.Description,
				"picture":         user.Picture,
			})

			return
		}

		if user.SocialLogin == true && user.SocialProvider == "Google" {
			ctx.JSON(iris.Map{
				"ID":              user.ID,
				"name":       		 user.Name,
				"email":           user.Email,
				"description":     user.Description,
				"picture":         user.Picture,
			})

			return
		}

		utils.CreateEmailAlreadyRegistered(ctx)
		
		return
	}
}

// Sign in and Sign up with Apple
func AppleLoginOrSignUp(ctx iris.Context) {
	var userInput AppleUserInput
	err := ctx.ReadJSON(&userInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	res, httpErr := http.Get("https://appleid.apple.com/auth/keys")
	if httpErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	defer res.Body.Close()

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	jwks, jwksErr := keyfunc.NewJSON(body)
	token, tokenErr := jwt.Parse(userInput.IdentityToken, jwks.Keyfunc)

	if jwksErr != nil || tokenErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	if !token.Valid {
		utils.CreateError(iris.StatusUnauthorized, "Unauthorized", "Invalid user token.", ctx)
		return
	}

	email := fmt.Sprint(token.Claims.(jwt.MapClaims)["email"])
	if email != "" {
		var user models.User
		userExists, userExistsErr := getAndHandleUserExists(&user, email)

		if userExistsErr != nil {
			utils.CreateInternalServerError(ctx)
			return
		}

		if userExists == false {
			user = models.User{
				Name: "", 
				Email: email, 
				SocialLogin: true, 
				SocialProvider: "Apple",
			}
			storage.DB.Create(&user)

			ctx.JSON(iris.Map{
				"ID":              user.ID,
				"name":       		 user.Name,
				"email":           user.Email,
				"description":     user.Description,
				"picture":         user.Picture,
			})
			
			return
		}

		if user.SocialLogin == true && user.SocialProvider == "Apple" {
			ctx.JSON(iris.Map{
				"ID":              user.ID,
				"name":       		 user.Name,
				"email":           user.Email,
				"description":     user.Description,
				"picture":         user.Picture,
			})

			return
		}

		utils.CreateEmailAlreadyRegistered(ctx)

		return
	}
}

// Update the information of an user
func UpdateUser(ctx iris.Context) {
	params := ctx.Params()
	userID := params.Get("user_id")

	var userUpdateInput UserUpdate
	err := ctx.ReadJSON(&userUpdateInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	hashedPassword, hashErr := hashAndSaltPassword(userUpdateInput.Password)
	if hashErr != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	var user models.User
	userExists := storage.DB.
		Where("id = ?", userID).
		Find(&user)

	if userExists.Error != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	if userExists.RowsAffected == 0 {
		utils.CreateError(
			iris.StatusNotFound, 
			"Not Found", 
			"User not found", 
			ctx,
		)
		return
	}

	userUpdate := UserUpdate {
		Name: userUpdateInput.Name,
		Password: hashedPassword,
		Description: userUpdateInput.Description,
		Picture: userUpdateInput.Picture,
	}

	rowsUpdated := storage.DB.
		Model(&user).
		Updates(userUpdate)

	if rowsUpdated.Error != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"message": "Usuario actualizado exitosamente",
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

type RegisterUserInput struct {
	Name 			string 		`json:"name" validate:"required,max=256"`
	Email     string  	`json:"email" validate:"required,max=256,email"`
	Password  string  	`json:"password" validate:"required,min=8,max=256"`
}

type LoginUserInput struct {
	Email    	string 	`json:"email" validate:"required,email"`
	Password 	string 	`json:"password" validate:"required"`
}

type FacebookOrGoogleUserInput struct {
	AccessToken string `json:"accessToken" validate:"required"`
}

type AppleUserInput struct {
	IdentityToken string `json:"identityToken" validate:"required"`
}

type FacebookUserRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GoogleUserRes struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
}

type UserUpdate struct {
	Name     				string 				 `json:"name"`
	Password 				string 				 `json:"password"`
	Description 		string 				 `json:"description"`
	Picture 				string 				 `json:"picture"`
}
