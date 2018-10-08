package dijkstra

import (
	"github.com/xiaomeng79/go-algorithm/data-structures/graph"
	"github.com/xiaomeng79/go-algorithm/data-structures/priority_queue"
)

//迪杰斯特拉算法计算的是从网中一个顶点到其它顶点之间的最短路径问题
//http://data.biancheng.net/view/46.html

func ShortestPath(g *graph.UnGraph, source graph.VertexId) map[graph.VertexId]graph.VertexId {
	visited := make(map[graph.VertexId]bool, g.VerticesCount())
	dist := make(map[graph.VertexId]int)
	prev := make(map[graph.VertexId]graph.VertexId)
	Q := priority_queue.NewMin()
	vertices := g.VerticesIter()

	dist[source] = 0
	for vertex := range vertices {
		if source != vertex {
			dist[vertex] = 1000
			prev[vertex] = 0
		}
		Q.Insert(*priority_queue.NewItem(vertex, dist[vertex]))
	}

	for Q.Len() > 0 {
		u := Q.Extract().Value.(graph.VertexId)
		visited[u] = true

		for neighbour := range g.GetNeighbours(u).VerticesIter() {
			if !visited[neighbour] {
				alt := dist[u] + g.GetEdgeWeight(u, neighbour)

				if alt < dist[neighbour] {
					dist[neighbour] = alt
					prev[neighbour] = u
					Q.ChangePriority(neighbour, alt)
				}
			}
		}
	}
	return prev
}
