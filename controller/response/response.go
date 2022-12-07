package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    StatusCode  `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, code StatusCode) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Decode(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code StatusCode, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Decode(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Decode(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
