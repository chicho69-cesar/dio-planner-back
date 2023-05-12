package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	ID      int    `json:"id" gorm:"primaryKey"`
	Opinion string `json:"opinion"`
	Grade   int    `json:"grade"`
	EventID uint   `json:"eventID"`
	UserID  uint   `json:"userID"`
}
