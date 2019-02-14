package quick

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/testdata"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/utils"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, QuickSort(v.Nosort), "no eq")
	}
}


func benchmarkQuickSort(n int, b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(list)
	}
}

func BenchmarkQuickSort100(b *testing.B)    { benchmarkQuickSort(100, b) }
func BenchmarkQuickSort1000(b *testing.B)   { benchmarkQuickSort(1000, b) }
func BenchmarkQuickSort10000(b *testing.B)  { benchmarkQuickSort(10000, b) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkQuickSort(100000, b) }
