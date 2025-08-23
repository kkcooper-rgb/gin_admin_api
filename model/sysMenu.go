package model

import "go_admin_api/utils"

// SysMenu 系统菜单结构体，对应数据库菜单表
type SysMenu struct {
	ID         uint        `gorm:"column:id;primaryKey;NOT NULL" json:"id" comment:"主键"`
	ParentId   uint        `gorm:"column:parent_id" json:"parentId" comment:"父菜单id"`
	MenuName   string      `gorm:"column:menu_name;type:varchar(100)" json:"menuName" comment:"菜单名称"`
	Icon       string      `gorm:"column:icon;type:varchar(100)" json:"icon" comment:"菜单图标"`
	Value      string      `gorm:"column:value;type:varchar(100)" json:"value" comment:"权限值"`
	MenuType   uint        `gorm:"column:menu_type" json:"menuType" comment:"菜单类型：1->目录；2->菜单；3->按钮"`
	Url        string      `gorm:"column:url;type:varchar(100)" json:"url" comment:"菜单url"`
	MenuStatus uint        `gorm:"column:menu_status" json:"menuStatus" comment:"启用状态：1->启用；2->禁用"`
	Sort       uint        `gorm:"column:sort" json:"sort" comment:"排序"`
	CreateTime utils.HTime `gorm:"column:create_time" json:"createTime" comment:"创建时间"`
	Children   []SysMenu   `gorm:"-" json:"children" comment:"子集"` // 忽略数据库映射，用于树形结构展示
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

// AddSysMenuDto 新增菜单参数
type AddSysMenuDto struct {
	ParentId   uint   `json:"parentId"`   // 父菜单id
	MenuName   string `json:"menuName"`   // 菜单名称
	Icon       string `json:"icon"`       // 菜单图标
	Value      string `json:"value"`      // 权限值
	MenuType   uint   `json:"menuType"`   // 菜单类型：1->目录；2->菜单；3->按钮
	Url        string `json:"url"`        // 菜单URL
	MenuStatus uint   `json:"menuStatus"` // 启用状态：1->启用；2->禁用
	Sort       uint   `json:"sort"`       // 排序
}

// UpdateSysMenuDto 修改菜单参数
type UpdateSysMenuDto struct {
	ID         uint   `json:"id"`         // ID
	ParentId   uint   `json:"parentId"`   // 父菜单id
	MenuName   string `json:"menuName"`   // 菜单名称
	Icon       string `json:"icon"`       // 菜单图标
	Value      string `json:"value"`      // 权限值
	MenuType   uint   `json:"menuType"`   // 菜单类型：1->目录；2->菜单；3->按钮
	Url        string `json:"url"`        // 菜单url
	MenuStatus uint   `json:"menuStatus"` // 启用状态；1->启用；2->禁用
	Sort       uint   `json:"sort"`       // 排序
}

// SysMenuIdDto id参数
type SysMenuIdDto struct {
	ID uint `json:"id"` // ID
}
