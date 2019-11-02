package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	id int `json:"id"`
	name string `json:"name"`
	age int `json:"age"`
}

func GetFirstPerson() *Person{
	temp := &Person{};

	err:=GetDB().Table("people").First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		println("db eror")
	}

	return temp
}
