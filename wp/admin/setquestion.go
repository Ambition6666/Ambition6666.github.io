package admin

import (
	data "wp/database"
	"wp/question"

	"github.com/gin-gonic/gin"
)

func SetQuestion(c *gin.Context) {
	db := data.GetDB()
	var a question.Que
	c.Bind(&a)
	question := question.Question{
		Templet:  a.Templet,
		Answer:   a.Answer,
		Language: a.Language,
		Data:     a.Data,
	}
	db.Create(&question)
}
