package heap

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/testdata"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/utils"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, HeapSort(v.Nosort), "no eq")
	}
}

func benchmarkHeapSort(n int, b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		HeapSort(list)
	}
}

func BenchmarkHeapSort100(b *testing.B)   { benchmarkHeapSort(100, b) }
func BenchmarkHeapSort1000(b *testing.B)  { benchmarkHeapSort(1000, b) }
func BenchmarkHeapSort10000(b *testing.B) { benchmarkHeapSort(10000, b) }

//func BenchmarkHeapSort100000(b *testing.B) { benchmarkHeapSort(100000, b) }
