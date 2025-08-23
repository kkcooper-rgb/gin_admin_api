package result

// 定义状态
type Codes struct {
	Message        map[uint]string
	Success        uint
	Failed         uint
	SysMenuIsExist uint
}

// 状态码
var ApiCode = &Codes{
	Success:        200,
	Failed:         501,
	SysMenuIsExist: 600,
}

// 状态信息初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Success:        "成功",
		ApiCode.Failed:         "失败",
		ApiCode.SysMenuIsExist: "菜单名称已存在，请重新输入！",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
