package response


// JSONResponse 定义JSONResponse的结构，最常用的结构体
type JSONResponse struct {
	Status string `json:"status"`
	Code Code   `json:"-"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}


// JSONData 用于正常情况 参数为body 如果没有body则可以传入nil
func JSONData(data interface{}) Response {
	return Response{
		Type: TypeJSON,
		Json: JSONResponse{
			Status: "success",
			Msg: "",
			Data: data,
		},
	}
}

// JSONError 用于错误情况，需要返回status为fail，同时有默认的code和msg
func JSONError(code Code) Response  {
	return Response{
		Type: TypeJSON,
		Json:JSONResponse{
			Status: "error",
			Code: code,
			Msg: GetMsg(code),
			Data: nil,
		},
	}
}

// JSONErrorWithMsg 用于错误情况，当code中没有对应的表示时，用于自定义msg返回
func JSONErrorWithMsg(msg string)Response  {
	return Response{
		Type: TypeJSON,
		Json: JSONResponse{
			Status: "error",
			Code: ERROR_DEFAULT,
			Msg: msg,
			Data: nil,
		},
	}
}

