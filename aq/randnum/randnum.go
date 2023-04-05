package randnum

import (
	"math/rand"
	"time"
)

func Randnum(a int) int {
	rand.New(rand.NewSource(time.Now().Unix()))
	i := rand.Intn(a) + 1
	return i
}
