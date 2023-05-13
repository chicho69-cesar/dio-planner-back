package types

import "time"

/* ***** USER ***** */
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

type UserUpdate struct {
	Name     				string 				 `json:"name"`
	Password 				string 				 `json:"password"`
	Description 		string 				 `json:"description"`
	Picture 				string 				 `json:"picture"`
}

/* ***** EVENTS *****  */
type EventInput struct {
	Name        string 		 	`json:"name"`
	Date 				time.Time  	`json:"date"`
	Description string 		 	`json:"description"`
	Img 				string 		 	`json:"img"`
	Location 		string 		 	`json:"location"`
	UserID      int    		 	`json:"user_id"`
}
