package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/italoservio/clouddrive/internal/clouddrive/db"
	"github.com/italoservio/clouddrive/internal/clouddrive/http/routes"
	"github.com/italoservio/clouddrive/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	env_filename := path.Join(pwd, "cmd", "clouddrive", ".env")
	fmt.Println(env_filename)
	if err := godotenv.Load(env_filename); err != nil {
		panic(err)
	}

	db.SetupConnection()
	routes.SetupRoutes()

	logger.Info("Golang application is running and listening at http://localhost:6000")
	http.ListenAndServe(":6000", nil)
}
