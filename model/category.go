package model

//钢材分类
type Category struct {
	MODEL
	CategoryID   uint   `json:"category_id" gorm:"category_id"`
	CategoryName string `json:"category_name" gorm:"category_name"`
}

func (Category) TableName() string {
	return "ap_category"
}
