package middlewares

import (
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/errors"
)

func Method(method string) Middleware {
	return func(wri http.ResponseWriter, req *http.Request) *HttpResponse {
		if req.Method == http.MethodOptions {
			wri.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,PATCH,DELETE")
			wri.Header().Set("Access-Control-Allow-Headers", "*")
			wri.Header().Set("Access-Control-Allow-Origin", "*")
			wri.Header().Set("Access-Control-Max-Age", "86400")
			return CreateHttpResponse(
				http.StatusNoContent,
				"",
				nil,
			)
		}

		if (method == http.MethodPost && req.Method != http.MethodPost) ||
			(method == http.MethodGet && req.Method != http.MethodGet) ||
			(method == http.MethodPut && req.Method != http.MethodPut) ||
			(method == http.MethodPatch && req.Method != http.MethodPatch) ||
			(method == http.MethodDelete && req.Method != http.MethodDelete) {

			return CreateHttpResponse(
				http.StatusNotImplemented,
				errors.BAD_HTTP_METHOD_VERB_ERR,
				nil,
			)
		}

		return nil
	}
}
