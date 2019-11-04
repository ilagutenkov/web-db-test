package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Position struct {
	gorm.Model
	Name string
	Money float32
}

func PutPosition(pos *Position){
	GetDB().Debug().Create(pos)
}

func GetPositionById(id int) *Position{

	pos:=&Position{}

	err	:=GetDB().Where("id = ?",id).Find(pos).Error;

	if(err!=nil){
		fmt.Println(err)
	}

	return pos;
}