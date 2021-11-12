package response

import (
	"annotation/utils/jsonu"
	"annotation/utils/logging"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type Response struct {
	Type         Type         // 返回类型
	Json         JSONResponse // JSON数据
	File         []byte       // 文件数据
	FileName     string       // 文件名
	FailedCode   int          // 出错时的http状态码
	RedirectURL  string       // 重定向url
	RedirectCode int          // 重定向code
}

func (r *Response) Write(c *gin.Context) {
	switch r.Type {
	case TypeJSON:
		marshal:= jsonu.Marshal(r.Json)
		c.JSON(http.StatusOK, json.RawMessage(marshal))
	case TypeFile:
		escape := url.QueryEscape(r.FileName)
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", escape))
		c.Data(http.StatusOK, "application/octet-stream", r.File)
	case TypeRedirect:
		c.Redirect(r.RedirectCode, r.RedirectURL)
	case TypeFailed:
		c.Status(r.FailedCode)

	case TypeImage:
		c.File(r.FileName)
	}
}


// Logging need to use after logging.Setup(),会向日志输出当前
func (r *Response)Logging() error {
	logging.Info(r.String())
	return nil
}



// String 将response的概要信息转为字符串 用于打日志
func (r *Response) String() string {
	switch r.Type {
	case TypeJSON: // JSON
		return fmt.Sprintf("%+v", struct {
			Type Type   // 类型
			Json string // JSON数据
		}{r.Type, r.ToJSON()})

	case TypeFile:
		return fmt.Sprintf("%+v", struct {
			Type     Type   // 类型
			File     string // 文件数据（只显示大小）
			FileName string // 文件名
		}{r.Type, fmt.Sprintf("<%d byte>", len(r.File)), r.FileName})
	case TypeFailed:
		return fmt.Sprintf("%+v", struct {
			Type Type //类型
			FailedCode int //错误码
		}{r.Type,r.FailedCode})
	case TypeRedirect:
		return fmt.Sprintf("%+v", struct {
			Type Type
			RedirectURL  string       // 重定向url
			RedirectCode int          // 重定向code
		}{r.Type,r.RedirectURL,r.RedirectCode})
	default:
		return "<unknown>"
	}
}



// ToJSON 用于json类型的调试， 将response的json data 转换成JSON字符串，方便打印，注意 只会打印json部分
func (r *Response) ToJSON() string {
	marshal:= jsonu.Marshal(r.Json)
	return marshal
}