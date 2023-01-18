package errors

type HttpError struct {
	Code    string               `json:"code,omitempty"`
	Message HttpErrorMessageLang `json:"message,omitempty"`
}

type HttpErrorMessageLang struct {
	PtBr string `json:"pt_br,omitempty"`
	EnUs string `json:"en_us,omitempty"`
}
