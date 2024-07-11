package main

import (
	"Taichi/response"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var req = LoginRequest{}
	r.POST("/login", func(c *gin.Context) {
		err := c.ShouldBindJSON(&req)
		if err != nil {
			response.Fail("", "参数不正确", response.NOT_PASS_AUTH, c)
			return
		}
		if req.Pwd == "11111" && req.Name == "11111" {
			response.Success(true, "", c)
			return
		}
		response.Fail("", "参数不正确", response.NOT_PASS_AUTH, c)
		return
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

type LoginRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}
