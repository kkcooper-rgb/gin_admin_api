package model

import (
	"go_admin_api/utils"
)

// SysRole 角色模型
type SysRole struct {
	ID          uint        `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                         // ID
	RoleName    string      `gorm:"column:role_name;varchar(64);comment:'角色名称';NOT NULL" json:"roleName"`         // 角色名称
	RoleKey     string      `gorm:"column:role_key;varchar(64);comment:'权限字符串';NOT NULL" json:"roleKey"`          // 权限字符串
	Status      int         `gorm:"column:status;default:1;comment:'帐号启用状态: 1->启用,2->禁用';NOT NULL" json:"status"` // 帐号启用状态: 1->启用,2->禁用
	Description string      `gorm:"column:description;varchar(500);comment:'描述'" json:"description"`              // 描述
	CreateTime  utils.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                 // 创建时间
}

func (SysRole) TableName() string {
	return "sys_role"
}

type AddSysRoleDto struct {
	RoleName    string      `json:"roleName"`    // 角色名称
	RoleKey     string      `json:"roleKey"`     // 权限字符串
	Status      int         `json:"status"`      // 帐号启用状态: 1->启用,2->禁用
	Description string      `json:"description"` // 描述
	CreateTime  utils.HTime `json:"createTime"`  // 创建时间
}
