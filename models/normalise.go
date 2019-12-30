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

func GetNonNormalized(db *gorm.DB ){
	numbers:=nonNormalizedNo{}
	fmt.Println("here")
	err:=db.Debug().Find(&numbers)
	if err!=nil{
		fmt.Println(err)
	}
	db.SingularTable(true)

	fmt.Println(numbers)
}