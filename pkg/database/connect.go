package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	User     string `default:"root"`
	Password string `default:"root"`
	Host     string `default:"127.0.0.1"`
	Port     string `default:"3306"`
	DBName   string `default:"test"`
}

var Db *gorm.DB
var DbErr error

func InitDatabase(config DatabaseConfig) (*gorm.DB, error) {
	// Construct DSN
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	Db = db
	DbErr = err

	return db, err
}
