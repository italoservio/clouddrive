package middlewares

import (
	"fmt"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/structs"
)

func Method(method string) Middleware {
	return func(wri http.ResponseWriter, req *http.Request) *structs.HttpError {
		if req.Method == http.MethodOptions {
			wri.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,PATCH,DELETE")
			wri.Header().Set("Access-Control-Allow-Headers", "*")
			wri.Header().Set("Access-Control-Allow-Origin", "*")
			wri.Header().Set("Access-Control-Max-Age", "86400")
			return structs.CreateHttpError(
				"",
				http.StatusNoContent,
				http.StatusText(http.StatusNoContent),
			)
		}

		if (method == http.MethodPost && req.Method != http.MethodPost) ||
			(method == http.MethodGet && req.Method != http.MethodGet) ||
			(method == http.MethodPut && req.Method != http.MethodPut) ||
			(method == http.MethodPatch && req.Method != http.MethodPatch) ||
			(method == http.MethodDelete && req.Method != http.MethodDelete) {

			return structs.CreateHttpError(
				fmt.Sprintf(
					"Method %v not allowed for this endpoint. Maybe you are looking for the allowed method: %v.",
					req.Method,
					method,
				),
				http.StatusNotImplemented,
				http.StatusText(http.StatusNotImplemented),
			)
		}

		return nil
	}
}
