package models

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	ID 			 int 		`json:"id" gorm:"primaryKey"`
	UserID 	 uint 	`json:"userID" gorm:"primaryKey"`
  EventID  uint 	`json:"eventID" gorm:"primaryKey"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
	Event    Event  `json:"event" gorm:"foreignKey:EventID"`
	Status 	 *bool  `json:"status"`
}
