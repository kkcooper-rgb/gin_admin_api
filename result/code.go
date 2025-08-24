package result

// 定义状态
type Codes struct {
	Message           map[uint]string
	Success           uint
	Failed            uint
	SysMenuIsExist    uint
	DelSysMenuFailed  uint
	NorDeleteMenu     uint
	RoleAlreadyExists uint
	DelSysRoleFailed  uint
}

// 状态码
var ApiCode = &Codes{
	Success:           200,
	Failed:            501,
	SysMenuIsExist:    600,
	DelSysMenuFailed:  601,
	NorDeleteMenu:     602,
	RoleAlreadyExists: 603,
	DelSysRoleFailed:  604,
}

// 状态信息初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Success:           "成功",
		ApiCode.Failed:            "失败",
		ApiCode.SysMenuIsExist:    "菜单名称已存在，请重新输入！",
		ApiCode.DelSysMenuFailed:  "菜单已分配不能删除",
		ApiCode.NorDeleteMenu:     "存在子菜单不能删除",
		ApiCode.RoleAlreadyExists: "角色名称或角色key已存在，请重新输入！",
		ApiCode.DelSysRoleFailed:  "角色已分配不能删除",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
