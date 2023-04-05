package user

import "github.com/gin-gonic/gin"

func SearchErrorRecord(c *gin.Context) {
	db := GetDB()
	y, err := c.Get("user") //得到登录用户的信息
	if !err {
		c.JSON(200, "提取失败")
		return
	}
	user1 := y.(User)
	a := make([]UserRecord, 0)
	db.Model(UserRecord{}).Where("name=? and Error=?", user1.Name, true).Find(&a)
	for i := len(a) - 1; i >= 0; i-- {
		c.JSON(200, gin.H{
			"1.QID":         a[i].QID,
			"2.ErrorRecord": a[i].ErrorRecord,
			"3.time":        a[i].CreateAt,
		})
	}
}
