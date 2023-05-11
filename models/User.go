package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       				int    				 `json:"id" gorm:"primaryKey"`
	Name     				string 				 `json:"name"`
	Email    				string 				 `json:"email"`
	Password 				string 				 `json:"password"`
	Description 		string 				 `json:"description"`
	Picture 				string 				 `json:"picture"`
	SocialLogin     bool           `json:"socialLogin"`
	SocialProvider  string         `json:"socialProvider"`
	Grades 					[]Grade		 		 `json:"grades"`
}
