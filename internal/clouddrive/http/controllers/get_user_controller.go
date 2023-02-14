package controllers

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	custom_errors "github.com/italoservio/clouddrive/internal/clouddrive/errors"
	mid "github.com/italoservio/clouddrive/internal/clouddrive/http/middlewares"
	"github.com/italoservio/clouddrive/internal/clouddrive/services"
)

func GetUserById(wri http.ResponseWriter, req *http.Request) {
	path_arr := strings.Split(req.URL.Path, "/")
	raw_id := path_arr[len(path_arr)-1]
	id := (regexp.MustCompile(`[ \/]+`)).ReplaceAllString(raw_id, "")
	if id == "" {
		HandleExecutionError(wri, errors.New(custom_errors.BAD_PATH_PARAMS))
		return
	}

	req_payload := dtos.DTOGetUserByIdReq{Id: id}
	res_payload, err := services.UserById(req_payload)
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
