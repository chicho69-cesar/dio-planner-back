package models

import "gorm.io/gorm"

type Memory struct {
	gorm.Model
	ID          int    		`json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Picture     string    `json:"picture"`
	EventID 		int 			`json:"eventID" gorm:"primaryKey"`
	Event 			Event 		`json:"event" gorm:"foreignKey:EventID"`
}
