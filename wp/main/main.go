package main

import (
	data "wp/database"
	"wp/question"
	"wp/routes"
	"wp/user"

	"github.com/gin-gonic/gin"
)

func main() {
	data.InitDB()
	r := gin.Default()
	CreateTable()
	routes.Routes(r)
	r.Run(":9090")
}
func CreateTable() {
	db := data.GetDB()
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&user.ChaprterAnswerRecord{})
	db.AutoMigrate(&user.DailyAnswerRecord{})
	db.AutoMigrate(&question.Question{})
	db.AutoMigrate(&user.CorrteForm{})
}
