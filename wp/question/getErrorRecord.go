package question

import (
	data "wp/database"
	"wp/identify"
	"wp/user"

	"github.com/gin-gonic/gin"
)

func GetDailyErrorRecord(c *gin.Context) {
	//获取用户信息
	u, e := identify.GetUser(c)
	if e != nil {
		c.JSON(333, e)
	}
	//打开数据库
	db := data.GetDB()
	//发送数据
	records := make([]user.DailyAnswerRecord, 0)
	db.Where("uid=? and error =?", u.ID, false).Find(&records)
	c.JSON(200, records)
}
func GetCharpterErrorRecord(c *gin.Context) {
	//获取用户信息
	u, e := identify.GetUser(c)
	if e != nil {
		c.JSON(333, e)
	}
	//打开数据库
	db := data.GetDB()
	//发送数据
	records := make([]user.ChaprterAnswerRecord, 0)
	db.Where("uid=? and error =?", u.ID, false).Find(&records)
	c.JSON(200, records)
}
