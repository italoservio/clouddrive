package middlewares

import "net/http"

type Middleware func(http.ResponseWriter, *http.Request) *HttpResponse
