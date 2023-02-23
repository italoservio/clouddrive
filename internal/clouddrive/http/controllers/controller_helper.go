package controllers

import (
	"encoding/json"
	"errors"
	"github.com/italoservio/clouddrive/pkg/logger"

	custom_errors "github.com/italoservio/clouddrive/internal/clouddrive/errors"
	mid "github.com/italoservio/clouddrive/internal/clouddrive/http/middlewares"
	"io"
	"log"
	"net/http"
)

func ParseJsonToStruct[T any](
	wri http.ResponseWriter,
	input io.ReadCloser,
	output *T,
) error {
	if err := json.NewDecoder(input).Decode(&output); err != nil {
		logger.Error(err.Error())
		custom_err := errors.New(custom_errors.BAD_JSON_MARSHALING)
		HandleExecutionError(wri, custom_err)
		return custom_err
	}
	return nil
}

func HandleExecutionError(wri http.ResponseWriter, err error) {
	err_str := err.Error()
	response := mid.CreateHttpResponse(err_str, mid.GetErrorHttpStatusByCode(err_str))
	wri.WriteHeader(mid.GetErrorHttpStatusByCode(err_str))
	if _, err := wri.Write(response.ToJson()); err != nil {
		log.Fatal(err)
	}
}
