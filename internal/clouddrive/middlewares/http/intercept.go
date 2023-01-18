package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

func Intercept(
	functions []Middleware,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(wri http.ResponseWriter, req *http.Request) {

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
