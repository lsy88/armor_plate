package controller

import (
	"armor_plate/controller/request"
	. "armor_plate/controller/response"
	"armor_plate/service"
	"armor_plate/service/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateOrderHandler 创建订单业务
func CreateOrderHandler(c *gin.Context) {
	var orderMes request.ReqOrderMes
	if err := c.ShouldBindJSON(&orderMes); err != nil {
		zap.L().Error("shouldBind orderMes failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	//从c解析出当前用户id
	employeeId := utils.GetUserID(c)
	if err := service.CreateOrderMes(&orderMes, employeeId); err != nil {
		zap.L().Error("service.CreateOrderMes failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

//交付订单业务
func DeliveryOrderHandler(c *gin.Context) {
	var delivery request.ReqDeliveryOrder
	if err := c.ShouldBindJSON(&delivery); err != nil {
		zap.L().Error("shouldBind ReqDeliveryOrder failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if err := service.DeliveryOrder(&delivery); err != nil {
		zap.L().Error("service.DeliveryOrder failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
