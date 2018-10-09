package insertion

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/testdata"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/utils"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, InsertionSort(v.Nosort), "no eq")
	}
}

func TestInsertionSort3(t *testing.T) {
	for _, v := range testdata.Values {
		InsertionSort3(v.Nosort)
		assert.Exactly(t, v.Sort, v.Nosort, "no eq")
	}
}

func benchmarkInsertionSort(n int, b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(list)
	}
}

func BenchmarkInsertionSort100(b *testing.B)    { benchmarkInsertionSort(100, b) }
func BenchmarkInsertionSort1000(b *testing.B)   { benchmarkInsertionSort(1000, b) }
func BenchmarkInsertionSort10000(b *testing.B)  { benchmarkInsertionSort(10000, b) }
func BenchmarkInsertionSort100000(b *testing.B) { benchmarkInsertionSort(100000, b) }

func benchmarkInsertionSort3(n int, b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort3(list)
	}
}

func BenchmarkInsertionSort3_100(b *testing.B)    { benchmarkInsertionSort3(100, b) }
func BenchmarkInsertionSort3_1000(b *testing.B)   { benchmarkInsertionSort3(1000, b) }
func BenchmarkInsertionSort3_10000(b *testing.B)  { benchmarkInsertionSort3(10000, b) }
func BenchmarkInsertionSort3_100000(b *testing.B) { benchmarkInsertionSort3(100000, b) }
