package controller

import (
	"armor_plate/controller/request"
	. "armor_plate/controller/response"
	"armor_plate/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//新建货物
func AddProductHandler(c *gin.Context) {
	var cp request.ReqCreateProduct
	if err := c.ShouldBindJSON(&cp); err != nil {
		zap.L().Error("should bind product failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	if err := service.CreateProduct(&cp); err != nil {
		zap.L().Error("service product failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// SetProductPathHandler 货物入仓
func SetProductPathHandler(c *gin.Context) {
	var pPath request.ReqProductPath
	if err := c.ShouldBindJSON(&pPath); err != nil {
		zap.L().Error("should bind product path failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	if err := service.SetPathProduct(&pPath); err != nil {
		zap.L().Error("service SetPathProduct failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// GetPathListHandler 获取货物存储位置列表
func GetPathListHandler(c *gin.Context) {
	pName := c.Query("product_name")
	page, size := request.GetPageInfo(c)
	list, count, err := service.GetPathList(pName, page, size)
	if err != nil {
		zap.L().Error("service GetPathList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"ProductName": pName,
		"Amount":      count,
		"PathList":    list,
	})
	
}
