package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"unique"`
	Pwd  string
}
type UserRecord struct {
	Name        string
	CreateAt    time.Time
	QID         int
	Language    string
	TD          string
	Error       bool
	ErrorRecord string
}
type Question struct {
	QID      int
	TD       string
	Language string
	Date     string
	Answer   string
}
