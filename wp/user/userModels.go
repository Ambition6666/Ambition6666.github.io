package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Openid string `gorm:"index"`
	Role   int
}
type WxSession struct {
	SessionKey string `json:"session_key"`
	ExpireIn   int    `json:"expires_in"`
	OpenID     string `json:"openid"`
}
type ChaprterAnswerRecord struct {
	CreateAt    time.Time `gorm:"primaryKey"`
	Uid         uint
	Qid         int
	Language    string
	Templet     string
	Error       bool `gorm:"default:false"`
	ErrorRecord string
}
type DailyAnswerRecord struct {
	CreateAt    time.Time `gorm:"primaryKey"`
	Uid         uint
	Qid         int
	Language    string `gorm:"index"`
	Templet     string `gorm:"index"`
	Error       bool   `gorm:"default:false"`
	ErrorRecord string
}
type CorrteForm struct {
	Uid      uint   `json:"uid"`
	Qid      int    `gorm:"primaryKey"`
	Templet  string `json:"templet"`
	Language string `json:"language"`
} //收藏所需信息
