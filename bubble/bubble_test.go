package bubble

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/xiaomeng79/go-algorithm/testdata"
)


func TestBubbleSort(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, BubbleSort(v.Nosort), "no eq")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		BubbleSort(testdata.Value)
	}
}
