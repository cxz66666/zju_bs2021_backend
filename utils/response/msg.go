package response


var MsgFlags =map[Code]string{
	ERROR_PARAM_FAIL:           "参数错误，请检查调用的参数",
	ERROR_TOKEN_GENERATE_FAIL:  "token生成错误",
	ERROR_DEFAULT:              "未知错误",
	ERROR_NOT_LOGIN:            "未登录，请先登录",
	ERROR_TOKEN_NOT_VAILD:      "令牌不合法，请重新登录",
	ERROR_TOKEN_EXPIRED:        "token令牌已过期,请重新获取",
	ERROR_AUTH_NO_VALID_HEADER: "请求头格式错误，请检查NOTIFY_AUTH_BEARER字段",
	ERROR_NOT_ADMIN:            "您不是管理员",
	ERROR_PASSWORD :            "密码错误",
	ERROR_USERID :              "账号错误不存在",
	ERROR_USER_NOT_FOUND :		"账号目前未在数据库中找到",
	ERROR_UPDATE_FAIL : 		"更新模型错误，请重试",
	ERROR_DATABASE_QUERY: 		"数据库内部查询错误，请重试",
	ERROR_IP_BLOCK: 			"登录失败次数过多，暂时封禁IP，请5min后重试",
	ERROR_ADMIN_INVALID_PASSWORD: "管理员密码错误",
	ERROR_NOT_VALID_USER_PARAM: "用户创建参数不合法",
}

func GetMsg(code Code) string {
	msg,ok:=MsgFlags[code]
	if ok{
		return msg
	}
	return MsgFlags[ERROR_DEFAULT]
}