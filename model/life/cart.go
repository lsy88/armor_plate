package life

//购物车添加参数模型
type CartAddParam struct {
	UserId     int `form:"userId"`
	GoodId     int `form:"goodId"`
	MerchantId int `form:"merchantId"`
	GoodCount  int `form:"goodCount"`
}

//购物车删除参数
type CartDeleteParam struct {
	UserId     int `form:"userId"`
	GoodId     int `form:"goodId"`
	MerchantId int `form:"merchantId"`
}

//购物车清除参数模型
type CartClearParam struct {
	UserId string `form:"userId"`
}
