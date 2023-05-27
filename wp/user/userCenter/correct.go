package usercenter

import (
	data "wp/database"
	"wp/identify"
	"wp/question"
	"wp/user"

	"github.com/gin-gonic/gin"
)

type m struct {
	Q int
	L string
	T string
}

func AddCorrectQuestion(c *gin.Context) {
	user1, _ := identify.GetUser(c)
	db := data.GetDB()
	var data m
	c.Bind(&data)
	cf := user.CorrteForm{
		Uid:      user1.ID,
		Qid:      data.Q,
		Templet:  data.T,
		Language: data.L,
	}
	db.Create(&cf)
	c.JSON(200, "创建成功")
}
func GetCorrectQuestion(c *gin.Context) {
	user1, _ := identify.GetUser(c)
	db := data.GetDB()
	cq := make([]user.CorrteForm, 0)
	db.Where("uid=?", user1.ID).Last(&cq)
	data1 := make([]map[string]any, 0)
	for i := 0; i < len(cq); i++ {
		form := gin.H{
			"u": cq[i].Uid,
			"l": cq[i].Language,
			"t": cq[i].Templet,
			"q": cq[i].Qid,
		}
		data1 = append(data1, form)
	}
	c.JSON(200, data1)
}
func PublishCorrectQuestion(c *gin.Context) {
	db := data.GetDB()
	var data m
	c.Bind(&data)
	var q []question.Question
	db.Where("Qid=?", data.Q).Find(&q)
	a := question.TurnH(q)
	c.JSON(200, a)
}
func AnswerCorrectQuestion(c *gin.Context) {
	question.AnswerCharpterQuestion(c)
}
func DeleteCorrectForm(c *gin.Context) {
	user1, _ := identify.GetUser(c)
	db := data.GetDB()
	var data m
	c.Bind(&data)
	d := user.CorrteForm{
		Uid:      user1.ID,
		Qid:      data.Q,
		Language: data.L,
		Templet:  data.T,
	}
	db.Delete(&d)
	c.JSON(200, "删除成功")
}
