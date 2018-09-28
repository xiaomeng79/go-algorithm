package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	for i := 0; i < 10; i++ {
		err := s.Push(i)
		assert.NoError(t, err)
	}
	err := s.Push(11)
	assert.Error(t, err)
	for i := 9; i >= 0; i-- {
		k, _ := s.Pop()
		assert.Equal(t, i, k)
	}
	_, err = s.Pop()
	assert.Error(t, err)
}

func TestNewListStack(t *testing.T) {
	s := NewListStack()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	for i := 9; i >= 0; i-- {
		k, _ := s.Pop()
		assert.Equal(t, i, k)
	}
	_, err := s.Pop()
	assert.Error(t, err)
}
