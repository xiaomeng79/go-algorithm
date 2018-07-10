package selection

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/testdata"
)

func TestSelectionSort(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, SelectionSort(v.Nosort), "no eq")
	}
}


func BenchmarkSelectionSort(b *testing.B) {
	b.ReportAllocs()
	for i:=0;i<b.N;i++ {
		SelectionSort(testdata.Value)
	}
}

