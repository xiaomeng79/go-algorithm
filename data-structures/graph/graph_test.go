package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUndirected(t *testing.T) {
	g := NewUndirected()
	//增加顶点
	for i := 0; i < 10; i++ {
		g.AddVertex(VertexId(i))
	}
	assert.Equal(t, 10, g.VerticesCount())
	//增加边
	for i := 0; i < 10; i++ {
		g.AddEdge(VertexId(i), VertexId(i%2), 1)
	}
	//
	assert.Equal(t, true, g.CheckEdge(2, 0))
	assert.Equal(t, false, g.CheckEdge(2, 1))
	assert.Equal(t, 1, g.GetEdgeWeight(2, 0))
	//删除顶点
	assert.NoError(t, g.RemoveVertex(VertexId(2)))
	assert.Equal(t, false, g.CheckVertex(VertexId(2)))
	assert.Equal(t, false, g.CheckEdge(2, 0))
	//增加边，存在修改
	assert.NoError(t, g.AddEdge(3, 0, 1))
	assert.Equal(t, true, g.CheckEdge(3, 0))
	//删除边
	assert.NoError(t, g.RemoveEdge(3, 0))
	assert.Equal(t, false, g.CheckEdge(3, 0))

	//统计边
	c := g.EdgesIter()
	countEdge := 0
	for _ = range c {
		countEdge++
	}
	assert.Equal(t, countEdge, g.EdgesCount())

}

func TestNewDirected(t *testing.T) {
	g := NewDirected()
	//增加顶点
	for i := 0; i < 10; i++ {
		g.AddVertex(VertexId(i))
	}
	assert.Equal(t, 10, g.VerticesCount())
	//增加边
	for i := 0; i < 10; i++ {
		g.AddEdge(VertexId(i), VertexId(i%2), 1)
	}
	//
	assert.Equal(t, true, g.CheckEdge(2, 0))
	assert.Equal(t, false, g.CheckEdge(2, 1))
	assert.Equal(t, 1, g.GetEdgeWeight(2, 0))
	//删除顶点
	assert.NoError(t, g.RemoveVertex(VertexId(2)))
	assert.Equal(t, false, g.CheckVertex(VertexId(2)))
	assert.Equal(t, false, g.CheckEdge(2, 0))
	//增加边，存在修改
	assert.NoError(t, g.AddEdge(3, 0, 1))
	assert.Equal(t, true, g.CheckEdge(3, 0))
	//删除边
	assert.NoError(t, g.RemoveEdge(3, 0))
	assert.Equal(t, false, g.CheckEdge(3, 0))

	//统计边
	c := g.EdgesIter()
	countEdge := 0
	for _ = range c {
		countEdge++
	}
	assert.Equal(t, countEdge, g.EdgesCount())
	//查看
	//for p := range g.EdgesIter() {
	//	t.Log(p)
	//}
	//入度
	gp := g.GetPredecessors(VertexId(1)).VerticesIter()
	for p := range gp {
		if p != 3 && p != 5 && p != 7 && p != 9 {
			t.Error()
		}
	}

	for p := range g.GetSuccessors(VertexId(4)).VerticesIter() {
		assert.Equal(t, VertexId(0), p)
	}
}
