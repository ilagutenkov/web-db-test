package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"web-db-test/models"
	u "web-db-test/utils"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set(
			"Content-Type",
			"text/html",
		)
		io.WriteString(
			writer,
			`<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
		)
	})

	router.HandleFunc("/random", func(writer http.ResponseWriter, request *http.Request) {
		person := models.GetFirstPerson()

		//err := json.NewDecoder(request.Body).Decode(person)
		//if err != nil {
		//	u.Respond(writer, u.Message(false, "Error while decoding request body"))
		//	return
		//}
		resp := u.Message(true, "success")
		resp["data"] = person

		u.Respond(writer, resp)

	})

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
