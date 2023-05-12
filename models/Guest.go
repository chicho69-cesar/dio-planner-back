package models

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	ID      int   `json:"id" gorm:"primaryKey"`
	UserID  uint  `json:"userID"`
	EventID uint  `json:"eventID"`
	Status  bool  `json:"status"`
}
