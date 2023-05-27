package question

import (
	"fmt"
	"strings"
	"time"
	data "wp/database"
	"wp/identify"
	"wp/tools"
	"wp/user"

	"github.com/gin-gonic/gin"
)

func PublishCharpterQuestion(c *gin.Context) {
	var lt MidData
	c.Bind(&lt)
	db := data.GetDB()
	quetions := make([]Question, 0)
	fmt.Println()
	if lt.T == "all" {
		db.Where("Language=?", lt.L).Find(&quetions)
		index := make([]int, 0)
		k := 0
		for k == 10 {
			j := tools.Randnum(len(quetions))
			if tools.DontRepeat(index, j) {
				continue
			}
			index = append(index, j)
			k++
		}
		m := make([]Question, 0)
		for i := 0; i < len(index); i++ {
			m = append(m, quetions[index[i]])
		}
		data := make([]map[string]any, 0)
		for i := 0; i < len(m); i++ {
			str1 := strings.Split(m[i].Data, "|||")
			str2 := strings.Split(m[i].Answer, "|||")
			d := gin.H{
				"id":       m[i].Qid,
				"language": m[i].Language,
				"templet":  m[i].Templet,
				"Question": str1[0],
				"Option":   str1[1:],
				"Answer":   str2,
			}
			data = append(data, d)
		}
		c.JSON(200, data)
		return
	}
	db.Where("Language=? and Templet=?", lt.L, lt.T).Find(&quetions)
	data := TurnH(quetions)
	c.JSON(200, data)
}
func AnswerCharpterQuestion(c *gin.Context) {
	var qlta MidAnData
	c.Bind(&qlta)
	db := data.GetDB()
	var quetion Question
	db.Where("Qid=? ", qlta.Q).Find(&quetion)
	u, e := identify.GetUser(c)
	if e != nil {
		c.JSON(333, e)
		return
	}
	str1 := strings.Split(quetion.Answer, "|||")
	for i := 0; i < len(str1)-1; i++ {
		if qlta.A[i] == str1[i] && i != len(str1)-2 {
			continue
		} else if qlta.A[i] == str1[i] && i == len(str1)-2 {
			record := user.ChaprterAnswerRecord{
				CreateAt:    time.Now(),
				Uid:         u.ID,
				Qid:         qlta.Q,
				Language:    qlta.L,
				Templet:     qlta.T,
				Error:       true,
				ErrorRecord: "",
			}
			c.JSON(200, "你答对了")
			db.Create(&record)
		} else {
			str := "你的错误回答："
			for j := 0; j < len(qlta.A); j++ {
				str += (qlta.A[i] + ",")
			}
			str += ("解析:" + str1[len(str1)-1])
			c.JSON(200, str)
			record := user.ChaprterAnswerRecord{
				CreateAt:    time.Now(),
				Uid:         u.ID,
				Qid:         qlta.Q,
				Language:    qlta.L,
				Templet:     qlta.T,
				Error:       false,
				ErrorRecord: str,
			}
			db.Create(&record)
			break
		}
	}
}
func TurnH(quetions []Question) []map[string]any {
	data := make([]map[string]any, 0)
	for i := 0; i < len(quetions); i++ {
		str1 := strings.Split(quetions[i].Data, "|||")
		str2 := strings.Split(quetions[i].Answer, "|||")
		d := gin.H{
			"id":       quetions[i].Qid,
			"language": quetions[i].Language,
			"templet":  quetions[i].Templet,
			"Question": str1[0],
			"Option":   str1[1:],
			"Answer":   str2,
		}
		data = append(data, d)
	}
	return data
}
