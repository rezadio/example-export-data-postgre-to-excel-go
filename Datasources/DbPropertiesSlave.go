package Datasources

import (
	"ReportingEmoney/App/Datasources/Db"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

func init() {
	database = OpenConnectionProd()
}

func GetConnectionProd() *gorm.DB {
	if database == nil {
		database = OpenConnectionProd()
	}
	return database
}

func OpenConnectionProd() *gorm.DB {
	DB_HOST := Db.DB_HOST_PROD
	DB_USER := Db.DB_USER_PROD
	DB_NAME := Db.DB_NAME_PROD
	DB_PASSWORD := Db.DB_PASSWORD_PROD

	dbinfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		DB_HOST, DB_USER, DB_NAME, DB_PASSWORD)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	} else {
		db.DB().SetMaxOpenConns(100)
		db.DB().SetMaxIdleConns(1)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected To " + DB_HOST)
	return db
}
