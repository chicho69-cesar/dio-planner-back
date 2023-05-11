package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	ID          int    		`json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	EventID 		int 			`json:"eventID" gorm:"primaryKey"`
	Event 			Event 		`json:"event" gorm:"foreignKey:EventID"`
}
