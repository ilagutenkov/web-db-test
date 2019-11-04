package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"web-db-test/models"
	u "web-db-test/utils"
)

func CreateRouting(router *mux.Router) {
	home("/", router)
	randomPerson("/random", router)
	putPosition("/position/put", router)
	getPosition("/position", router)
	sumAge("/sumAge", router)
	sumAgeParallel("/parallel/sumAge",router)

}

func sumAgeParallel(path string, router *mux.Router) {
router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
	age := models.SumAgeBeforeParallel(50)

	resp := u.Message(true, "success")
	resp["data"] = age

	u.Respond(writer, resp)

})

}

func sumAge(path string, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		age := models.SumAgeBefore(50)

		resp := u.Message(true, "success")
		resp["data"] = age

		u.Respond(writer, resp)

	})
}

func getPosition(path string, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		idStr := request.URL.Query()["id"][0]

		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			fmt.Print(err)
		}

		pos := models.GetPositionById(int(id))
		resp := u.Message(true, "success")
		resp["data"] = pos

		u.Respond(writer, resp)

	})
}

func putPosition(path string, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		positionName := request.URL.Query()["name"][0]
		money := request.URL.Query()["money"][0]
		m, err := strconv.ParseFloat(money, 32)
		if (err != nil) {
			fmt.Println(err)
		}

		pos := models.Position{Name: positionName, Money: float32(m)};
		models.PutPosition(&pos);
	})
}

func randomPerson(path string, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
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
}

func home(path string, router *mux.Router) *mux.Route {
	return router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
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
}
