package binaryTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Compare(t *testing.T) {
	n := NewNode(1)
	m := NewNode(2)
	k := NewNode(1)

	assert.Equal(t, -1, n.Compare(m))
	assert.Equal(t, 0, n.Compare(k))
	assert.Equal(t, 1, m.Compare(k))
}

func TestTree(t *testing.T) {
	max_size := 10
	tree := NewTree(nil)
	for i := 0; i < max_size; i++ {
		tree.Insert(ElementType(i))
	}
	assert.Equal(t, 10, tree.Size)
	assert.Equal(t, ElementType(5), tree.Search(5).Value)
	assert.Equal(t, true, tree.Delete(5))
	assert.Equal(t, 9, tree.Size)

}
