package graph

import (
	"fmt"
)

type AdjacencyList struct {
	List          map[int][]int
	Visited       map[int]bool
	TraversalPath []int
}

func (al *AdjacencyList) DFS(v int) {
	al.Visited[v] = true
	al.TraversalPath = append(al.TraversalPath, v)

	for _, av := range al.List[v] {
		if _, ok := al.Visited[av]; !ok {
			al.DFS(av)
		}
	}
}

// Breadth First Search
// The "patient" strategy: search nearest nodes first
// Useful for: Explore all short-distance options before proceeding
// Find distance to starting point node
func (al *AdjacencyList) BFS(v int) {
	al.Visited[v] = true
	queue := al.List[v]
	al.TraversalPath = append(al.TraversalPath, v)

	for len(queue) > 0 {
		al.TraversalPath = append(al.TraversalPath, queue[0])
		al.Visited[queue[0]] = true

		for _, v := range al.List[queue[0]] {
			if _, ok := al.Visited[v]; ok {
				continue
			}
			if queue[len(queue)-1] == v {
				continue
			}

			queue = append(queue, v)
		}

		queue = queue[1:]
	}
}

// Strategy: Handle each node before its children
// For printing a hierarchy
func (al *AdjacencyList) Preorder(v, level int) {
	al.Visited[v] = true
	fmt.Printf("Vertex: %d level=%d \n", v, level)

	for _, av := range al.List[v] {
		if _, ok := al.Visited[av]; !ok {
			al.Preorder(av, level+1)
		}
	}
}

func init() {
	// Examples:
	adjacencyList := &AdjacencyList{
		List:    map[int][]int{},
		Visited: map[int]bool{},
	}

	// Graph: DFS
	adjacencyList.List[1] = []int{2, 3}
	adjacencyList.List[2] = []int{1, 4, 5}
	adjacencyList.List[3] = []int{1, 6}
	adjacencyList.List[4] = []int{2, 7}
	adjacencyList.List[5] = []int{2, 7}
	adjacencyList.List[6] = []int{3, 7, 8, 9}
	adjacencyList.List[7] = []int{4, 5, 6}
	adjacencyList.List[8] = []int{6}
	adjacencyList.List[9] = []int{6}

	adjacencyList.DFS(1)
	fmt.Printf("DFS traversal: %+v \n", adjacencyList.TraversalPath)

	// Graph: BFS
	adjacencyList.TraversalPath = []int{}
	adjacencyList.Visited = map[int]bool{}
	adjacencyList.BFS(1)
	fmt.Printf("BFS traversal: %+v \n", adjacencyList.TraversalPath)

	// Tree: preorder DFS
	adjacencyList = &AdjacencyList{
		List:    map[int][]int{},
		Visited: map[int]bool{},
	}
	adjacencyList.List[1] = []int{2, 3}
	adjacencyList.List[2] = []int{1, 4, 5}
	adjacencyList.List[3] = []int{1, 6}
	adjacencyList.List[4] = []int{2}
	adjacencyList.List[5] = []int{2}
	adjacencyList.List[6] = []int{3, 7, 8, 9}
	adjacencyList.List[7] = []int{6}
	adjacencyList.List[8] = []int{6}
	adjacencyList.List[9] = []int{6}
	adjacencyList.Preorder(1, 0)
}
