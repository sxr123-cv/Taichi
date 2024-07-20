package middleware

import (
	auth_const "Taichi/auth/const"
	jwt "Taichi/auth/jwt"
	"Taichi/response"
	"github.com/gin-gonic/gin"
)

func AuthJwt() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwtData := context.GetHeader("Authorization")
		var preload auth_const.Preload
		err := jwt.VerifyJWT(jwtData, &preload)
		if err != nil {
			response.Fail("", "参数不正确", response.NOT_PASS_AUTH, context)
			context.Abort()
			return
		}
		context.Set("user", preload)
		context.Next()
	}
}

func Logout() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwtData := context.GetHeader("Authorization")
		var preload auth_const.Preload
		err := jwt.Logout(jwtData, &preload)
		if err != nil {
			response.Fail("", "参数不正确", response.NOT_PASS_AUTH, context)
			context.Abort()
			return
		}
		context.Next()
	}
}
