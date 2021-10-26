package response

//Failed用于异常检测状态，设置HTTP状态码，比如404 not found, 403 forbidden
func Failed(code int) Response {
	return Response{
		Type: TypeFailed,
		FailedCode: code,
	}
}
