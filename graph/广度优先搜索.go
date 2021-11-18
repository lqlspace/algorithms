package graph

import (
	"fmt"
)



// search path by BFS
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}

	// init prev
	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}

	// search by queue
	var queue []int
	visited := make([]bool, g.v)
	queue = append(queue, s)
	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedList := g.adj[top]
		for e := linkedList.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}

	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path  found from  %d to %d", s, t)
	}
}


// print path recursively
func printPrev(prev []int, s, t int) {
	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}
}




