package bitmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBitMap(t *testing.T) {
	max := 100
	b := NewBitMap(max)
	b.Add(13)
	b.Add(100)
	b.Add(0)
	assert.Equal(t, true, b.IsExist(13))
	assert.Equal(t, true, b.IsExist(0))
	assert.Equal(t, true, b.IsExist(100))
	b.Remove(13)
	assert.Equal(t, false, b.IsExist(13))
	assert.Equal(t, max, b.Max())
}
