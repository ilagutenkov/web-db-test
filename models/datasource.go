package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB //база данных

func init() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort:=os.Getenv("db_port")



	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort) //Создать строку подключения
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println("error pg")
		fmt.Print(err)
	}

	db = conn

	db.SingularTable(true)
	db.Debug().AutoMigrate(&Person{},&Position{}) //Миграция базы данных
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}
