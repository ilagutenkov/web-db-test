package main

import (
	"./controller"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	fmt.Println("...starting...")

	router := mux.NewRouter()

	controller.CreateRouting(router)

	//router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена

	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
