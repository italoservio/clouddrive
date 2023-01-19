package main

import (
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
	wri.Write(response.ToJson())
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
