package config

import (
	"github.com/jinzhu/gorm"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Auth/structs"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.User{})
	return db
}
