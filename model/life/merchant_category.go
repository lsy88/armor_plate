package life

import "github.com/flipped-aurora/gin-vue-admin/server/global"

//商家分类表
type MerchantCategory struct {
	global.GVA_MODEL
	CategoryName string `json:"categoryName" gorm:"category_name"` //商家分类名
}
