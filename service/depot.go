package service

import (
	"armor_plate/controller/request"
	"armor_plate/dao/mysql"
	"armor_plate/model"
	"armor_plate/pkg/baiduMap"
	"armor_plate/pkg/snowflake"
	"go.uber.org/zap"
)

func CreateDepot(req *request.ReqDepotMes) (err error) {
	did, _ := snowflake.GetID()
	latitude, longitude := baiduMap.Get_JWData_By_Ctiy(req.Path[:6])
	var depot = model.Depot{
		DepotID:   did,
		DepotName: req.DepotName,
		Path:      req.Path,
		Priority:  req.Priority,
		Latitude:  latitude,
		Longitude: longitude,
	}
	if err = mysql.CreateDepot(&depot); err != nil {
		zap.L().Error("mysql.CreateDepot failed", zap.Error(err))
		return
	}
	return nil
}
