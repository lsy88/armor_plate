package casbin

import (
	"armor_plate/model"
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"strconv"
)

type CasbinService struct{}

var Enforcer *casbin.Enforcer
var CasbinServiceApp = new(CasbinService)

// Casbin 持久化到数据库
func (casbinServer *CasbinService) Casbin() *casbin.Enforcer {
	//初始化一个gorm适配器
	a, err := gormadapter.NewAdapter("mysql", "root:716523@tcp(127.0.0.1:3306)/armor_plate", true)
	if err != nil {
		zap.L().Error("casbin adaptor failed init", zap.Error(err))
	}
	e, err := casbin.NewEnforcer("conf/rbac_model.conf", a)
	if err != nil {
		zap.L().Error("casbin adaptor failed init", zap.Error(err))
	}
	//从数据库加载策略
	e.LoadPolicy()
	Enforcer = e
	return Enforcer
}

// ClearCasbin 清除匹配的权限
func (casbinServer *CasbinService) ClearCasbin(v int, p ...string) bool {
	success, err := Enforcer.RemoveFilteredPolicy(v, p...)
	if err != nil {
		zap.L().Error("clear casbin failed", zap.Error(err))
	}
	return success
}

// UpdateCasbin 添加权限
func (casbinServer *CasbinService) UpdateCasbin(AuthorityId uint, casbinInfo []model.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityId))
	casbinServer.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfo {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	//添加策略
	success, _ := Enforcer.AddPolicies(rules)
	if !success {
		return errors.New("存在相同的api,添加失败，请联系管理员")
	}
	return nil
}

// GetPolicyByAuthorityIdList 获取权限列表
func (casbinServer *CasbinService) GetPolicyByAuthorityIdList(AuthorityId uint) (pathMap []model.CasbinInfo) {
	e := Enforcer
	autherrityId := strconv.Itoa(int(AuthorityId))
	//GetFilteredPolicy获取筛选策略中获取策略中的所有授权规则，可以指定字段筛选器
	list := e.GetFilteredPolicy(0, autherrityId)
	for _, v := range list {
		pathMap = append(pathMap, model.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return
}
