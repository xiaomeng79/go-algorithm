package merge

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/testdata"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/utils"
	"testing"
)

func TestMergeSort(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, MergeSort(v.Nosort), "no eq")
	}
}

func benchmarkMergeSort(n int, b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//sort(list)
		MergeSort(list)
	}
}

func BenchmarkMergeSort100(b *testing.B)    { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)   { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B)  { benchmarkMergeSort(10000, b) }
func BenchmarkMergeSort100000(b *testing.B) { benchmarkMergeSort(100000, b) }

func benchmarkMergeSort2(n int, b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sort(list)
	}
}

func BenchmarkMergeSort2_100(b *testing.B)    { benchmarkMergeSort2(100, b) }
func BenchmarkMergeSort2_1000(b *testing.B)   { benchmarkMergeSort2(1000, b) }
func BenchmarkMergeSort2_10000(b *testing.B)  { benchmarkMergeSort2(10000, b) }
func BenchmarkMergeSort2_100000(b *testing.B) { benchmarkMergeSort2(100000, b) }
