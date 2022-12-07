package model

type CasbinInfo struct {
	Path   string `json:"path" gorm:"v1"`
	Method string `json:"method" gorm:"v2"`
}

type CasbinModel struct {
	MODEL
	AuthorityId uint       `json:"authority_id" gorm:"authority_id"` //用户角色id
	Role        string     `json:"role" gorm:"v0"`                   //用户角色
	CasbinInfo  CasbinInfo `json:"casbinInfo" gorm:"-"`
	Users       []Employee `json:"-" gorm:"many2many:ap_user_authority;"`
}

func (CasbinModel) TableName() string {
	return "ap_casbin_model"
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/user/login", Method: "POST"},
		{Path: "/user/register", Method: "POST"},
		{Path: "/user/updateUser", Method: "POST"},
		{Path: "/user/setUserAuth", Method: "POST"},
		{Path: "/user/getUserList", Method: "GET"},
	}
}
