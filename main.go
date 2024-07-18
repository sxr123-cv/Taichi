package main

import (
	"Taichi/jwt"
	"Taichi/middleware"
	"Taichi/response"
	"Taichi/session"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	needAuthRouter := r.Group("client").Use(middleware.AuthJwt())
	noAuthRouter := r.Group("client")
	gin.Logger()
	var req = LoginRequest{}
	noAuthRouter.POST("/login", func(c *gin.Context) {
		err := c.ShouldBindJSON(&req)
		if err != nil {
			response.Fail("", "参数不正确", response.NOT_PASS_AUTH, c)
			return
		}
		if req.Pwd == "11111" && req.Name == "11111" {
			saveSession, err := jwt.GetJwt(session.Preload{
				Role:   "student",
				UserId: 1,
			}, 30*60)
			if err != nil {
				response.Fail("", "签发失败", response.NOT_PASS_AUTH, c)
				return
			} else {
				response.Success(saveSession, "", c)
				return
			}

		}
		response.Fail("", "参数不正确", response.NOT_PASS_AUTH, c)
		return
	})
	needAuthRouter.POST("/login2", func(c *gin.Context) {
		v, _ := c.Get("user")
		response.Success(v, "", c)
		return
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

type LoginRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}
