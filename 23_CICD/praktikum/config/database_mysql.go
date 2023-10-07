package config

import (
	"clean/api/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlConn(config *AppConfig) *gorm.DB {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUSER, config.DBPASS, config.DBHOST, config.DBPORT, config.DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Model : cannot connect to database, ", err.Error())
		return nil

	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}



