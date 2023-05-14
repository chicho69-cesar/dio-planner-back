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

/* ***** GUESTS ***** */
type GuestInput struct {
	UserID  uint  	`json:"userID"`
	EventID uint  	`json:"eventID"`
	Status  string  `json:"status"`
}

/* ***** GRADES ***** */
type GradeInput struct {
	Opinion string `json:"opinion"`
	Grade   int    `json:"grade"`
	EventID uint   `json:"eventID"`
	UserID  uint   `json:"userID"`
}

/* ***** MEMORIES ***** */
type MemoryInput struct {
	Title       string    `json:"title"`
	Picture     string    `json:"picture"`
	EventID 		uint 			`json:"eventID"`
}

/* ***** TODOS ***** */
type TodoInput struct {
	Text 				string 			`json:"text"`
	Date 				time.Time  	`json:"date"`
	Complete 		bool 				`json:"complete"`
	EventID 		uint 				`json:"eventID"`
}

/* ***** PURCHASES ***** */
type PurchaseInput struct {
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	EventID 		uint 			`json:"eventID"`
}
