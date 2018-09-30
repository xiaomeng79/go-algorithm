package bfs_shortest_path

import (
	"github.com/xiaomeng79/go-algorithm/algorithms/graphs/bfs"
	"github.com/xiaomeng79/go-algorithm/data-structures/graph"
)

func ShortestPath(g *graph.DirGraph, start graph.VertexId) (dist map[graph.VertexId]int) {
	dist = make(map[graph.VertexId]int)
	visited := make(map[graph.VertexId]bool)

	getDist := func(v graph.VertexId) { //目标顶点
		neighbours := g.GetNeighbours(v).VerticesIter()
		visited[v] = true

		for neighbour := range neighbours {

			ok, _ := visited[neighbour]
			if !ok {
				dist[neighbour] = dist[v] + 1
			}
		}
	}
	bfs.Bfs(g, start, getDist)
	return
}

func GetDist(g *graph.DirGraph, from, to graph.VertexId) int {
	return ShortestPath(g, from)[to]
}
