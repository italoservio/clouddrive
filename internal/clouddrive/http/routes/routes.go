package routes

import (
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/http/controllers"
	mid "github.com/italoservio/clouddrive/internal/clouddrive/http/middlewares"
)

func SetupRoutes() {
	http.HandleFunc(
		"/authenticate",
		mid.ReqHandler(
			[]mid.Middleware{
				mid.JsonOut(),
				mid.JsonIn(),
				mid.Method(http.MethodPost),
			},
			http.HandlerFunc(controllers.Authenticate),
		),
	)
}
