package task_management_db

import (
	"fmt"
	"os"
	"task-management/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	Schema   string
	User     string
	Password string
}

func (dc *DBConfig) getDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.User,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.Schema,
	)
}

func GetDB() *gorm.DB {
	dbConfig := &DBConfig{
		Host:     os.Getenv(constants.DbHost),
		Port:     os.Getenv(constants.DbPort),
		Schema:   os.Getenv(constants.DbName),
		User:     os.Getenv(constants.DbUser),
		Password: os.Getenv(constants.DbPassword),
	}
	dsn := dbConfig.getDSN()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in creating db client: ", err)
	}

	return db
}
