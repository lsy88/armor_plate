package request

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// LoginForm 用户登录请求结构
type ReqLoginForm struct {
	EmployeeID     uint64 `json:"employee_id" binding:"required"`
	EmployeeName   string `json:"employee_name" binding:"required"`
	Password       string `json:"password" binding:"required"`
	AuthorityId    uint   `json:"authority_id" gorm:"authority_id"`
	DepartmentName string `json:"department_name" binding:"required"`
}

// RegisterForm 用户注册请求结构
type ReqRegisterForm struct {
	EmployeeName    string `json:"employee_name" binding:"required"`
	Password        string `json:"password" binding:"required"`
	AuthorityId     uint   `json:"authority_id"`
	ConfirmPassword string `json:"confirm_password" binding:"required"` //确认密码
	DepartmentName  string `json:"department_name" binding:"required"`
	Email           string `json:"email"`
	PhoneNumber     uint64 `json:"phone_number"`
}

//修改用户信息
type ReqChangeEmployeeReq struct {
	ID          uint64 `json:"-"` //从jwt中提取EmployeeID ,避免越权
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password"`
	Email       string `json:"email"`
	PhoneNumber uint64 `json:"phone_number"`
}

//分页信息
type PageInfo struct {
	Page    int    `json:"page" form:"page"`       //页码
	Size    int    `json:"size" form:"size"`       //页面大小
	Keyword string `json:"keyword" form:"keyword"` //关键字
}

// SetUserAuth 设置用户角色

type ReqSetUserAuthorities struct {
	EID          uint64 `json:"eid"`
	AuthorityIds []uint `json:"authority_ids"`
}
type SetUserAuth struct {
	EID         uint64 `json:"eid"`
	AuthorityId uint   `json:"authority_id"`
}

// SetRoleInfo 设置角色拥有权限
type ReqSetRoleInfo struct {
	AuthorityId   uint   `json:"authority_id"`   //角色id
	AuthorityName string `json:"authority_name"` //角色名称
	Path          string `json:"path"`           //请求路径
	Method        string `json:"method"`         //请求方法[post][get][put]...
}

//新增货物
type ReqCreateProduct struct {
	ProductName string  `json:"product_name" binding:"required"`
	CategoryID  uint    `json:"category_id" binding:"required"` //物品种类id
	Info        string  `json:"info"`                           //货物信息描述
	OnStorage   uint    `json:"onstorage"`                      //货物状态
	Price       float64 `json:"price"`                          //货物价格
}

//给货物设置存储位置
type ReqProductPath struct {
	ProductName string `json:"product_name"`
	DepotName   string `json:"depot_name"`
	Count       int    `json:"count"`
}

//获取分页参数
func GetPageInfo(c *gin.Context) (int, int) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	
	var (
		page int
		size int
	)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	return page, size
}

//请求订单信息
type ReqOrderMes struct {
	CustomerName string `json:"customer_name"` //订单客户
	ProductName  string `json:"product_name"`  //货物名称
	Amount       int    `json:"amount"`        //需求数量
	State        int    `json:"state"`         //订单支付状态已支付,未支付,待支付
	Destination  string `json:"destination"`   //目的地址
}

//添加仓库信息请求
type ReqDepotMes struct {
	DepotName string `json:"depot_name"`                         //仓库名
	Path      string `json:"path" binding:"required"`            //地址
	Priority  int    `json:"priority" gorm:"priority;default:0"` //优先级
}

//交付订单请求
type ReqDeliveryOrder struct {
	OrderID uint64  `json:"order_id"` //要交付的订单号
	Money   float64 `json:"money"`    //交付金额
}
