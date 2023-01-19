package controllers

import (
	"encoding/json"
	"net/http"

	mid "github.com/italoservio/clouddrive/internal/clouddrive/http/middlewares"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases/dtos"
)

func Authenticate(wri http.ResponseWriter, req *http.Request) {
	var req_payload dtos.DTOAuthenticateReq

	if err := json.NewDecoder(req.Body).Decode(&req_payload); err != nil {
		err_str := err.Error()
		response := mid.CreateHttpResponse(err_str, mid.GetErrorHttpStatusByCode(err_str))
		wri.WriteHeader(mid.GetErrorHttpStatusByCode(err_str))
		wri.Write(response.ToJson())
		return
	}

	res_payload, err := usecases.Authenticate(req_payload)
	if err != nil {
		err_str := err.Error()
		response := mid.CreateHttpResponse(err_str, mid.GetErrorHttpStatusByCode(err_str))
		wri.WriteHeader(mid.GetErrorHttpStatusByCode(err_str))
		wri.Write(response.ToJson())
		return
	}

	wri.WriteHeader(http.StatusOK)
	res := mid.CreateHttpResponse(res_payload, http.StatusOK)
	wri.Write(res.ToJson())
}
