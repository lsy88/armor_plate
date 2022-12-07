package life

import "github.com/flipped-aurora/gin-vue-admin/server/global"

//商家信息表
type Merchant struct {
	global.GVA_MODEL
	MerchantName  string `json:"merchantName" gorm:"merchant_name"`
	BrandImages   string `json:"brandImages" gorm:"brand_images comment:'招牌图片'"`
	Description   string `json:"description" gorm:"description size:200 comment:'商家简介'"`
	Phone         string `json:"phone" gorm:"phone"` //联系方式
	Address       string `json:"address" gorm:"address"`
	MerCategoryId uint   `json:"merCategoryId" gorm:"comment:'商家类别id'"`
}
