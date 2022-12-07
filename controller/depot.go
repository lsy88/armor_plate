package controller

import (
	"armor_plate/controller/request"
	. "armor_plate/controller/response"
	"armor_plate/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateDepotHandler(c *gin.Context) {
	var depotMes request.ReqDepotMes
	if err := c.ShouldBindJSON(&depotMes); err != nil {
		zap.L().Error("shouldBind create auth failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if err := service.CreateDepot(&depotMes); err != nil {
		zap.L().Error("shouldBind CreateDepot failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
