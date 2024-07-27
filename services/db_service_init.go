package services

import (
	dbb "task-management/db"

	"gorm.io/gorm"
)

var SERVICE_DB *gorm.DB

func InitializeDB() {
	SERVICE_DB = dbb.GetDB()
}
