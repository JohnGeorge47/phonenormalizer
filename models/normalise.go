package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type nonNormalizedNo struct{
	ID       int		`gorm:"column:id"`
	PhoneNumber string `gorm:"column:phone_number"`
}

func GetNonNormalized(db *gorm.DB )*[]nonNormalizedNo{
	numbers:=[]nonNormalizedNo{}
	fmt.Println("here")
	db.Debug().Find(&numbers)
	db.SingularTable(true)
	fmt.Println(numbers)
	return &numbers
}