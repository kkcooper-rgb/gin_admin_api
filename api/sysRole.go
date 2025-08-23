package api

import (
	"go_admin_api/core"
	"go_admin_api/model"
	"go_admin_api/result"
	"go_admin_api/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateSysRole 新增角色
// @Summary 新增角色
// @Tags 菜单相关接口
// @Produce json
// @Description 新增角色
// @Param data body model.AddSysRoleDto true "data"
// @Success 200 {object} result.Result
// @Router /api/sysRole/add [post]
func CreateSysRole(c *gin.Context) {
	var dto model.AddSysRoleDto
	_ = c.BindJSON(&dto)
	sysRoleByName := GetSysRoleByName(dto.RoleName)
	if sysRoleByName.ID > 0 {
		result.Failed(c, int(result.ApiCode.RoleAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.RoleAlreadyExists))
		return
	}
	sysRoleByKey := GetSysRoleByKey(dto.RoleKey)
	if sysRoleByKey.ID > 0 {
		result.Failed(c, int(result.ApiCode.RoleAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.RoleAlreadyExists))
		return
	}
	addSysRole := model.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Description: dto.Description,
		Status:      dto.Status,
		CreateTime:  utils.HTime{Time: time.Now()},
	}
	tx := core.Db.Create(&addSysRole)
	if tx.RowsAffected > 0 {
		result.Success(c, true)
		return
	} else {
		result.Failed(c, int(result.ApiCode.Failed), result.ApiCode.GetMessage(result.ApiCode.Failed))
		return
	}
}

// GetSysRoleByName 根据角色名称查询角色
func GetSysRoleByName(roleName string) (sysRole model.SysRole) {
	core.Db.Where("role_name = ?", roleName).First(&sysRole)
	return sysRole
}

// GetSysRoleByKey 根据角色key名称查询角色
func GetSysRoleByKey(roleKey string) (sysRole model.SysRole) {
	core.Db.Where("role_key = ?", roleKey).First(&sysRole)
	return sysRole
}
