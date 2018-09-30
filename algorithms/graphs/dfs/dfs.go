package dfs

import (
	"github.com/xiaomeng79/go-algorithm/data-structures/graph"
	"github.com/xiaomeng79/go-algorithm/data-structures/stack"
)

//无向图的深度搜索

func UndirectedDfs(g *graph.UnGraph, v graph.VertexId, fn func(graph.VertexId)) {
	s := stack.New() //初始化一个栈
	s.Push(int(v))
	visited := make(map[graph.VertexId]bool) //记录是否访问过

	for s.Len() > 0 { //只要栈里面有，就一直
		v, _ := s.Pop()
		_v := graph.VertexId(v)
		if _, ok := visited[_v]; !ok {
			visited[_v] = true
			fn(_v)
			//将此结点的邻接点压栈
			neighbours := g.GetNeighbours(_v).VerticesIter()
			for neighbour := range neighbours {
				s.Push(int(neighbour))
			}
		}
	}
}

//有向图的深度搜索
func DirectedDfs(g *graph.DirGraph, v graph.VertexId, fn func(graph.VertexId)) {
	s := stack.New()
	s.Push(int(v))
	visited := make(map[graph.VertexId]bool)

	for s.Len() > 0 {
		v, _ := s.Pop()
		_v := graph.VertexId(v)
		if _, ok := visited[_v]; !ok {
			visited[_v] = true
			fn(_v)
			neighbours := g.GetSuccessors(_v).VerticesIter()
			for neighbour := range neighbours {
				s.Push(int(neighbour))
			}
		}
	}
}
