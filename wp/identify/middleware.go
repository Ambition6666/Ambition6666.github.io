package identify

import (
	"errors"
	"strings"
	data "wp/database"
	"wp/user"

	"github.com/gin-gonic/gin"
)

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
			ctx.JSON(401, "解析失败")
			ctx.Abort() //中间件不通过
			return
		}
		Uid := c.Model.ID
		db := data.GetDB()
		var user1 user.User
		db.Where("ID=?", Uid).First(&user1)
		if user1.ID == 0 {
			ctx.JSON(401, "没有该用户")
			ctx.Abort()
			return
		}
		ctx.Set("user", user1)
		ctx.Next()
	}

}
func GetUser(c *gin.Context) (user.User, error) {
	y, err := c.Get("user") //得到登录用户的信息
	if !err {

		e := errors.New("提取用户失败")
		return user.User{}, e
	}
	user1 := y.(user.User) //将any格式转化为USER格式
	return user1, nil
}
