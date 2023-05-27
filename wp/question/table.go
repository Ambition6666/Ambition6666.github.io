package question

import (
	"strconv"
	"time"
	data "wp/database"
	"wp/identify"
	"wp/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Showdata(c *gin.Context) {
	db := data.GetDB()
	totalData := make([]map[string]any, 0)
	totalerrorData := make([]map[string]any, 0)
	user1, err := identify.GetUser(c)
	if err != nil {
		c.JSON(333, err)
		return
	}

	day := c.Query("day")
	totalDay, _ := strconv.Atoi(day)
	for i := 0; i < totalDay; i++ {
		a, b := num(user1, db, totalDay)
		totalerrorData = append(totalerrorData, b)
		totalData = append(totalData, a)
	}
	c.JSON(200, totalData)
	c.JSON(200, totalerrorData)
}
func num(user1 user.User, db *gorm.DB, i int) (map[string]any, map[string]any) {
	var (
		num1 int64
		num2 int64
		num3 int64
		num4 int64
	)
	db.Model(user.ChaprterAnswerRecord{}).Where("uid=? and create_at BETWEEN ? AND ?", user1.ID, time.Now().Add((-1)*time.Duration(i)*24*time.Hour), time.Now().Add((-1)*time.Duration(i-1)*24*time.Hour)).Count(&num1)
	db.Model(user.DailyAnswerRecord{}).Where("uid=? and create_at BETWEEN ? AND ?", user1.ID, time.Now().Add((-1)*time.Duration(i)*24*time.Hour), time.Now().Add((-1)*time.Duration(i-1)*24*time.Hour)).Count(&num2)
	db.Model(user.ChaprterAnswerRecord{}).Where("uid=? and error =? and create_at BETWEEN ? AND ?", user1.ID, true, time.Now().Add((-1)*time.Duration(i)*24*time.Hour), time.Now().Add((-1)*time.Duration(i-1)*24*time.Hour)).Count(&num3)
	db.Model(user.DailyAnswerRecord{}).Where("uid=? and error =? and create_at BETWEEN ? AND ?", user1.ID, true, time.Now().Add((-1)*time.Duration(i)*24*time.Hour), time.Now().Add((-1)*time.Duration(i-1)*24*time.Hour)).Count(&num4)
	return gin.H{
			"time": time.Now().Add((-1) * time.Duration(i*int(time.Hour)*24)),
			"全部":   num1 + num2,
		}, gin.H{
			"time": time.Now().Add((-1) * time.Duration(i*int(time.Hour)*24)),
			"正确":   num3 + num4,
		}
}
