package errors

const (
	BAD_ACCEPT_HEADER_ERR       = "BAD_ACCEPT_HEADER_ERR"
	BAD_CONTENT_TYPE_HEADER_ERR = "BAD_CONTENT_TYPE_HEADER_ERR"
	BAD_HTTP_METHOD_VERB_ERR    = "BAD_HTTP_METHOD_VERB_ERR"
	UNEXPECTED_ERR              = "UNEXPECTED_ERR"
	BAD_JSON_MARSHALING         = "BAD_JSON_MARSHALING"

	BAD_DB_CALL        = "BAD_DB_CALL"
	BAD_CONFLICT       = "BAD_CONFLICT"
	NOT_FOUND_REGISTRY = "NOT_FOUND_REGISTRY"
	BAD_PATH_PARAMS    = "BAD_PATH_PARAMS"
	BAD_BODY_PARAMS    = "BAD_BODY_PARAMS"
	BAD_QUERY_PARAMS   = "BAD_QUERY_PARAMS"
)
