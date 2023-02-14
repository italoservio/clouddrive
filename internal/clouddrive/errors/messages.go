package errors

type ErrorMessageLang struct {
	PtBr string `json:"pt_br,omitempty"`
	EnUs string `json:"en_us,omitempty"`
}

func GetErrorMessageByCode(error_code string) ErrorMessageLang {
	switch error_code {
	case BAD_ACCEPT_HEADER_ERR:
		return ErrorMessageLang{
			PtBr: "Cabeçalho 'Accept' inválido, por favor verifique o tipo de saída deste endpoint e acrescente o valor correto",
			EnUs: "Invalid 'Accept' header, please verify the response type for this endpoint and set the correct value",
		}

	case BAD_CONTENT_TYPE_HEADER_ERR:
		return ErrorMessageLang{
			PtBr: "Cabeçalho 'Content-Type' inválido, por favor verifique o tipo de entrada deste endpoint e acrescente o valor correto",
			EnUs: "Invalid 'Content-Type' header, please verify the request type for this endpoint and set the correct value",
		}

	case BAD_HTTP_METHOD_VERB_ERR:
		return ErrorMessageLang{
			PtBr: "Verbo HTTP incorreto, por favor verifique o verbo correto para este endpoint",
			EnUs: "Invalid HTTP verb, please verify the correct verb for this endpoint",
		}

	case BAD_JSON_MARSHALING:
		return ErrorMessageLang{
			PtBr: "Falha na conversão do JSON de entrada",
			EnUs: "Failed on input JSON marshaling",
		}

	case BAD_DB_CALL:
		return ErrorMessageLang{
			PtBr: "Falha na comunicação com o banco de dados",
			EnUs: "Failed to establish communication with database",
		}

	case BAD_CONFLICT:
		return ErrorMessageLang{
			PtBr: "Há existe um registro semelhante e não pode haver duplicidade",
			EnUs: "There is a similar record and there can be no duplicate",
		}

	case NOT_FOUND_REGISTRY:
		return ErrorMessageLang{
			PtBr: "Nenhum registro encontrado para os parâmetros de busca informados",
			EnUs: "No records found for the informed search parameters",
		}

	case BAD_PATH_PARAMS:
		return ErrorMessageLang{
			PtBr: "Parâmetros de rota em formato inválido para esta rota",
			EnUs: "Malformed path params for this route",
		}

	case BAD_BODY_PARAMS:
		return ErrorMessageLang{
			PtBr: "Parâmetro do corpo da requisição são inválidos para esta rota",
			EnUs: "Malformed body params for this route",
		}

	case BAD_QUERY_PARAMS:
		return ErrorMessageLang{
			PtBr: "Parâmetros de busca da requisição são inválidos para esta rota",
			EnUs: "Malformed query params for this route",
		}

	default:
		return ErrorMessageLang{
			PtBr: "Erro interno e inesperado",
			EnUs: "Internal and unexpected error",
		}
	}
}
