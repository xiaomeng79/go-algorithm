package insertion

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/xiaomeng79/go-algorithm/testdata"
)


func TestInsertionSort(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, InsertionSort(v.Nosort), "no eq")
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		InsertionSort(testdata.Value)
	}
}
