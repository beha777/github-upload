package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

/*func (user *User) TableName() string {
	return "user"
}*/
