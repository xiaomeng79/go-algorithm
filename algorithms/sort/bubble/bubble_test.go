package bubble

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/testdata"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, v := range testdata.Values {
		BubbleSort(v.Nosort)
		assert.Exactly(t, v.Sort, v.Nosort, "no eq")
	}
}


func benchmarkBubbleSort(n int, b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(list)
	}
}


func BenchmarkBubbleSort100(b *testing.B)    { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkBubbleSort(100000, b) }
