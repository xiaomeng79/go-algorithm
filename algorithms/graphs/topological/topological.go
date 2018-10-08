package topological

import (
	"github.com/xiaomeng79/go-algorithm/data-structures/graph"
	"github.com/xiaomeng79/go-algorithm/data-structures/stack"
)

//拓扑排序指的是将有向无环图（又称“DAG”图）中的顶点按照图中指定的先后顺序进行排序。

//拓扑排序实质上就是一种偏序到全序的排序算法

//应用场景:排课问题，课程有依赖关系

//对有向无环图进行拓扑排序，只需要遵循两个原则：
//在图中选择一个没有前驱的顶点 V；
//从图中删除顶点 V 和所有以该顶点为尾的弧。

func Sort(g *graph.DirGraph) *stack.Stack {
	var visit func(v graph.VertexId)
	sorted := stack.New()
	visited := make(map[graph.VertexId]bool)
	marked := make(map[graph.VertexId]bool)

	visit = func(v graph.VertexId) {
		//判断是不是一个无环图
		if marked[v] {
			panic("not a DAG")
		}
		marked[v] = true
		//获取出度的顶点
		successors := g.GetSuccessors(v).VerticesIter()
		for successor := range successors {
			if !marked[successor] && !visited[successor] {
				visit(successor)
			}
		}
		visited[v] = true
		marked[v] = false
		sorted.Push(int(v))
	}

	vertices := g.VerticesIter()
	for vertex := range vertices {
		if !visited[vertex] {
			visit(vertex)
		}
	}

	return sorted
}
