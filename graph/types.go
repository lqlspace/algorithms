package graph

import (
	"container/list"
)

// adjacency table, 无向图
type Graph struct {
	adj []*list.List
	v  int
}

// init graph according to capacity
func newGraph(v int) *Graph {
	graph :=  &Graph{}
	graph.v  =  v
	graph.adj = make([]*list.List, v)
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}

	return  graph
}


// insert as add edge, 一条边存2次
func (g *Graph) addEdge(s, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}
