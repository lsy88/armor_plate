package controller

import (
	"armor_plate/controller/request"
	. "armor_plate/controller/response"
	"armor_plate/dao/mysql"
	"armor_plate/model"
	"armor_plate/pkg/jwt"
	"armor_plate/service"
	"armor_plate/service/casbin"
	"armor_plate/service/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

//登录业务
func LoginHandler(c *gin.Context) {
	var lg request.ReqLoginForm
	if err := c.ShouldBindJSON(&lg); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	//用户登录
	if err := service.LoginUser(&lg); err != nil {
		zap.L().Error("mysql Login failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	//生成jwt
	j := jwt.NewJWT()
	aToken, rToken, _ := j.GenToken(lg.EmployeeName)
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       lg.EmployeeID,
		"userName":     lg.EmployeeName,
	})
}

// RegisterHandler
// @Tags 用户业务接口
// @Summary 用户注册
// @Accept application/json
// @Produce application/json
// @host localhost:9090
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query request.RegisterForm true "查询参数"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /server/v1/user/register [post]
func RegisterHandler(c *gin.Context) {
	var fo request.ReqRegisterForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("shouldBind register failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if fo.Password != fo.ConfirmPassword {
		ResponseErrorWithMsg(c, CodeInvalidPassword, "密码不一致")
		return
	}
	//注册用户
	err := service.RegisterUser(&fo)
	if errors.Is(err, mysql.ErrorUserExit) {
		ResponseError(c, CodeUserExist)
		return
	}
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "恭喜"+fo.EmployeeName+",注册成功")
}

// SetUserAuthHandler 为用户设置权限
func SetUserAuthHandler(c *gin.Context) {
	var sau request.SetUserAuth
	if err := c.ShouldBindJSON(&sau); err != nil {
		zap.L().Error("shouldBind set auth failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if err := service.SetUserAuth(sau.EID, sau.AuthorityId); err != nil {
		zap.L().Error("set auth failed", zap.Error(err))
		ResponseError(c, CodeAuthCreateFail)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// CreateAuthHandler 创建角色
func CreateAuthHandler(c *gin.Context) {
	var auth request.ReqSetRoleInfo
	if err := c.ShouldBindJSON(&auth); err != nil {
		zap.L().Error("shouldBind create auth failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	auth.Path = c.Request.URL.Path
	auth.Method = c.Request.Method
	if authBack, err := service.CreateUserAuth(&auth); err != nil {
		zap.L().Error("create auth failed", zap.Error(err))
		ResponseError(c, CodeAuthCreateFail)
	} else {
		info := model.DefaultCasbin()
		_ = casbin.CasbinServiceApp.UpdateCasbin(auth.AuthorityId, info)
		ResponseSuccess(c, gin.H{
			"Authority": authBack,
		})
	}
	
}

// RefreshTokenHandler 刷新token
func RefreshTokenHandler(c *gin.Context) {
	//rt := c.Query("refresh_token")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	j := jwt.NewJWT()
	aToken, rToken, err := j.RefreshToken(parts[1])
	if err != nil {
		zap.L().Error("refresh token failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

// UpdateHandler 修改用户信息
func UpdateHandler(c *gin.Context) {
	var fo request.ReqChangeEmployeeReq
	if err := c.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("shouldBind Req failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	id := utils.GetUserID(c)
	if err := service.UpdateUser(&fo, id); err != nil {
		zap.L().Error("service update failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "恭喜"+strconv.FormatUint(id, 10)+"修改信息成功")
}

// GetEmpListHandler 分页获取用户列表
func GetEmpListHandler(c *gin.Context) {
	page, size := request.GetPageInfo(c)
	//if err := c.ShouldBindJSON(&pageInfo); err != nil {
	//	zap.L().Error("shouldBind pageInfo failed", zap.Error(err))
	//	ResponseError(c, CodeInvalidParams)
	//	return
	//}
	var pageInfo request.PageInfo
	pageInfo.Page = page
	pageInfo.Size = size
	list, total, err := service.GetUserList(&pageInfo)
	if err != nil {
		zap.L().Error("service getUserList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"Total":        total,
		"EmployeeList": list,
	})
}
