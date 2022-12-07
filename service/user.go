package service

import (
	"armor_plate/controller/request"
	"armor_plate/dao/mysql"
	"armor_plate/model"
	"armor_plate/pkg/md5"
	"armor_plate/pkg/snowflake"
	"errors"
	"go.uber.org/zap"
)

func LoginUser(u *request.ReqLoginForm) (err error) {
	user := &model.Employee{
		EmployeeID:     u.EmployeeID,
		EmployeeName:   u.EmployeeName,
		Password:       u.Password,
		AuthorityId:    u.AuthorityId,
		DepartmentName: u.DepartmentName,
	}
	if err = mysql.Login(user); err != nil {
		zap.L().Error("mysql login failed, ", zap.Error(err))
		return err
	}
	return nil
}

func RegisterUser(user *request.ReqRegisterForm) (err error) {
	id, _ := snowflake.GetID()
	depId, err := mysql.GetDepIDByName(user.DepartmentName)
	if err != nil {
		zap.L().Error("department is not exist", zap.Error(err))
		return errors.New("您所选的部门不存在，请重新选择")
	}
	u := model.Employee{
		EmployeeID:   id,
		EmployeeName: user.EmployeeName,
		AuthorityId:  user.AuthorityId,
		//Role:           user.Role,
		Password:       user.Password,
		PhoneNumber:    user.PhoneNumber,
		Email:          user.Email,
		DepartmentID:   depId,
		DepartmentName: user.DepartmentName,
	}
	if err = mysql.Register(&u); err != nil {
		zap.L().Error("mysql register failed", zap.Error(err))
		return
	}
	//创建成功
	return nil
}

func UpdateUser(user *request.ReqChangeEmployeeReq, id uint64) (err error) {
	if user.Password == user.NewPassword {
		return errors.New("新密码与旧密码一致,请重新输入新密码")
	}
	u := model.Employee{
		EmployeeID:  id,
		Password:    md5.MD5([]byte(user.Password)),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}
	if err = mysql.Update(&u); err != nil {
		return err
	}
	return nil
}

func GetUserList(info *request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.Size
	offset := info.Size * (info.Page - 1)
	list, total, err = mysql.GetList(limit, offset)
	return
}

func SetUserAuth(id uint64, authId uint) (err error) {
	if err = mysql.SetUserAuthority(id, authId); err != nil {
		return err
	}
	return nil
	
}

func CreateUserAuth(s *request.ReqSetRoleInfo) (auth *model.CasbinModel, err error) {
	cm := model.CasbinModel{
		AuthorityId: s.AuthorityId,
		Role:        s.AuthorityName,
		CasbinInfo: model.CasbinInfo{
			Path:   s.Path,
			Method: s.Method,
		},
	}
	auth, err = mysql.CreateAuth(&cm)
	return auth, err
}
