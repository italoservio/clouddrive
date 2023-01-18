package middlewares

import (
	"net/http"

	"github.com/italoservio/clouddrive/internal/clouddrive/structs"
)

type Middleware func(http.ResponseWriter, *http.Request) *structs.HttpError
