package api

import (
	"go_admin_api/core"
	"go_admin_api/model"
	"go_admin_api/result"
	"go_admin_api/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateSysAdmin 新增用户
// @Summary 新增用户
// @Tags 用户相关接口
// @Produce json
// @Description 新增用户
// @Param data body model.AddSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/add [post]
func CreateSysAdmin(c *gin.Context) {
	// 绑定参数并校验参数必填
	var dto model.AddSysAdminDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 用户名称不能重复
	sysAdminByUsername := GetSysAdminByUsername(dto.Username)
	if sysAdminByUsername.ID > 0 {
		result.Failed(c, int(result.ApiCode.UsernameAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.UsernameAlreadyExists))
		return
	}
	// 新增用户
	sysAdmin := model.SysAdmin{
		Username:   dto.Username,
		Nickname:   dto.Nickname,
		Password:   utils.EncryptionMd5(dto.Password),
		Phone:      dto.Phone,
		Email:      dto.Email,
		Sex:        dto.Sex,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: utils.HTime{Time: time.Now()},
	}
	core.Db.Create(&sysAdmin)
	// 增加用户角色表信息
	sysAdminExist := GetSysAdminByUsername(dto.Username)
	var sysAdminRole model.SysAdminRole
	sysAdminRole.AdminId = sysAdminExist.ID
	sysAdminRole.RoleId = dto.RoleId
	core.Db.Create(&sysAdminRole)
	result.Success(c, true)
}

// GetSysAdmin 根据id查询用户
// @Summary 根据id查询用户
// @Tags 用户相关接口
// @Produce json
// @Description 根据id查询用户
// @Param id query int true "用户id"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/info [get]
func GetSysAdmin(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var sysAdminInfo model.SysAdminInfo
	core.Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_admin_role.role_id = sys_role.id").
		First(&sysAdminInfo, Id)
	result.Success(c, sysAdminInfo)
}

// UpdateSysAdmin 修改用户
// @Summary 修改用户
// @Tags 用户相关接口
// @Produce json
// @Description 修改用户
// @Param data body model.UpdateSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/update [put]
func UpdateSysAdmin(c *gin.Context) {
	// 绑定参数并校验参数必填
	var dto model.UpdateSysAdminDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 先查询再修改
	var sysAdmin model.SysAdmin
	core.Db.First(&sysAdmin, dto.Id)
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	if dto.Note != "" {
		sysAdmin.Note = dto.Note
	}
	sysAdmin.Status = dto.Status
	sysAdmin.Sex = dto.Sex
	core.Db.Save(&sysAdmin)
	// 删除之前的角色在分配新的角色
	var sysAdminRole model.SysAdminRole
	core.Db.Where("admin_id = ?", dto.Id).Delete(&model.SysAdminRole{})
	sysAdminRole.AdminId = dto.Id
	sysAdminRole.RoleId = dto.RoleId
	core.Db.Create(&sysAdminRole)
	result.Success(c, true)
}

// DeleteSysAdmin 根据id删除用户
// @Summary 根据id删除用户
// @Tags 用户相关接口
// @Produce json
// @Description 根据id删除用户
// @Param data body model.SysAdminIDDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/delete [delete]
func DeleteSysAdmin(c *gin.Context) {
	var dto model.SysAdminIDDto
	_ = c.BindJSON(&dto)
	core.Db.Delete(&model.SysAdmin{}, dto.Id)
	core.Db.Where("admin_id = ?", dto.Id).Delete(&model.SysAdminRole{})
	result.Success(c, true)
}

// UpdateSysAdminStatus 设置用户状态
// @Summary 设置用户状态
// @Tags 用户相关接口
// @Produce json
// @Description 设置用户状态
// @Param data body model.UpdateSysAdminStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/updateStatus [put]
func UpdateSysAdminStatus(c *gin.Context) {
	var dto model.UpdateSysAdminStatusDto
	_ = c.BindJSON(&dto)
	var sysAdmin model.SysAdmin
	core.Db.First(&sysAdmin, dto.Id)
	sysAdmin.Status = dto.Status
	core.Db.Save(&sysAdmin)
	result.Success(c, true)
}

// ResetSysAdminPassword 重置密码
// @Summary 重置密码
// @Tags 用户相关接口
// @Produce json
// @Description 重置密码
// @Param data body model.ResetSysAdminPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/updatePassword [put]
func ResetSysAdminPassword(c *gin.Context) {
	var dto model.ResetSysAdminPasswordDto
	_ = c.BindJSON(&dto)
	var sysAdmin model.SysAdmin
	core.Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = utils.EncryptionMd5(dto.Password)
	core.Db.Save(&sysAdmin)
	result.Success(c, true)
}

// GetSysAdminList 分页查询用户列表
// @Summary 分页查询用户列表
// @Tags 用户相关接口
// @Produce json
// @Description 分页查询用户列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param username query string false "用户名称"
// @Param status query string false "帐号启用状态: 1->启用,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/list [get]
func GetSysAdminList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var sysAdminVo []model.SysAdminVo
	var count int64
	curDb := core.Db.Table("sys_admin").
		Select("sys_admin.*, sys_role.role_name").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_role.id = sys_admin_role.role_id")
	if Username != "" {
		curDb = curDb.Where("sys_admin.sys_admin.username = ?", Username)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("sys_admin.create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if Status != "" {
		curDb = curDb.Where("sys_admin.status = ?", Status)
	}
	curDb.Count(&count).Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("sys_admin.create_time DESC").Find(&sysAdminVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysAdminVo})
}

// GetSysAdminByUsername 根据用户名称查询用户
func GetSysAdminByUsername(username string) (sysAdmin model.SysAdmin) {
	core.Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}
