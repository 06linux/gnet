package gbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilAdd(t *testing.T) {

	value := Util.Add(1, 2)

	assert.Equal(t, value, 3)

	var a string = "Hello"
	var b string = "Hello"
	assert.Equal(t, a, b, "The two words should be the same.")

}

func BenchmarkUtilAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Util.Add(1, 2)
	}
}

// 基准测试就是在一定的工作负载之下检测程序性能的一种方法
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}
