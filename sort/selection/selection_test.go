package selection

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/sort/testdata"
	"github.com/xiaomeng79/go-algorithm/sort/utils"
)

func TestSelectionSort(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, SelectionSort(v.Nosort), "no eq")
	}
}


func benchmarkSelectionSort(n int,b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i:=0;i < b.N;i++ {
		SelectionSort(list)
	}
}


func BenchmarkSelectionSort100(b *testing.B) { benchmarkSelectionSort(100, b) }
func BenchmarkSelectionSort1000(b *testing.B)   { benchmarkSelectionSort(1000, b) }
func BenchmarkSelectionSort10000(b *testing.B)  { benchmarkSelectionSort(10000, b) }
func BenchmarkSelectionSort100000(b *testing.B) { benchmarkSelectionSort(100000, b) }

