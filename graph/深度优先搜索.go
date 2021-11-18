package graph


//  search by DFS
func (g *Graph) DFS(s, t int) {
	prev := make([]int, g.v)
	for i :=  range prev {
		prev[i] = -1
	}

	visited := make([]bool, g.v)
	visited[s] = true

	g.recurse(prev, visited, s, t)

	printPrev(prev, s, t)
}

func (g *Graph) recurse(prev []int, visited []bool, s, t int) {

	visited[s] = true
	if s == t {
		return
	}

	linkedList := g.adj[s]
	for e := linkedList.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurse(prev, visited, k, t)
		}
	}
}
