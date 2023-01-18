package errors

func GetErrorMessageByCode(error_code string) HttpErrorMessageLang {
	switch error_code {
	case BAD_ACCEPT_HEADER_ERR:
		return HttpErrorMessageLang{
			PtBr: "Cabeçalho 'Accept' inválido, por favor verifique o tipo de saída deste endpoint e acrescente o valor correto",
			EnUs: "Invalid 'Accept' header, please verify the response type for this endpoint and set the correct value",
		}

	case BAD_CONTENT_TYPE_HEADER_ERR:
		return HttpErrorMessageLang{
			PtBr: "Cabeçalho 'Content-Type' inválido, por favor verifique o tipo de entrada deste endpoint e acrescente o valor correto",
			EnUs: "Invalid 'Content-Type' header, please verify the request type for this endpoint and set the correct value",
		}

	case BAD_HTTP_METHOD_VERB_ERR:
		return HttpErrorMessageLang{
			PtBr: "Verbo HTTP incorreto, por favor verifique o verbo correto para este endpoint",
			EnUs: "Invalid HTTP verb, please verify the correct verb for this endpoint",
		}

	default:
		return HttpErrorMessageLang{
			PtBr: "Erro interno e inesperado",
			EnUs: "Internal and unexpected error",
		}
	}
}
