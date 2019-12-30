package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	conn, err := ConnectToSql()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
}

func ConnectToSql() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:password@/phoneinfo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}
