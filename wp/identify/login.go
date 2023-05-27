package identify

import (
	"time"
	data "wp/database"
	"wp/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var Msk []byte = []byte("ztyyyds666")

func GetToken(b []byte, c uint) (string, error) { //得到token,c为用户名
	a := Claim{
		gorm.Model{
			ID: c,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Issuer:    "zty",
		},
	} //获取claim实例
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a) //获取token
	return token.SignedString(b)                          //返回加密串
}
func ParseToken(a string) (*jwt.Token, *Claim, error) { //解析token
	claim := &Claim{}
	t, err := jwt.ParseWithClaims(a, claim, func(t *jwt.Token) (interface{}, error) {
		return Msk, nil
	}) //接收前端发来加密字段
	return t, claim, err
}

func Login(c *gin.Context) {
	user1 := new(user.User)
	jscode := c.Query("jscode")
	u, e := user1.WxLogin(jscode)
	if e != nil {
		c.JSON(401, e)
		return
	}
	if u.OpenID == "" {
		c.JSON(401, "登录失败")
		return
	}
	db := data.GetDB()
	var user2 user.User
	db.Where("openid=?", u.OpenID).Find(&user2)
	if user2.ID == 0 {
		a := user.User{
			Openid: u.OpenID,
			Role:   0,
		}
		db.Create(&a)
		db.Model(user.User{}).Where("openid=?", u.OpenID).Find(&user2.ID)
	}
	str, err := GetToken(Msk, user2.ID)
	if err != nil {
		c.JSON(401, "加密失败")
		return
	}

	c.JSON(200, gin.H{
		"token":   str,
		"message": "登录成功"})
}
