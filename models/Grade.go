package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	ID          int    		`json:"id" gorm:"primaryKey"`
	Opinion 	  string		`json:"opinion"`
	Grade 	    int				`json:"grade"`
	EventID 		int 			`json:"eventID" gorm:"primaryKey"`
	Event 			Event 		`json:"event" gorm:"foreignKey:EventID"`
	UserID 	 		uint 			`json:"userID" gorm:"primaryKey"`
	User     		User   		`json:"user" gorm:"foreignKey:UserID"`
}
