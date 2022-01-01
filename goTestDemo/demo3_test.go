package demo

import "testing"

func TestGetArea3(t *testing.T) {
	area := GetArea(40, 50)

	if area != 2000 {
		t.Error("测试失败！")
	}
}

func BenchmarkGetArea3(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}
