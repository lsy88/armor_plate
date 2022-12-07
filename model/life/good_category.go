package life

import "github.com/flipped-aurora/gin-vue-admin/server/global"

//商品分类
type GoodCategory struct {
	global.GVA_MODEL
	MerchantId   uint   `json:"merchantId" gorm:"merchant_id"` //商家ID
	CategoryId   uint   `json:"categoryId" gorm:"category_id comment:'分类id'"`
	CategoryName string `json:"categoryName" gorm:"category_name comment:'商品分类名称'"`
}
