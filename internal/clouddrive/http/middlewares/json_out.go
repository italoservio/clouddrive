package middlewares

import (
	"net/http"
	"strings"

	"github.com/italoservio/clouddrive/internal/clouddrive/errors"
)

func JsonOut() Middleware {
	return func(wri http.ResponseWriter, req *http.Request) *HttpResponse {

		wri.Header().Set("content-type", "application/json")

		accept := req.Header.Get("accept")
		if accept != "*/*" && !strings.Contains(accept, "application/json") {
			return CreateHttpResponse(
				errors.BAD_ACCEPT_HEADER_ERR,
				GetErrorHttpStatusByCode(errors.BAD_ACCEPT_HEADER_ERR),
			)
		}

		return nil
	}
}
