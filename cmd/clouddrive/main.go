package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	mid "github.com/italoservio/clouddrive/internal/clouddrive/middlewares/http"
)

func handle(wri http.ResponseWriter, req *http.Request) {
	user := entities.CreateUser("foo", "bar")

	response := mid.CreateHttpResponse(
		http.StatusOK,
		"",
		user,
	)

	obj, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	wri.Write(obj)
}

func main() {
	http.HandleFunc(
		"/",
		mid.ReqHandler(
			[]mid.Middleware{
				mid.JsonOut(),
				mid.JsonIn(),
				mid.Method(http.MethodGet),
			},
			http.HandlerFunc(handle),
		),
	)

	fmt.Println("Server listening at http://localhost:6000")
	http.ListenAndServe(":6000", nil)
}
