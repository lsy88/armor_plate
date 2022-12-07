package model

//货物
type Product struct {
	MODEL
	ProductID   uint64      `json:"product_id" gorm:"product_id;not null"`
	ProductName string      `json:"product_name" gorm:"product_name;not null"`
	CategoryID  uint        `json:"category_id" gorm:"category_id"`
	Category    []*Category `json:"-" gorm:"foreignKey:CategoryID;references:CategoryID"`
	Info        string      `json:"info" gorm:"size:1000"`                //货物描述信息
	Price       float64     `json:"price" gorm:"price"`                   //货物单价
	OnStorage   uint        `json:"onstorage" gorm:"onstorage;default:0"` //货物状态
	Addresses   []*Address  `json:"-" gorm:"foreignKey:ProductID;references:ProductID"`
}

func (Product) TableName() string {
	return "ap_product"
}
