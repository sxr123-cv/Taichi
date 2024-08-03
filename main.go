package main

import (
	authconst "Taichi/auth/const"
	jwt "Taichi/auth/jwt"
	"Taichi/config"
	"Taichi/db"
	"Taichi/db/model"
	"Taichi/middleware"
	"Taichi/response"
	"Taichi/sdk"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	readConfig, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	err = db.InitMySQL(readConfig.MySQL)
	sdk.InitRedis(readConfig.Redis)
	if err != nil {
		return
	}
	needAuthRouter := r.Group("client").Use(middleware.AuthJwt())
	noAuthRouter := r.Group("client")
	var req = LoginRequest{}
	noAuthRouter.POST("/login", func(c *gin.Context) {
		err := c.ShouldBindJSON(&req)
		if err != nil {
			response.Fail("", "参数不正确", response.NOT_PASS_AUTH, c)
			return
		}
		err = model.UserModel.InsertUser(&model.User{
			Id:       0,
			Name:     "111",
			Password: "2222",
			Sex:      "女",
		})
		if err != nil {
			return
		}
		if req.Pwd == "11111" && req.Name == "11111" {
			saveSession, err := jwt.GetJwt(authconst.Preload{
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
	needAuthRouter.Use(middleware.Logout()).POST("/logout", func(c *gin.Context) {
		response.Success("OK", "", c)
		return
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

type LoginRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}
