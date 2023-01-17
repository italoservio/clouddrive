package structs

type HttpError struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func CreateHttpError(message string, status_code int, status string) *HttpError {
	return &HttpError{
		Status:     status,
		StatusCode: status_code,
		Message:    message,
	}
}

func (he *HttpError) Error() string {
	return he.Message
}
