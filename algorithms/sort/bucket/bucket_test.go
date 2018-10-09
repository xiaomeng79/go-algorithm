package bucket

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/testdata"
	"github.com/xiaomeng79/go-algorithm/algorithms/sort/utils"
	"testing"
)

func TestBucketSort(t *testing.T) {
	i := 0
	for _, v := range testdata.Values {
		i++
		if i == 2 {
			break
		}
		assert.Exactly(t, v.Sort, BucketSort(v.Nosort), "no eq")
	}
}

func benchmarkBucketSort(n int, b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		BucketSort(list)
	}
}

func BenchmarkBucketSort100(b *testing.B)    { benchmarkBucketSort(100, b) }
func BenchmarkBucketSort1000(b *testing.B)   { benchmarkBucketSort(1000, b) }
func BenchmarkBucketSort10000(b *testing.B)  { benchmarkBucketSort(10000, b) }
func BenchmarkBucketSort100000(b *testing.B) { benchmarkBucketSort(100000, b) }
