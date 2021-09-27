package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB       *gorm.DB
	username string = "root"
	password string = ""
	dbName   string = "test"
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	if err != nil {
		log.Fatalf("gorm open err : %v", err)
	}
	DB.SingularTable(true)
}
