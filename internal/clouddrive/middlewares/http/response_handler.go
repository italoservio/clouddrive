package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/errors"
	"github.com/italoservio/clouddrive/pkg/logger"
)

type HttpResponse struct {
	StatusCode int               `json:"status_code"`
	Status     string            `json:"status"`
	Err        *errors.HttpError `json:"error,omitempty"`
	Data       interface{}       `json:"data,omitempty"`
}

func CreateHttpResponse(
	status_code int,
	error_code string,
	data interface{},
) *HttpResponse {
	http_response := &HttpResponse{
		Status:     http.StatusText(status_code),
		StatusCode: status_code,
	}

	if error_code != "" {
		http_response.Err = &errors.HttpError{
			Code:    error_code,
			Message: errors.GetErrorMessageByCode(error_code),
		}
	} else if data != nil {
		http_response.Data = data
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
