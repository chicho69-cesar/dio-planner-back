package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          int    			`json:"id" gorm:"primaryKey"`
	Text 				string 			`json:"text"`
	Date 				time.Time  	`json:"date"`
	Complete 		bool 				`json:"complete"`
	EventID 		int 				`json:"eventID" gorm:"primaryKey"`
	Event 			Event 			`json:"event" gorm:"foreignKey:EventID"`
}
