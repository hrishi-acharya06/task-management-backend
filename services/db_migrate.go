package services

import (
	models "task-management/models"
)

func MigrateDatabase() {
	var userTable models.User
	SERVICE_DB.AutoMigrate(&userTable)
}
