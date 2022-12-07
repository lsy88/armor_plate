package model

// Order 订单单据
type Order struct {
	MODEL
	OrderID      uint64  `json:"order_id" gorm:"order_id"` //订单号
	EmployeeID   uint64  `json:"employee_id" gorm:"employee_id;not null"`
	ProductID    uint64  `json:"product_id" gorm:"product_id;not null"`
	CustomerName string  `json:"customer_name" gorm:"customer_name"`         //客户信息
	Amount       int     `json:"amount" gorm:"amount"`                       //数量
	State        int     `json:"state" gorm:"default:0;comment:0-未支付 1-已支付"` //订单状态
	Money        float64 `json:"money" gorm:"money"`                         //订单金额
	Destination  string  `json:"destination" gorm:"destination"`             //目的地址
}

func (Order) TableName() string {
	return "ap_order"
}

// OutgoingOrder 出库单
type OutgoingOrder struct {
	MODEL
	DeliveryID   uint64 `json:"delivery_id" gorm:"delivery_id"` //出库单号
	OrderID      uint64 `json:"order_id" gorm:"order_id"`       //订单号
	EmployeeID   uint64 `json:"employee_id" gorm:"employee_id;not null"`
	DepotID      uint64 `json:"depot_id" gorm:"depot_id"`
	ProductID    uint64 `json:"product_id" gorm:"product_id;not null"`
	CustomerName string `json:"customer_name" gorm:"customer_name"` //客户信息
	
	Amount      int     `json:"amount" gorm:"amount"`           //出库数量
	Money       float64 `json:"money" gorm:"money"`             //订单金额
	Destination string  `json:"destination" gorm:"destination"` //目的地址
	Path        string  `json:"path" gorm:"path"`               //发货地址
}

func (OutgoingOrder) TableName() string {
	return "ap_outgoing_order"
}
