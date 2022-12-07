package mysql

import (
	"armor_plate/model"
	"errors"
	"gorm.io/gorm"
)

func GetIDByDepotName(name string) (uint64, error) {
	var depot model.Depot
	err := db.Table("ap_depot").Where("depot_name = ?", name).First(&depot).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.New("不存在此仓库")
	}
	return depot.DepotID, nil
}

func CreateDepot(depot *model.Depot) (err error) {
	err = db.Table("ap_depot").Create(&depot).Error
	return
}

//根据id获取仓库的优先级
func GetPriByID(id uint64) (priority int) {
	var p model.Depot
	_ = db.Table("ap_depot").Where("depot_id = ?", id).First(&p)
	return p.Priority
}

//根据id获取仓库的持有量
func GetAmuByID(pid uint64, did uint64) (count int) {
	var p model.Address
	_ = db.Table("ap_address").Where("depot_id = ? and product_id = ?", did, pid).First(&p)
	return p.Count
}

//根据id获取仓库的地址
func GetPathByID(pid uint64) (path string) {
	var p model.Depot
	_ = db.Table("ap_depot").Where("depot_id = ?", pid).First(&p)
	return p.Path
}
