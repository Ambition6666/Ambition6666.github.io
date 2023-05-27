package tools

import (
	"math/rand"
	"time"
)

func Randnum(a int) int { //a是可抽题目的数量
	rand.New(rand.NewSource(time.Now().Unix()))
	i := rand.Intn(a) + 1
	return i
} //产生随机数
func DontRepeat(a []int, num int) bool {
	for i := 0; i < len(a); i++ {
		if num == a[i] {
			return true
		}
	}
	return false
}
