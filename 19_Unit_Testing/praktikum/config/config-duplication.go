package config

import (
	"fmt"

	"db_API/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBTest() {
	const DB_USER_TEST = "root"
	const DB_PASS_TEST = ""
	const DB_HOST_TEST = "127.0.0.1"
	const DB_PORT_TEST = "3306"
	const DB_NAME_TEST = "api_test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})

	DB.Migrator().DropTable(&models.Blog{})
	DB.AutoMigrate(&models.Blog{})

	DB.Migrator().DropTable(&models.Book{})
	DB.AutoMigrate(&models.Book{})
}
