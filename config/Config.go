package config

import (
	"level_5/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var Database_URI = "root@tcp(localhost:3306)/office?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {

	var err error

	Database, err = gorm.Open(mysql.Open(Database_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic("Database tidak tehubung")
	}

	Database.AutoMigrate(&model.Department{})
	Database.AutoMigrate(&model.Employee{})

	return nil

}
