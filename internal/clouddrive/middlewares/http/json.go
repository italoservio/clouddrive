package middlewares

import (
	"net/http"
	"strings"

	"github.com/italoservio/clouddrive/internal/clouddrive/structs"
)

func Json() Middleware {
	return func(wri http.ResponseWriter, req *http.Request) *structs.HttpError {
		var accept_json = false
		var sends_json = false

		wri.Header().Set("content-type", "application/json")

		for key, values := range req.Header {
			if !accept_json && strings.ToLower(key) == "accept" {
				for _, val := range values {
					if strings.HasPrefix(val, "application/json") || val == "*/*" {
						accept_json = true
					}
				}
				continue
			}

			if !sends_json && strings.ToLower(key) == "content-type" {
				for _, val := range values {
					if strings.HasPrefix(val, "application/json") || val == "*/*" {
						sends_json = true
					}
				}
			}
		}

		if !accept_json {
			return structs.CreateHttpError(
				"This endpoint only works with JSON. Please set the correct 'Accept' header.",
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
			)
		}

		if !sends_json {
			return structs.CreateHttpError(
				"This endpoint only works with JSON. Please set the correct 'Content-Type' header.",
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
			)
		}

		return nil
	}
}
