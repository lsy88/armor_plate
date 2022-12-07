package mysql

import (
	"armor_plate/model"
	"armor_plate/pkg/md5"
	"errors"
	"gorm.io/gorm"
)

//通过部门名称获取部门id
func GetDepIDByName(name string) (uint, error) {
	var u model.Department
	err := db.Where("department_name = ?", name).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, errors.New("该部门不存在,请重新确定部门")
		}
		return 0, err
	}
	return u.DepartmentID, nil
}

//用户登录

func Login(user *model.Employee) (err error) {
	originPassword := user.Password
	var use model.Employee
	u := db.Where("employee_id = ? and department_name = ?", user.EmployeeID, user.DepartmentName).
		Find(&use)
	if u.RowsAffected == 0 {
		//用户不存在
		return ErrorUserNotExit
	}
	db.Table("ap_employee").Select("password").
		Where("employee_name", user.EmployeeName).Find(&use)
	password := md5.MD5([]byte(originPassword))
	if use.Password != password {
		return ErrorPasswordWrong
	}
	return nil
}

//用户注册
func Register(user *model.Employee) (err error) {
	var use model.Employee
	u := db.Model(&model.Employee{}).Where("employee_name = ? and department_name = ?", user.EmployeeName, user.DepartmentName).
		Find(&use)
	if u.RowsAffected != 0 {
		//用户存在
		return ErrorUserExit
	}
	//密码加密
	user.Password = md5.MD5([]byte(user.Password))
	err = db.Model(&model.Employee{}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// Update 用户修改信息
func Update(user *model.Employee) (err error) {
	var u model.Employee
	err = db.Model(&model.Employee{}).Where("employee_id = ?", user.EmployeeID).First(&u).Error
	if err == gorm.ErrRecordNotFound {
		return ErrorUserNotExit
	}
	err = db.Model(&model.Employee{}).Save(&user).Error
	return err
}

// GetList 查询用户列表
func GetList(limit, offset int) (list []*model.Employee, total int64, err error) {
	err = db.Table("ap_employee").Offset(offset).Limit(limit).Count(&total).Find(&list).Error
	return
}

//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: eid uint64, authorityId uint
//@return: err error

func SetUserAuthority(eid uint64, authorityId uint) (err error) {
	err = db.Where("employee_id = ? and authority_id = ?", eid, authorityId).First(&model.UserAuthority{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	//err = db.Preload("ap_employee").Preload("ap_casbin_model").Model(&model.UserAuthority{}).Where("employee_id = ?", eid).First(&model.Employee{}).
	//	Update("authority_id", authorityId).Error
	var u = model.UserAuthority{
		EmployeeID:  eid,
		AuthorityId: authorityId,
	}
	err = db.Model(&model.UserAuthority{}).Save(&u).Error
	return err
}

// SetUserAuthorities 设置用户权限
//@function: SetUserAuth
//@description: 设置一个用户的权限
//@param: eid uint, authIds []string
//@return: error
func SetUserAuthorities(eid uint64, authIds []uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.UserAuthority{}).Delete("employee_id = ?", eid).Error
		if err != nil {
			return err
		}
		var ua []model.UserAuthority
		for _, v := range authIds {
			ua = append(ua, model.UserAuthority{
				EmployeeID:  eid,
				AuthorityId: v,
			})
		}
		err = tx.Model(&model.UserAuthority{}).Create(&ua).Error
		if err != nil {
			return err
		}
		err = tx.Where("employee_id = ?", eid).First(&model.Employee{}).Update("authority_id", authIds[0]).Error
		if err != nil {
			return err
		}
		return nil
	})
	
}

//@function: CreateAuth
//@description: 创建一个角色
//@param: auth *model.CasbinModel
//@return: authority *model.CasbinModel, err error

func CreateAuth(auth *model.CasbinModel) (authority *model.CasbinModel, err error) {
	err = db.Where("authority_id = ?", auth.AuthorityId).First(&model.CasbinModel{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, errors.New("存在相同角色id")
	}
	
	err = db.Create(&auth).Error
	return auth, err
}
