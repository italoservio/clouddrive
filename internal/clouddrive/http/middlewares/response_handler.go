package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/errors"
	"github.com/italoservio/clouddrive/pkg/logger"
)

type HttpResponse struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Err        *HttpError  `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type HttpError struct {
	Code    string                  `json:"code,omitempty"`
	Message errors.ErrorMessageLang `json:"message,omitempty"`
}

func CreateHttpResponse(
	content interface{},
	status_code int,
) *HttpResponse {
	http_response := &HttpResponse{
		Status:     http.StatusText(status_code),
		StatusCode: status_code,
	}

	if fmt.Sprintf("%T", content) == "string" {
		error_code := fmt.Sprint(content)
		http_response.Err = &HttpError{
			Code:    error_code,
			Message: errors.GetErrorMessageByCode(error_code),
		}
	} else {
		http_response.Data = content
	}

	return http_response
}

func (hr *HttpResponse) ToJson() []byte {
	obj, err := json.Marshal(hr)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return obj
}

func GetErrorHttpStatusByCode(error_code string) int {
	switch error_code {
	case errors.BAD_ACCEPT_HEADER_ERR:
		return http.StatusBadRequest

	case errors.BAD_CONTENT_TYPE_HEADER_ERR:
		return http.StatusBadRequest

	case errors.BAD_HTTP_METHOD_VERB_ERR:
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}
