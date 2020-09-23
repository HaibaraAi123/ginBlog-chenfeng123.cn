package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var db *gorm.DB
type user struct {
	id       int
	role     string
	name     string
	password string
}

func main() {
	var err error
	db, err := gorm.Open("mysql", "root:HaibaraAi0715@/chenfeng123?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&user{})
	user1 := user{123, "nihoa ", "adsfd", "sdfasdf "}
	db.NewRecord(user1)
	db.Create(&user1)
}
