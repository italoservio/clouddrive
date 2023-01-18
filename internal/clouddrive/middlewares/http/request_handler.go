package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/logger"
)

func ReqHandler(
	functions []Middleware,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(wri http.ResponseWriter, req *http.Request) {
		logRequest(req)

		for _, fn := range functions {
			if err0 := fn(wri, req); err0 != nil {
				json, err1 := json.Marshal(err0)
				if err1 != nil {
					log.Fatalf(err1.Error())
				}
				wri.WriteHeader(err0.StatusCode)
				wri.Write(json)
				return
			}
		}

		next(wri, req)
	}
}

func logRequest(req *http.Request) {
	logger.Info(fmt.Sprintf(
		"[%s - %s] BODY %v HEADERS %v",
		req.Method,
		req.RemoteAddr,
		req.Body,
		req.Header,
	))
}
