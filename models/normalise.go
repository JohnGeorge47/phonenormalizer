package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type nonNormailizedNo struct{
	ID       int		`gorm:"column:id"`
	PhoneNumber string `gorm:"column:phone_number"`
}

func GetNonNormalized(db *gorm.DB ){
	numbers:=nonNormailizedNo{}
	db.First(&numbers)
	db.SingularTable(true)

	fmt.Println(numbers)
}