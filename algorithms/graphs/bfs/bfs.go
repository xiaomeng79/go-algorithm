package bfs

import "github.com/xiaomeng79/go-algorithm/data-structures/graph"

//有向图的广度优先搜索
func Bfs(g *graph.DirGraph, start graph.VertexId, fn func(graph.VertexId)) {
	queue := []graph.VertexId{start}
	visited := make(map[graph.VertexId]bool)

	var next []graph.VertexId

	for len(queue) > 0 {
		next = []graph.VertexId{}
		for _, vertex := range queue {
			visited[vertex] = true
			//获取邻接点
			neighbours := g.GetNeighbours(vertex).VerticesIter()
			fn(vertex)
			for neighbour := range neighbours {
				if _, ok := visited[neighbour]; !ok {
					next = append(next, neighbour)
				}
			}
		}
		queue = next
	}
}
