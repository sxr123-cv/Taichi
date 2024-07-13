package main

import (
	"Taichi/response"
	"Taichi/session"
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
			saveSession, err := session.SaveSession(session.Preload{
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

	r.POST("/login2", func(c *gin.Context) {
		//err := c.ShouldBindJSON(&req)
		//if err != nil {
		//	response.Fail("", "参数不正确", response.NOT_PASS_AUTH, c)
		//	return
		//}
		key := c.GetHeader("Cookie")
		var preload session.Preload
		err := session.VerifySession(key, &preload)
		if err != nil {
			response.Fail("", "参数不正确", response.NOT_PASS_AUTH, c)
			return
		}
		response.Success(preload, "", c)
		//if req.Pwd == "11111" && req.Name == "11111" {
		//	saveSession, err := session.SaveSession(session.Preload{
		//		Role:   "student",
		//		UserId: 1,
		//	}, 30*60)
		//	if err != nil {
		//		response.Fail("", "签发失败", response.NOT_PASS_AUTH, c)
		//		return
		//	} else {
		//		response.Success(saveSession, "", c)
		//		return
		//	}
		//
		//}
		//response.Fail("", "参数不正确", response.NOT_PASS_AUTH, c)
		return
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

type LoginRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}
