package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20)"`
	Email    string `gorm:"type:varchar(70);unique_index"`
	Password string `json:"-"`
}
