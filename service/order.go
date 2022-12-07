package service

import (
	"armor_plate/controller/request"
	"armor_plate/dao/mysql"
	"armor_plate/model"
	"armor_plate/pkg/snowflake"
	"go.uber.org/zap"
)

func CreateOrderMes(mes *request.ReqOrderMes, eid uint64) (err error) {
	oid, _ := snowflake.GetID() //生成订单id
	pid, err := mysql.GetIDByProductName(mes.ProductName)
	if err != nil {
		return
	}
	var order = model.Order{
		OrderID:     oid, //生成订单id
		EmployeeID:  eid, //记录用户id
		ProductID:   pid, //记录货物id
		Amount:      mes.Amount,
		State:       mes.State,
		Destination: mes.Destination, //目的地址
	}
	if err = mysql.CreateOrder(&order); err != nil {
		zap.L().Error("mysql.create Order failed", zap.Error(err))
		return err
	}
	return nil
}

func DeliveryOrder(delivery *request.ReqDeliveryOrder) (err error) {
	order, err := mysql.GetOrderByID(delivery.OrderID)
	if order.Money == delivery.Money {
		//钱交足了，将状态标位已交付
		order.State = 1
	}
	//进行交付业务
	err = mysql.DeliveryOrder(order)
	return
}
