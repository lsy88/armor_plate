package mysql

import (
	"armor_plate/model"
	"errors"
	"gorm.io/gorm"
)

//新建货物
func CreateProduct(cp *model.Product) (err error) {
	u := db.Where("product_id = ?", cp.ProductID).Find(&model.Product{})
	if u.RowsAffected != 0 {
		return errors.New("该产品已存在")
	}
	err = db.Model(&model.Product{}).Create(&cp).Error
	return err
}

func GetIDByProductName(name string) (uint64, error) {
	var m model.Product
	err := db.Model(&model.Product{}).Where("product_name = ?", name).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.New("不存在该货物")
	}
	return m.ProductID, nil
	
}

//设置货物存储位置
func SetPathProduct(path *model.Address) (err error) {
	err = db.Model(&model.Address{}).Create(&path).Error
	return
}

// GetPathListByName 获取货物存储位置列表
func GetPathListByName(name string, page, size int) (list []*model.Address, count int, err error) {
	pid, err := GetIDByProductName(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, errors.New("不存在该货物,无法获取存储位置")
	}
	err = db.Table("ap_address").Preload("ap_address").Where("product_id = ?", pid).Offset((page - 1) * size).Limit(size).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	//该类货物的所有数量
	for _, v := range list {
		count += v.Count
	}
	
	return
}

// GetPathListByID 获取持有货物所有仓库列表
func GetPathListByID(pid uint64) (depotList []uint64, err error) {
	list := []*model.Address{}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("不存在该货物,无法获取存储位置")
	}
	err = db.Table("ap_address").Where("product_id = ?", pid).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		depotList = append(depotList, v.DepotID)
	}
	return
}

//获取货物总数量
func GetProductAmount(pid uint64) (amount int, err error) {
	var list []model.Address
	err = db.Table("ap_address").Where("product_id = ?", pid).Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.New("仓库中没有该类物品")
	}
	//获取总数量
	for _, v := range list {
		amount += v.Count
	}
	return
}

//根据物品ID获取物品价格
func GetPriceById(pid uint64) (float64, error) {
	var ap model.Product
	err := db.Table(("ap_product")).Where("product_id", pid).First(&ap).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.New("仓库中没有该类物品")
	}
	return ap.Price, nil
}
