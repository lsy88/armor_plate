package utils

import (
	"armor_plate/controller/response"
	"armor_plate/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetClaims(c *gin.Context) (*jwt.CustomClaims, error) {
	//token := c.Request.Header.Get("Authorization")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		response.ResponseErrorWithMsg(c, response.CodeNoAuthLogin, "请求头缺少Auth Token")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.ResponseErrorWithMsg(c, response.CodeInvalidToken, "Token格式不对")
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	j := jwt.NewJWT()
	claims, err := j.ParseToken(parts[1])
	//fmt.Println(claims)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.CodeServerBusy, "解析jwt信息失败")
	}
	return claims, err
}

// GetUserID 从gin的context中获取jwt解析出来的用户id
func GetUserID(c *gin.Context) uint64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.EmployeeID
		}
	} else {
		waitUser := claims.(*jwt.CustomClaims)
		return waitUser.EmployeeID
	}
}

// GetUserAuthorityId 从gin的context中获取jwt解析出用户角色
func GetUserAuthorityId(c *gin.Context) uint64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.EmployeeID
		}
	} else {
		waitUse := claims.(*jwt.CustomClaims)
		return waitUse.EmployeeID
	}
}

//从gin的context中获取jwt解析出用户
func GetUserInfo(c *gin.Context) *jwt.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*jwt.CustomClaims)
		return waitUse
	}
}
