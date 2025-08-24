package model

// SysAdminRole 用户与角色关系模型
type SysAdminRole struct {
	RoleId  uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"roleId"`   // 角色id
	AdminId uint `gorm:"column:admin_id;comment:'菜单id';NOT NULL" json:"adminId"` // 角色id
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}
