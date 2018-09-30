package graph

import "errors"

const Infinity int = 65535 //无穷大

//非线程安全
type VertexId uint

type Vertices []VertexId

type Edge struct {
	From VertexId
	To   VertexId
}

type graph struct {
	edges      map[VertexId]map[VertexId]int
	edgesCount int
	isDirected bool //定向图
}

type EdgesIterable interface {
	EdgesIter() <-chan Edge
}

type VerticesIterable interface {
	VerticesIter() <-chan VertexId
}

//边
func (g *graph) EdgesIter() <-chan Edge {
	ch := make(chan Edge)
	go func() {
		for from, connectedVertices := range g.edges {
			for to, _ := range connectedVertices {
				if g.isDirected { //有向边
					ch <- Edge{from, to}
				} else {
					if from < to { //无向边，只输出一次
						ch <- Edge{from, to}
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

//顶点
func (g *graph) VerticesIter() <-chan VertexId {
	ch := make(chan VertexId)
	go func() {
		for vertex, _ := range g.edges {
			ch <- vertex
		}
		close(ch)
	}()
	return ch
}

//检查顶点是否存在
func (g *graph) CheckVertex(vertex VertexId) bool {
	_, exists := g.edges[vertex]
	return exists
}

//增加顶点
func (g *graph) AddVertex(vertex VertexId) error {
	if !g.CheckVertex(vertex) {
		g.edges[vertex] = make(map[VertexId]int)
		return nil
	} else {
		return errors.New("Vertex already exists")
	}
}

//删除顶点
func (g *graph) RemoveVertex(vertex VertexId) error {
	if !g.CheckVertex(vertex) {
		return errors.New("unknow vertex")
	}
	//先删除边
	for _to, _ := range g.edges[vertex] {
		g.RemoveEdge(vertex, _to)
	}
	delete(g.edges, vertex)
	for _, connectedVertices := range g.edges {
		delete(connectedVertices, vertex)
	}
	return nil
}

//统计顶点
func (g *graph) VerticesCount() int {
	return len(g.edges)
}

//判断边是否存在
func (g *graph) CheckEdge(from, to VertexId) bool {
	if _, ok := g.edges[from][to]; ok {
		return true
	}
	return false
}

//增加边,存在就修改权
func (g *graph) AddEdge(from, to VertexId, weight int) error {
	//自身循环
	if from == to {
		return errors.New("cannot add self loop")
	}
	//不存在边
	if !g.CheckVertex(from) || !g.CheckVertex(to) {
		return errors.New("vertices not exist")
	}
	//判断边存在不
	if g.edges[from][to] > 0 {
		return errors.New("edge  exist")
	}
	g.edges[from][to] = weight
	if !g.isDirected {
		g.edges[to][from] = weight
	}
	g.edgesCount++
	return nil
}

//删除边
func (g *graph) RemoveEdge(from, to VertexId) error {
	//判断边是否存在
	if !g.CheckEdge(from, to) {
		return errors.New("edge not exist")
	}
	//删除
	delete(g.edges[from], to)
	if !g.isDirected {
		delete(g.edges[to], from)
	}
	g.edgesCount--
	return nil
}

//统计边
func (g *graph) EdgesCount() int {
	return g.edgesCount
}

//获取边权
func (g *graph) GetEdgeWeight(from, to VertexId) int {
	if !g.CheckEdge(from, to) {
		return Infinity
	}
	return g.edges[from][to]
}

//获取邻结点和权
func (g *graph) GetNeighbours(vertex VertexId) VerticesIterable {
	iterator := func() <-chan VertexId {
		ch := make(chan VertexId)
		go func() {
			if connected, ok := g.edges[vertex]; ok {
				for vid, _ := range connected {
					ch <- vid
				}
			}
			close(ch)
		}()
		return ch
	}
	return VerticesIterable(&VerticesIterableHelp{iter: iterator})
}

//帮助获取
type VerticesIterableHelp struct {
	iter func() <-chan VertexId
}

func (v *VerticesIterableHelp) VerticesIter() <-chan VertexId {
	return v.iter()
}
