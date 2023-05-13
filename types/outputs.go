package types

import "time"

/* ***** EVENTS ***** */
type EventOutput struct {
	ID          int    		 	`json:"id" gorm:"primaryKey"`
	Name        string 		 	`json:"name"`
	Date 				time.Time  	`json:"date"`
	Description string 		 	`json:"description"`
	Img 				string 		 	`json:"img"`
	Location 		string 		 	`json:"location"`
	UserID      int    		 	`json:"user_id"`
}
