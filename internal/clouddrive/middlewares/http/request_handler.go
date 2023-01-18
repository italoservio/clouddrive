package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ReqHandler(
	functions []Middleware,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(wri http.ResponseWriter, req *http.Request) {
		time := time.Now()
		date_time := fmt.Sprintf(
			"%s/%s/%d %d:%d:%d",
			fmt.Sprintf("%02d", time.Day()),
			fmt.Sprintf("%02d", time.Month()),
			time.Year(),
			time.Hour(),
			time.Minute(),
			time.Second(),
		)

		fmt.Printf(
			"[%s - %s - %s] BODY %v HEADERS %v\n",
			req.Method,
			date_time,
			req.RemoteAddr,
			req.Body,
			req.Header,
		)

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
