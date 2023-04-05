package main

import (
	"aq/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", user.Login)
	r.POST("/register", user.Register)
	answer := r.Group("/answer")
	answer.Use(user.Mid())
	{
		answer.GET("/", user.PublishQuestion)
		answer.POST("/", user.Answer)
		answer.POST("/errorRecord", user.SearchErrorRecord)
	}
	r.POST("/", user.SetQuestion)
	r.Run(":9090")
}
