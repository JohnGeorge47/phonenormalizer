package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type NonNormalizedNo struct{
	ID       int		`gorm:"column:id"`
	PhoneNumber string `gorm:"column:phone_number"`
}

type NormalizedNo struct {
	ID int `gorm:"column:id""`
	Phonenumber string `gorm:"column:phone_number"`
}

func GetNonNormalized(db *gorm.DB )*[]NonNormalizedNo{
	numbers:=[]NonNormalizedNo{}
	db.Debug().Find(&numbers)
	db.SingularTable(true)
	return &numbers
}

func InsertNormalized(db *gorm.DB,normalizedNumbers []string){
	for _,num:=range normalizedNumbers{
		phonenumber:=NormalizedNo{Phonenumber:num}
		db.Debug().Create(&phonenumber)
	}
}