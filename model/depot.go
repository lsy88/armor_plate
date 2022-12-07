package model

//仓库信息
type Depot struct {
	MODEL
	DepotID   uint64  `json:"depot_id" gorm:"depot_id"`
	DepotName string  `json:"depot_name" gorm:"depot_name"`
	State     int     `json:"state" gorm:"state;default:1;comment:0-关闭,1-运营中"` //仓库状态
	Path      string  `json:"path" gorm:"path;type:varchar(50);not null"`      //地址
	Longitude float64 `json:"longitude" gorm:"longitude"`                      //经度
	Latitude  float64 `json:"latitude" gorm:"latitude"`                        //纬度
	Priority  int     `json:"priority" gorm:"priority;default:0"`              //优先级
}

func (Depot) TableName() string {
	return "ap_depot"
}
