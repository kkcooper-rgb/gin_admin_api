package core

import (
	"go_admin_api/config"
	"go_admin_api/model"
	"time"

	"github.com/golang-jwt/jwt/v4" // 更换为维护中的官方分支
)

// 用户自定义Claims，包含用户信息和标准声明
type userStdClaims struct {
	model.JwtAdmin       // 嵌入用户信息结构体
	jwt.RegisteredClaims // 使用v4版本的标准声明（替代原StandardClaims）
}

var (
	// 从配置中获取token过期时间（小时）
	TokenExpireDuration = time.Duration(config.Config.Token.ExpireTime) * time.Hour
	// 从配置中获取签名密钥
	Secret = []byte(config.Config.Token.Secret)
	// 从配置中获取签发人
	Issuer = config.Config.Token.Issuer

	// 错误信息常量
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
)

// GenerateTokenByAdmin 根据用户信息生成JWT令牌
func GenerateTokenByAdmin(admin model.SysAdmin) (string, error) {
	// 计算过期时间
	expireTime := time.Now().Add(TokenExpireDuration)

	// 构建用户信息
	jwtAdmin := model.JwtAdmin{
		ID:         admin.ID,
		Username:   admin.Username,
		Nickname:   admin.Nickname,
		Status:     admin.Status,
		Icon:       admin.Icon,
		Email:      admin.Email,
		Phone:      admin.Phone,
		Note:       admin.Note,
		ExpireTime: expireTime.Unix(), // 存储过期时间戳
	}

	// 构建claims
	claims := userStdClaims{
		JwtAdmin: jwtAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 使用v4版本的时间格式
			Issuer:    Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 新增签发时间
		},
	}

	// 创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名生成token字符串
	return token.SignedString(Secret)
}
