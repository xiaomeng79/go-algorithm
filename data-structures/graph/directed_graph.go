package graph

type DirGraph struct {
	graph
}

//有向图
func NewDirected() *DirGraph {
	return &DirGraph{
		graph{
			edgesCount: 0,
			edges:      make(map[VertexId]map[VertexId]int),
			isDirected: true,
		},
	}
}

//入度
func (g *graph) GetPredecessors(vertex VertexId) VerticesIterable {
	iterator := func() <-chan VertexId {
		ch := make(chan VertexId)
		go func() {
			for VertexId, _ := range g.edges {
				if g.CheckEdge(VertexId, vertex) {
					ch <- VertexId
				}
			}
			close(ch)
		}()
		return ch
	}

	return VerticesIterable(&VerticesIterableHelp{iter: iterator})
}

//出度
func (g *graph) GetSuccessors(vertex VertexId) VerticesIterable {
	iterator := func() <-chan VertexId {
		ch := make(chan VertexId)
		go func() {
			if connected, ok := g.edges[vertex]; ok {
				for VertexId, _ := range connected {
					if g.CheckEdge(vertex, VertexId) {
						ch <- VertexId
					}
				}
			}
			close(ch)
		}()
		return ch
	}

	return VerticesIterable(&VerticesIterableHelp{iter: iterator})
}

func (g *DirGraph) Reverse() *DirGraph {
	r := NewDirected()

	vertices := g.VerticesIter()
	for vertex := range vertices {
		r.AddVertex(vertex)
	}

	edges := g.EdgesIter()
	for edge := range edges {
		r.AddEdge(edge.To, edge.From, 1)
	}

	return r
}
