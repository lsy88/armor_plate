package service

import (
	"armor_plate/controller/request"
	"armor_plate/dao/mysql"
	"armor_plate/model"
	"armor_plate/pkg/snowflake"
)

func CreateProduct(product *request.ReqCreateProduct) (err error) {
	id, _ := snowflake.GetID()
	var p = model.Product{
		ProductID:   id,
		ProductName: product.ProductName,
		CategoryID:  product.CategoryID,
		OnStorage:   product.OnStorage,
	}
	if err = mysql.CreateProduct(&p); err != nil {
		return err
	}
	return nil
}

func SetPathProduct(path *request.ReqProductPath) (err error) {
	productID, err := mysql.GetIDByProductName(path.ProductName)
	if err != nil {
		return
	}
	did, err := mysql.GetIDByDepotName(path.DepotName)
	if err != nil {
		return
	}
	var p = model.Address{
		ProductID: productID,
		DepotID:   did,
		Count:     path.Count,
	}
	if err = mysql.SetPathProduct(&p); err != nil {
		return err
	}
	return nil
}

func GetPathList(name string, page, size int) (p interface{}, c int, err error) {
	return mysql.GetPathListByName(name, page, size)
}
