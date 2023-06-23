package config

import (
	"fmt"
	"log"

	"github.com/pwn233/golang-short-course/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

var DevSandbox Store

func InitializeDatabaseConnection() {
	var ConnectionMasterDB string
	ConnectionMasterDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		common.LocalDatabaseUsername,
		common.LocalDatabasePassword,
		common.LocalDatabaseHost,
		common.LocalDatabaseMacPort,
		common.LocalDatabaseName)

	db, err := gorm.Open(mysql.Open(ConnectionMasterDB), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("[Database] Failed to connect database : %s\n", err.Error()))
	}
	fmt.Printf("[Database] Successfully connected at : %s\n", ConnectionMasterDB)
	DevSandbox.DB = db
}
