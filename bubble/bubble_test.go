package bubble

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	values = []struct {
		nosort []int
		sort   []int
	}{
		{[]int{3, 44, 56, 38, 77, 38, 26}, []int{3, 26, 38, 38, 44, 56, 77}},
		{[]int{3}, []int{3}},
		{[]int{3, -1, -6, 34, -78}, []int{-78, -6, -1, 3, 34}},
	}
	value = []int{3, 44, 56, 38, 77, 38, 26}
)

func TestBubbleSort(t *testing.T) {
	for _, v := range values {
		assert.Exactly(t, v.sort, BubbleSort(v.nosort), "no eq")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		BubbleSort(value)
	}
}
