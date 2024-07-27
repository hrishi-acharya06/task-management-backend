package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
}

func (User) TableName() string {
	return "tabUser"
}
