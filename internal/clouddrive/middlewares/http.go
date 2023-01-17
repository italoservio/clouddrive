package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

func Http(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(wri http.ResponseWriter, req *http.Request) {
		if result := HttpHeaders(wri, req); result != nil {
			json, err := json.Marshal(result)
			if err != nil {
				log.Fatalf(err.Error())
			}
			wri.WriteHeader(result.StatusCode)
			wri.Write(json)
			return
		}

		if result := HttpMethod(method, wri, req); result != nil {
			json, err := json.Marshal(result)
			if err != nil {
				log.Fatalf(err.Error())
			}
			wri.WriteHeader(result.StatusCode)
			wri.Write(json)
			return
		}

		next(wri, req)
	}
}
