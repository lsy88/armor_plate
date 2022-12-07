package model

//存储位置
type Address struct {
	MODEL
	ProductID uint64 `json:"product_id" gorm:"product_id;not null"`
	DepotID   uint64 `json:"depot_id" gorm:"depot_id"` //仓库编号
	Count     int    `json:"count" gorm:"count"`       //仓库持有量
	Depot     Depot  `json:"-" gorm:"-;foreignKey:DepotID;references:DepotID"`
}

func (Address) TableName() string {
	return "ap_address"
}
