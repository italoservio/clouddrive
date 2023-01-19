package middlewares

import (
	"fmt"
	"net/http"

	"github.com/italoservio/clouddrive/pkg/logger"
)

type Middleware func(http.ResponseWriter, *http.Request) *HttpResponse

func ReqHandler(
	functions []Middleware,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(wri http.ResponseWriter, req *http.Request) {
		logRequest(req)

		for _, fn := range functions {
			if err := fn(wri, req); err != nil {
				wri.WriteHeader(err.StatusCode)
				wri.Write(err.ToJson())
				logErr(err)
				return
			}
		}

		next(wri, req)
	}
}

func logRequest(req *http.Request) {
	logger.Info(fmt.Sprintf(
		"%s %s %s",
		req.Method,
		req.RequestURI,
		req.RemoteAddr,
	))
}

func logErr(err *HttpResponse) {
	logger.Error(fmt.Sprintf(
		"STATUS %d MESSAGE %v",
		err.StatusCode,
		err.Err,
	))
}
