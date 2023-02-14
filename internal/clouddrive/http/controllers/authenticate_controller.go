package controllers

import (
	"log"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	mid "github.com/italoservio/clouddrive/internal/clouddrive/http/middlewares"
	"github.com/italoservio/clouddrive/internal/clouddrive/services"
)

func AuthenticateUser(wri http.ResponseWriter, req *http.Request) {
	var req_payload dtos.DTOAuthenticateReq

	err := ParseJsonToStruct(wri, req.Body, &req_payload)
	if err != nil {
		return
	}

	res_payload, err := services.Authenticate(req_payload)
	if err != nil {
		HandleExecutionError(wri, err)
		return
	}

	wri.WriteHeader(http.StatusOK)
	res := mid.CreateHttpResponse(res_payload, http.StatusOK)
	if _, err := wri.Write(res.ToJson()); err != nil {
		log.Fatal(err)
	}
}
