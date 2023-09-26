package api

var UNAUTHORIZED_REQUEST = ApiMessage{
	Error:   -102,
	Message: "unauthorized request",
	Data:    "",
}

var MISSING_PARAMS = ApiMessage{
	Error:   -103,
	Message: "missing params",
	Data:    "",
}
