package routes

import (
	"wp/admin"
	"wp/identify"
	"wp/question"
	usercenter "wp/user/userCenter"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/login", identify.Login)
	answer := r.Group("/answer")
	answer.Use(identify.Mid())
	{
		answer.GET("/", question.Showdata)
		answer.POST("/daily/a", question.AnswerDailyQuestions)
		answer.POST("/daily/q", question.PublishDailyQuestions)
		answer.POST("/charpter/a", question.AnswerCharpterQuestion)
		answer.POST("/charpter/q", question.PublishCharpterQuestion)
		answer.GET("/daily", question.GetDailyErrorRecord)
		answer.GET("/charpter", question.GetCharpterErrorRecord)
	}
	correct := r.Group("/correct")
	correct.Use(identify.Mid())
	{
		correct.POST("/add", usercenter.AddCorrectQuestion)
		correct.POST("/get", usercenter.GetCorrectQuestion)
		correct.POST("/answer/q", usercenter.PublishCorrectQuestion)
		correct.POST("/answer/a", usercenter.AnswerCorrectQuestion)
		correct.POST("/delete", usercenter.DeleteCorrectForm)
	}
	r.POST("/look", identify.Mid(), usercenter.LookQuestionbank)
	r.POST("/", admin.SetQuestion)
}
