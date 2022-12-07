package life

import "github.com/flipped-aurora/gin-vue-admin/server/global"

//商品信息
type Good struct {
	global.GVA_MODEL
	GoodName    string  `json:"goodName" gorm:"good_name"`
	MainPicture string  `json:"mainPicture" gorm:"main_picture;comment:'商品主图'"`
	DetailImage string  `json:"detailImage" gorm:"detail_image;comment:'商品具体图'"`
	Description string  `json:"description" gorm:"size:100;comment:'商品信息简介'"`
	CategoryId  uint    `json:"categoryId" gorm:"category_id;comment:'商品分类id'"`
	MerchantId  uint    `json:"merchantId" gorm:"comment:'商家id'"`
	OnSale      int     `json:"onSale" gorm:"on_sale;default:1;comment:'商品状态:1-在售,2-下架'"`
	Price       float64 `json:"price" gorm:"price"`
}
