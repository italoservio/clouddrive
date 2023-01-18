package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	"github.com/italoservio/clouddrive/internal/clouddrive/logger"
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

	logger.Info("Golang application is running and listening at http://localhost:6000")
	http.ListenAndServe(":6000", nil)
}
