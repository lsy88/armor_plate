package model

//部门
type Department struct {
	MODEL
	DepartmentID   uint   `json:"department_id" gorm:"department_id"`
	DepartmentName string `json:"department_name" gorm:"department_name;unique"`
	LeaderID       uint64 `json:"leader_id" gorm:"leader_id"` //部门主管ID
	//Employee       *Employee `json:"employee" gorm:"foreignKey:DepartmentID;"`
}

func (Department) TableName() string {
	return "ap_department"
}
