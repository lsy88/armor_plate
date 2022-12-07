package middleware

import (
	"armor_plate/controller/response"
	"armor_plate/core"
	"armor_plate/service/casbin"
	"armor_plate/service/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.EmployeeID))
		success, _ := casbin.Enforcer.Enforce(sub, obj, act)
		if core.Conf.Mode == "dev" || success {
			c.Next()
		} else {
			response.ResponseError(c, response.CodeNoAuthority)
			c.Abort()
			return
		}
	}
}
