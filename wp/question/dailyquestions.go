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

func PublishDailyQuestions(c *gin.Context) {
	user1, err := identify.GetUser(c)
	if err != nil {
		c.JSON(333, err)
		return
	}
	//定义数组接收
	var qt MidDailyData
	c.Bind(&qt)
	//打开数据库，查询符合要求的题目，并将它存到一个数组里
	db := data.GetDB()
	questions := make([]Question, 0)
	fmt.Println(qt.T)
	db.Where("Language=? and Templet IN ?", qt.L, qt.T).Find(&questions)
	//去除答对的题目
	questions = IfTrueDelete(questions, user1.ID)
	//如果题目数量小于要抽取的数量,则全部发送
	if len(questions) < 10 {
		data := TurnH(questions)
		c.JSON(200, data)
		c.JSON(200, "已做完全部题目")
		return
	}
	//开始随机抽题,定义一个数组，避免重复抽数据库中相同的题
	index := make([]int, 0)
	k := 0
	for k < 10 {
		j := tools.Randnum(len(questions) - 1)
		if tools.DontRepeat(index, j) {
			continue
		}
		index = append(index, j)
		k++
	}
	fmt.Println(index)
	//将题目塞到一个数组发给前端
	m := make([]Question, 0)
	for i := 0; i < len(index); i++ {
		m = append(m, questions[index[i]])
	}
	data := TurnH(m)
	c.JSON(200, data)
}
func AnswerDailyQuestions(c *gin.Context) {
	//获取用户信息
	u, e := identify.GetUser(c)
	if e != nil {
		c.JSON(333, e)
		return
	}
	//接收前端返回的值
	var qlta []MidAnData
	c.Bind(&qlta)
	//打开数据库
	db := data.GetDB()
	//定义一个存题号的数组
	index := make([]int, 0)
	for i := 0; i < len(qlta); i++ {
		index = append(index, qlta[i].Q)
	}
	//取出需要的答案
	questions := make([]Question, 0)
	for i := 0; i < len(index); i++ {
		var question Question
		db.Where("qid=?", index[i]).Find(&question)
		questions = append(questions, question)
	}
	//创建答题记录
	for i := 0; i < len(questions); i++ {
		str1 := strings.Split(questions[i].Answer, "|||")
		if len(qlta[i].A) != len(str1)-1 {
			str := "你的错误回答："
			for k := 0; k < len(qlta[i].A); k++ {
				str += (qlta[i].A[k] + ",")
			}
			str += ("解析:" + str1[len(str1)-1])
			c.JSON(200, str)
			record := user.DailyAnswerRecord{
				CreateAt:    time.Now(),
				Uid:         u.ID,
				Qid:         qlta[i].Q,
				Language:    qlta[i].L,
				Templet:     qlta[i].T,
				Error:       false,
				ErrorRecord: str,
			}
			db.Create(&record)
			continue
		}
		for j := 0; j < len(str1)-1; j++ {
			if qlta[i].A[j] == str1[j] && j != len(str1)-2 {
				continue
			} else if qlta[i].A[j] == str1[j] && j == len(str1)-2 {
				record := user.DailyAnswerRecord{
					CreateAt:    time.Now(),
					Uid:         u.ID,
					Qid:         qlta[i].Q,
					Language:    qlta[i].L,
					Templet:     qlta[i].T,
					Error:       true,
					ErrorRecord: "",
				}
				c.JSON(200, "你答对了")
				db.Create(&record)
			} else {
				str := "你的错误回答："
				for k := 0; k < len(qlta[i].A); k++ {
					str += (qlta[i].A[k] + ",")
				}
				str += ("解析:" + str1[len(str1)-1])
				c.JSON(200, str)
				record := user.DailyAnswerRecord{
					CreateAt:    time.Now(),
					Uid:         u.ID,
					Qid:         qlta[i].Q,
					Language:    qlta[i].L,
					Templet:     qlta[i].T,
					Error:       false,
					ErrorRecord: str,
				}
				db.Create(&record)
				break
			}
		}
	}
}
func IfTrueDelete(a []Question, b uint) []Question {
	db := data.GetDB()
	for i := 0; i < len(a); i++ {
		var tf bool
		db.Model(user.DailyAnswerRecord{}).Select("error").Where("qid=? and uid=?", a[i].Qid, b).Last(&tf)
		if i == len(a)-1 {
			if tf {
				a = a[:len(a)-1]
				break
			}
		}
		if tf {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}
