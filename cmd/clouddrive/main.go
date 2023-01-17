package main

import (
	"encoding/json"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	"github.com/italoservio/clouddrive/internal/clouddrive/middlewares"
)

func handle(wri http.ResponseWriter, req *http.Request) {
	user := entities.CreateUser("foo", "bar")
	json.NewEncoder(wri).Encode(user)
}

func main() {
	http.HandleFunc("/", middlewares.Http(http.MethodGet, http.HandlerFunc(handle)))
	http.ListenAndServe(":6000", nil)
}
