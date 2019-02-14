package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testData = []struct{
		s string
		_s string
		m,n int
	}{
		{"abcdef","defabc",3,6},
		{"12345678","45612378",3,6},
		{"gqdvf","qdvfg",1,5},
	}
)
func TestLeftRotateStritog(t *testing.T) {
	for _,ts := range testData {
		assert.Equal(t,ts._s,LeftRotateStritog(ts.s,ts.m,ts.n))
	}
}
