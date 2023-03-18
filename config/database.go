package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "root"
	dbName   = "docker_db"
)

//var user = os.Getenv("DB_USER")
//var password = os.Getenv("DB_PASS")
//var host = os.Getenv("DB_HOST")
//var port = os.Getenv("DB_PORT")
//var dbName = os.Getenv("DB_NAME")

var DB *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		log.Fatal("Failed to connect to database")
	}
}
