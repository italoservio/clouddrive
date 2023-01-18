package middlewares

import (
	"net/http"
	"strings"

	"github.com/italoservio/clouddrive/internal/clouddrive/errors"
)

func JsonIn() Middleware {
	return func(wri http.ResponseWriter, req *http.Request) *HttpResponse {

		content_type := req.Header.Get("content-type")
		if content_type != "*/*" && !strings.Contains(content_type, "application/json") {
			return CreateHttpResponse(
				http.StatusBadRequest,
				errors.BAD_CONTENT_TYPE_HEADER_ERR,
				nil,
			)
		}

		return nil
	}
}
