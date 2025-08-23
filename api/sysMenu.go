package api

import (
	"go_admin_api/core"
	"go_admin_api/model"
	"go_admin_api/result"
	"go_admin_api/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateSysMenu 新增菜单
// @Summary 新增菜单
// @Tags 菜单相关接口
// @Produce json
// @Description 新增菜单
// @Param data body model.AddSysMenuDto true "data"
// @Success 200 {object} result.Result
// @Router /api/sysMenu/add [post]
func CreateSysMenu(c *gin.Context) {
	// 绑定请求参数到DTO
	var dto model.AddSysMenuDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		// 参数绑定失败返回错误
		result.Failed(c, int(result.ApiCode.Failed), result.ApiCode.GetMessage(result.ApiCode.Failed))
		return
	}

	// 检查菜单名称是否已存在
	sysMenuByName := GetSysMenuByName(dto.MenuName)
	if sysMenuByName.ID != 0 {
		// 菜单已存在返回错误
		result.Failed(c, int(result.ApiCode.SysMenuIsExist), result.ApiCode.GetMessage(result.ApiCode.SysMenuIsExist))
		return
	}

	// 根据菜单类型(1-目录, 2-菜单, 3-按钮)执行不同创建逻辑
	switch dto.MenuType {
	case 1: // 目录类型
		sysMenu := model.SysMenu{
			ParentId:   0, // 目录没有父级
			MenuName:   dto.MenuName,
			Icon:       dto.Icon,
			MenuType:   dto.MenuType,
			Url:        dto.Url,
			MenuStatus: dto.MenuStatus,
			Sort:       dto.Sort,
			CreateTime: utils.HTime{Time: time.Now()},
		}
		core.Db.Create(&sysMenu)
	case 2: // 菜单类型
		sysMenu := model.SysMenu{
			ParentId:   dto.ParentId,
			MenuName:   dto.MenuName,
			Icon:       dto.Icon,
			MenuType:   dto.MenuType,
			Url:        dto.Url,
			Value:      dto.Value,
			MenuStatus: dto.MenuStatus,
			Sort:       dto.Sort,
			CreateTime: utils.HTime{Time: time.Now()},
		}
		core.Db.Create(&sysMenu)
	case 3: // 按钮类型
		sysMenu := model.SysMenu{
			ParentId:   dto.ParentId,
			MenuName:   dto.MenuName,
			MenuType:   dto.MenuType,
			Value:      dto.Value,
			MenuStatus: dto.MenuStatus,
			Sort:       dto.Sort,
			CreateTime: utils.HTime{Time: time.Now()},
		}
		core.Db.Create(&sysMenu)
	}

	// 新增成功返回
	result.Success(c, true)
}

// GetSysMenuByName 根据菜单名称查询菜单是否存在
func GetSysMenuByName(menuName string) (sysMenu model.SysMenu) {
	core.Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}
