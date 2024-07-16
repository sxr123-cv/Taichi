package middleware

import (
	"Taichi/response"
	"Taichi/session"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		key := context.GetHeader("Cookie")
		var preload session.Preload
		err := session.VerifySession(key, &preload)
		if err != nil {
			response.Fail("", "参数不正确", response.NOT_PASS_AUTH, context)
			context.Abort()
			return
		}
		context.Set("user", preload)
		context.Next()
	}
}
