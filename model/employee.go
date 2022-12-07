package model

import (
	"encoding/json"
	"errors"
)

type Employee struct {
	MODEL
	EmployeeID   uint64        `json:"employee_id" gorm:"employee_id"`
	EmployeeName string        `json:"employee_name" gorm:"employee_name"`
	Password     string        `json:"password" gorm:"password"`
	AuthorityId  uint          `json:"authority_id" gorm:"authority_id"`
	CasbinModel  CasbinModel   `json:"authority" gorm:"foreignKey:AuthorityId;association_foreignKey:AuthorityId;comment:用户角色ID"`
	CasbinModels []CasbinModel `json:"-" gorm:"many2many:ap_user_authority;"`
	//Role           string `json:"role" gorm:"role;default:common;comment:common-用户,root-管理员" binding:"required"` //员工角色
	DepartmentID   uint   `json:"department_id" gorm:"department_id"`
	DepartmentName string `json:"department_name" gorm:"department_name"`
	//Department     Department `json:"department" gorm:"foreignKey:DepartmentID;"` //部门
	Email       string `json:"email" gorm:"email"`
	PhoneNumber uint64 `json:"phone_number" gorm:"phone;type:varchar(11)"`
}

func (Employee) TableName() string {
	return "ap_employee"
}

func (e *Employee) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName    string `json:"username" gorm:"username"`
		Password    string `json:"password" gorm:"password"`
		AuthorityId uint   `json:"authority_id"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		e.EmployeeName = required.UserName
		e.Password = required.Password
		e.AuthorityId = required.AuthorityId
	}
	return
}
