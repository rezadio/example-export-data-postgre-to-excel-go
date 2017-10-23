package Datasources

import (
	"ReportingC2C/App/Datasources/Db"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

func init() {
	database = OpenConnectionDevel()
}

func GetConnectionDevel() *gorm.DB {
	if database == nil {
		database = OpenConnectionDevel()
	}
	return database
}

func OpenConnectionDevel() *gorm.DB {
	DB_HOST := Db.DB_HOST_DEVEL
	DB_USER := Db.DB_USER_DEVEL
	DB_NAME := Db.DB_NAME_DEVEL
	DB_PASSWORD := Db.DB_PASSWORD_DEVEL

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
	} else {
		fmt.Println("Connected To Devel " + DB_HOST)
	}
	return db.Debug()
}
