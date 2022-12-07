package model

type UserAuthority struct {
	MODEL
	EmployeeID  uint64 `json:"employee_id" gorm:"employee_id"`
	AuthorityId uint   `json:"authority_id" gorm:"authority_id"`
}

func (s *UserAuthority) TableName() string {
	return "ap_user_authority"
}
