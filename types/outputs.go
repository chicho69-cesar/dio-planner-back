package types

import "time"

/* ***** USERS ***** */
type UserOutput struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

/* ***** EVENTS ***** */
type EventOutput struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Img         string    `json:"img"`
	Location    string    `json:"location"`
	UserID      int       `json:"user_id"`
}

/* ***** GUESTS ***** */
type GuestsOutput struct {
	ID          int       `json:"id"`
	Name        string 		`json:"name"`
	Description string 		`json:"description"`
	Picture     string 		`json:"picture"`
	Status      string 		`json:"status"`
}

type InvitationsOutput struct {
	ID          int       `json:"id"`
	Name        string 		`json:"name"`
	Date        time.Time `json:"date"`
	Description string 		`json:"description"`
	Img     		string 		`json:"img"`
}

/* ***** GRADES ***** */
type GradeOutput struct {
	ID      int    `json:"id"`
	Opinion string `json:"opinion"`
	Grade   int    `json:"grade"`
	EventID uint   `json:"eventID"`
	UserID  uint   `json:"userID"`
}
