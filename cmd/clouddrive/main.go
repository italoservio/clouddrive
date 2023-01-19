package main

import (
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/http/routes"
	"github.com/italoservio/clouddrive/pkg/logger"
)

func main() {
	routes.SetupRoutes()
	logger.Info("Golang application is running and listening at http://localhost:6000")
	http.ListenAndServe(":6000", nil)
}
