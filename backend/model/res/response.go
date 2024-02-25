package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg})
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func Fail(code int, data any, msg string, c *gin.Context) {
	c.JSON(Error, Response{
		Code: code,
		Data: data,
		Msg:  msg})
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, Ok := ErrorMap[code]
	if Ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误", c)
}
