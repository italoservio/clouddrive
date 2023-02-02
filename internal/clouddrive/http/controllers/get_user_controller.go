package controllers

import (
	"log"
	"net/http"

	mid "github.com/italoservio/clouddrive/internal/clouddrive/http/middlewares"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases/dtos"
)

func GetUser(wri http.ResponseWriter, req *http.Request) {
	var req_payload dtos.DTOGetUserReq
	err := ParseJsonToStruct[dtos.DTOGetUserReq](wri, req.Body, &req_payload)
	if err != nil {
		return
	}

	res_payload, err := usecases.GetUser(req_payload)
	if err != nil {
		HandleExecutionError(wri, err)
		return
	}

	wri.WriteHeader(http.StatusCreated)
	res := mid.CreateHttpResponse(res_payload, http.StatusOK)
	if _, err := wri.Write(res.ToJson()); err != nil {
		log.Fatal(err)
	}
}