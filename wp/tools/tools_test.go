package tools

import (
	"testing"
)

func TestRandNum(t *testing.T) {
	for i := 0; i < 2000; i++ {
		if Randnum(100) > 100 || Randnum(100) < 1 {
			t.Error("范围超过")
		}
	}
}
