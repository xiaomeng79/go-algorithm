package bfs_shortest_path

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/data-structures/graph"
	"testing"
)

func TestShortestPath(t *testing.T) {
	h := graph.NewDirected()
	for i := 0; i < 10; i++ {
		h.AddVertex(graph.VertexId(i))
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}

	assert.Equal(t, 9, GetDist(h, graph.VertexId(0), graph.VertexId(9)))
}
