package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          int        `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Date        time.Time  `json:"date"`
	Description string     `json:"description"`
	Img         string     `json:"img"`
	Location    string     `json:"location"`
	UserID      uint       `json:"userID"`
}
