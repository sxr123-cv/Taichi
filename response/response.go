package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any, msg string, c *gin.Context) {
	c.JSON(200, Response{Code: SUCCESS, Msg: msg, Data: data})
	c.AbortWithStatus(200)
	return
}

func Fail(data any, msg string, code int, c *gin.Context) {
	c.JSON(200, Response{Code: code, Msg: msg, Data: data})
	return
}
