package user

import (
	"aq/randnum"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type midquestion struct {
	qid      int
	td       string
	language string
}

var middleID midquestion

func SetQuestion(c *gin.Context) {
	db := GetDB()
	var question Question
	err := c.Bind(&question)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		c.JSON(200, "创建成功")
	}
	db.Create(question)
}
func PublishQuestion(c *gin.Context) {
	y, err := c.Get("user") //得到登录用户的信息
	if !err {
		c.JSON(200, "提取用户失败")
		return
	}
	user1 := y.(User) //将any格式转化为USER格式

	//分界线-------------------
	db := GetDB()
	language := c.Query("language")
	td := c.Query("td")
	a := make([]Question, 0)
	db.Model(Question{}).Where("Language=? and TD=?", language, td).Find(&a)
	//得到问题切片
	var num int64
	db.Model(UserRecord{}).Where("Language=? and TD=? and Error=? and name=?", language, td, false, user1.Name).Count(&num)
	if num >= int64(len(a)) {
		c.JSON(200, "已全部答完")
		return
	}
	i := 0
	for {
		i = randnum.Randnum(len(a)) - 1
		var c UserRecord
		db.Where("name=? and Q_ID=? and Error=?", user1.Name, a[i].QID, false).Find(&c)
		if c.QID != 0 {
			continue
		} else {
			break
		}
	}
	str := strings.Split(a[i].Date, "|")
	c.JSON(200, gin.H{
		"QID":      a[i].QID,
		"Question": str[0],
		"date":     str[1:],
	})
	middleID = midquestion{
		a[i].QID,
		a[i].TD,
		a[i].Language,
	}
}
func Answer(c *gin.Context) {
	db := GetDB()
	var a Question
	y, err := c.Get("user") //得到登录用户的信息
	if !err {
		c.JSON(200, "提取失败")
		return
	}

	user1 := y.(User)
	db.Where("Q_ID=? and TD=? and Language=?", middleID.qid, middleID.td, middleID.language).Find(&a)
	b := c.PostForm("answer")
	if a.Answer == b {
		c.JSON(200, "你答对了")
		d := UserRecord{
			Name:        user1.Name,
			CreateAt:    time.Now(),
			QID:         a.QID,
			Language:    a.Language,
			TD:          a.TD,
			Error:       false,
			ErrorRecord: "",
		}
		db.Create(&d)
	} else {
		c.JSON(200, "你答错了")
		d := UserRecord{
			Name:        user1.Name,
			CreateAt:    time.Now(),
			QID:         a.QID,
			Language:    a.Language,
			TD:          a.TD,
			Error:       true,
			ErrorRecord: b,
		}
		db.Create(&d)
	}
}
