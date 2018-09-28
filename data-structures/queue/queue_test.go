package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	queue := New()
	assert.Equal(t, true, queue.isEmpty())
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	assert.Equal(t, 3, queue.Len())
	assert.Equal(t, 1, queue.Shift().(int))
	assert.Equal(t, 2, queue.Shift().(int))
	assert.Equal(t, 3, queue.Peek().(int))
	assert.Equal(t, 3, queue.Shift().(int))

}

func TestNewListQueue(t *testing.T) {
	queue := NewListQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	assert.Equal(t, 3, queue.Len())
	assert.Equal(t, 1, queue.Shift().(int))
	assert.Equal(t, 2, queue.Shift().(int))
	assert.Equal(t, 3, queue.Shift().(int))
	assert.Equal(t, 0, queue.Len())
}
