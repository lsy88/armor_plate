package middleware

import (
	"armor_plate/controller/response"
	"armor_plate/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ResponseErrorWithMsg(c, response.CodeNoAuthLogin, "请求头缺少Auth Token")
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.ResponseErrorWithMsg(c, response.CodeInvalidToken, "Token格式不对")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		j := jwt.NewJWT()
		mc, err := j.ParseToken(parts[1])
		//fmt.Println(mc)
		if err != nil {
			fmt.Println(err)
			response.ResponseError(c, response.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		//c.Set("claims", mc)
		c.Set("claims", mc)
		c.Next() // 后续的处理函数可以用过c.Get("userID")来获取当前请求的用户信息
	}
}
