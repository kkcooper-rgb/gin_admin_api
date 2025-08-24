package api

import (
	"go_admin_api/core"
	"go_admin_api/model"
	"go_admin_api/result"
	"go_admin_api/utils"
	"strconv"
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

// GetSysRoleList 分页查询角色列表
// @Summary 分页查询角色列表
// @Tags 角色相关接口
// @Produce json
// @Description 分页查询角色列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数量"
// @Param roleName query string false "角色名称"
// @Param status query string false "帐号启用状态: 1->启用,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @Router /api/sysRole/list [get]
func GetSysRoleList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	RoleName := c.Query("roleName")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var sysRole []model.SysRole
	var count int64
	curDb := core.Db.Table("sys_role")
	if RoleName != "" {
		curDb = curDb.Where("role_name = ?", RoleName)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if Status != "" {
		curDb = curDb.Where("status = ?", Status)
	}
	curDb.Count(&count).Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time DESC").Find(&sysRole)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysRole})
}

// GetSysRole 根据id查询角色
// @Summary 根据id查询角色
// @Tags 角色相关接口
// @Produce json
// @Description 根据id查询角色
// @Param id query int true "角色id"
// @Success 200 {object} result.Result
// @Router /api/sysRole/info [get]
func GetSysRole(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var sysRole model.SysRole
	core.Db.First(&sysRole, Id)
	result.Success(c, sysRole)
}

// UpdateSysRole 修改角色
// @Summary 修改角色
// @Tags 角色相关接口
// @Produce json
// @Description 修改角色
// @Param data body model.UpdateSysRoleDto true "data"
// @Success 200 {object} result.Result
// @Router /api/sysRole/update [put]
func UpdateSysRole(c *gin.Context) {
	var dto model.UpdateSysRoleDto
	_ = c.BindJSON(&dto)
	var sysRole model.SysRole
	core.Db.First(&sysRole, dto.ID)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	core.Db.Save(&sysRole)
	result.Success(c, true)
}

// DeleteSysRole 根据id删除角色
// @Summary 根据id删除角色
// @Tags 菜单相关接口
// @Produce json
// @Description 根据id删除角色
// @Param data body model.SysRoleIdDto true "data"
// @Success 200 {object} result.Result
// @Router /api/sysRole/delete [delete]
func DeleteSysRole(c *gin.Context) {
	var dto model.SysRoleIdDto
	_ = c.BindJSON(&dto)
	sysRoleById := GetSysAdminRole(dto.ID)
	if sysRoleById.RoleId > 0 {
		result.Failed(c, int(result.ApiCode.DelSysRoleFailed), result.ApiCode.GetMessage(result.ApiCode.DelSysRoleFailed))
		return
	}
	core.Db.Table("sys_role").Delete(&model.SysRole{}, dto.ID)
	core.Db.Table("sys_role_menu").Where("role_id = ?", dto.ID).Delete(&model.SysRoleMenu{})
	result.Success(c, true)
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

// GetSysAdminRole 查询是否分配角色
func GetSysAdminRole(id uint) (sysAdminRole model.SysAdminRole) {
	core.Db.Where("role_id = ?", id).First(&sysAdminRole)
	return sysAdminRole
}
