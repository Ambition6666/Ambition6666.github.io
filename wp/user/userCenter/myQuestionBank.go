package usercenter

import (
	data "wp/database"
	"wp/identify"
	"wp/user"

	"github.com/gin-gonic/gin"
)

func LookQuestionbank(c *gin.Context) {
	user1, _ := identify.GetUser(c)
	db := data.GetDB()
	qb1 := make([]user.ChaprterAnswerRecord, 0)
	qb2 := make([]user.DailyAnswerRecord, 0)
	db.Where("uid=?", user1.ID).Find(&qb1)
	db.Where("uid=?", user1.ID).Find(&qb2)
	c.JSON(200, gin.H{
		"每日": qb2,
	})
	c.JSON(200, gin.H{
		"章节": qb1,
	})
}
