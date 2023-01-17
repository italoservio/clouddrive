package middlewares

import (
	"fmt"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/structs"
)

func HttpMethod(method string, w http.ResponseWriter, r *http.Request) *structs.HttpError {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
		return structs.CreateHttpError(
			"",
			http.StatusNoContent,
			http.StatusText(http.StatusNoContent),
		)
	}

	if (method == http.MethodPost && r.Method != http.MethodPost) ||
		(method == http.MethodGet && r.Method != http.MethodGet) ||
		(method == http.MethodPut && r.Method != http.MethodPut) ||
		(method == http.MethodPatch && r.Method != http.MethodPatch) ||
		(method == http.MethodDelete && r.Method != http.MethodDelete) {

		return structs.CreateHttpError(
			fmt.Sprintf(
				"Method %v not allowed for this endpoint. Maybe you are looking for the allowed method: %v.",
				r.Method,
				method,
			),
			http.StatusNotImplemented,
			http.StatusText(http.StatusNotImplemented),
		)
	}

	return nil
}
