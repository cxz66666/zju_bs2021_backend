package response

type Code int
const (
	ERROR_AUTH_NO_VALID_HEADER Code=10008 //请求头格式错误（即放token的格式不对）
	ERROR_NOT_LOGIN Code=10009  //请求头没有token，未登录
	ERROR_TOKEN_NOT_VAILD Code=10010 //token不合法
	ERROR_TOKEN_EXPIRED Code=10011 //token已过期

	ERROR_NOT_ADMIN Code=10012//不是管理员

	ERROR_PARAM_FAIL          Code =10001  //登陆模型绑定参数错误
	ERROR_TOKEN_GENERATE_FAIL Code = 10002 //token生成错误
	ERROR_PASSWORD              Code   =10003  //密码错误
	ERROR_USERID                 Code  =10004  //账号错误
	ERROR_USER_NOT_FOUND  Code =10005 //账号不存在（已被删除）

	ERROR_UPDATE_FAIL Code = 10100 //更新模型错误

	ERROR_DATABASE_QUERY Code = 20000// 数据库内部错误

	//to be continue ....

	ERROR_DEFAULT Code= 30000 //未知错误
)