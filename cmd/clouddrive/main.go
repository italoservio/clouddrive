package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	mid "github.com/italoservio/clouddrive/internal/clouddrive/middlewares/http"
)

func handle(wri http.ResponseWriter, req *http.Request) {
	user := entities.CreateUser("foo", "bar")
	json.NewEncoder(wri).Encode(user)
}

func main() {
	http.HandleFunc(
		"/",
		mid.Intercept(
			[]mid.Middleware{
				mid.Json(),
				mid.Method(http.MethodGet),
			},
			http.HandlerFunc(handle),
		),
	)

	fmt.Println("Server listening at http://localhost:6000")
	http.ListenAndServe(":6000", nil)
}
