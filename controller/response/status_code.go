package response

type StatusCode int64

const (
	CodeSuccess         StatusCode = 1000
	CodeInvalidParams   StatusCode = 1001
	CodeUserExist       StatusCode = 1002
	CodeUserNotExist    StatusCode = 1003
	CodeInvalidPassword StatusCode = 1004
	CodeServerBusy      StatusCode = 1005
	
	CodeInvalidToken      StatusCode = 1006
	CodeInvalidAuthFormat StatusCode = 1007
	CodeNotLogin          StatusCode = 1008
	
	CodeNoAuthority    StatusCode = 1009
	CodeAuthCreateFail StatusCode = 1010
	CodeNoAuthLogin    StatusCode = 1011
)

var MessageFlag = map[StatusCode]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名重复",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	
	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
	//casbin
	CodeNoAuthority:    "用户权限不足",
	CodeAuthCreateFail: "创建权限失败",
	CodeNoAuthLogin:    "非法访问",
}

func (s StatusCode) Decode() string {
	message, ok := MessageFlag[s]
	if ok {
		//存在直接返回
		return message
	}
	//不存在返回服务忙
	return MessageFlag[CodeServerBusy]
}
