package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type NonNormailizedNo struct{
	Id       int
	PhoneNumber string
}

func GetNonNormalized(db *gorm.DB ){
	numbers:=[]NonNormailizedNo{}
	db.Find(&numbers)
	fmt.Println(numbers)
}