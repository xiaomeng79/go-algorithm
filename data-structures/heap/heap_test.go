package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMax(t *testing.T) {
	mh := NewMax()
	mh.Insert(Int(5))
	mh.Insert(Int(7))
	mh.Insert(Int(2))
	mh.Insert(Int(6))
	mh.Insert(Int(3))
	assert.Equal(t, Int(7), mh.Extract())
	assert.Equal(t, Int(6), mh.Extract())
	mh.Insert(Int(12))
	assert.Equal(t, 4, mh.Len())
	assert.Equal(t, Int(12), mh.Extract())
}

func TestNewMin(t *testing.T) {
	mh := NewMin()
	mh.Insert(Int(5))
	mh.Insert(Int(7))
	mh.Insert(Int(2))
	mh.Insert(Int(6))
	mh.Insert(Int(3))
	assert.Equal(t, Int(2), mh.Extract())
	assert.Equal(t, Int(3), mh.Extract())
	mh.Insert(Int(12))
	assert.Equal(t, 4, mh.Len())
	assert.Equal(t, Int(5), mh.Extract())
}
