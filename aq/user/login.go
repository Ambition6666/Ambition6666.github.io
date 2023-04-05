package user

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claim struct {
	Username string
	jwt.RegisteredClaims
} //创建用户登录标签
var Msk []byte = []byte("ztyyyds666")

func GetToken(b []byte, c string) (string, error) { //得到token,c为用户名
	a := Claim{
		c,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
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
func Register(ctx *gin.Context) {
	db := GetDB()                //打开数据库
	name := ctx.PostForm("name") //接收前端发来用户名
	pwd := ctx.PostForm("pwd")   //接收前端发来密码
	var a User                   //判断该用户名是否已存在
	db.Where("name=?", name).First(&a)
	if a.Name == name {
		ctx.JSON(400, "该用户名已存在")
		return
	}
	user := User{
		Pwd:  pwd,
		Name: name,
	} //创建用户
	db.Create(&user) //在数据库中创建用户
	ctx.JSON(200, gin.H{
		"message": "创建成功"})
}
func Login(ctx *gin.Context) {
	db := GetDB()
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd") //接收用户名和密码
	var user3 User
	db.Where("name=?", name).First(&user3)
	if user3.ID == 0 {
		ctx.JSON(401, "没有该用户")
		return
	}
	if user3.Pwd != pwd {
		ctx.JSON(401, "密码错误")
		return
	}
	str, err := GetToken(Msk, user3.Name)
	if err != nil {
		ctx.JSON(500, "加密失败")
		return
	}

	ctx.JSON(200, gin.H{
		"date":    str,
		"message": "登录成功"})
}

func Mid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.GetHeader("Authorization") //得到字串开头
		if t == "" || !strings.HasPrefix(t, "Bearer ") {
			ctx.JSON(401, "解析失败")
			ctx.Abort()
			return
		}

		t = t[7:]                 //扔掉头部
		tk, c, e := ParseToken(t) //c为claim结构体的实例
		if e != nil || !tk.Valid {
			ctx.JSON(401, "解析失败1")
			ctx.Abort() //中间件不通过
			return
		}
		name := c.Username
		db := GetDB()
		var user4 User
		db.Where("name=?", name).First(&user4)
		if user4.ID == 0 {
			ctx.JSON(401, "没有该用户")
			ctx.Abort()
			return
		}
		ctx.Set("user", user4)
		ctx.Next()
	}

}
