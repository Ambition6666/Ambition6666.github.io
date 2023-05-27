package identify

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Claim struct {
	gorm.Model
	jwt.RegisteredClaims
} //创建用户登录标签
