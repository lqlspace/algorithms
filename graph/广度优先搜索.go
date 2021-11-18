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
				visited[k] = true
				prev[k] = top
				queue = append(queue, k)
				if k == t {
					isFound = true
					break
				}
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
	// 递归终止条件
	if t == s {
		fmt.Printf("%d ", t)
		return
	}

	// 递归打印之前的节点
	printPrev(prev, s, prev[t])
	fmt.Printf("%d ", t)
}




