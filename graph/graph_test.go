package graph

import (
	"testing"
)

var graph *Graph
func init()  {
	graph = newGraph(8)
	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 4)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 7)

}

func TestGraph_BFS(t *testing.T) {
	graph.BFS(0, 7)
}

func TestGraph_DFS(t *testing.T) {
	graph.DFS(0, 7)
}
