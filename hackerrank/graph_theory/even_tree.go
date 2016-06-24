package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Result int

type Reader struct {
	scanner *bufio.Scanner
	nextInt []int
}

func (r *Reader) NextInt() int {
	if len(r.nextInt) > 0 {
		res := r.nextInt[0]
		r.nextInt = r.nextInt[1:]
		return res
	}

	r.scanner.Scan()
	line := strings.TrimSpace(r.scanner.Text())

	for _, v := range strings.Split(line, " ") {
		number, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		r.nextInt = append(r.nextInt, number)
	}

	res := r.nextInt[0]
	r.nextInt = r.nextInt[1:]

	return res
}

type AdjacencyList struct {
	List          map[int][]int
	Visited       map[int]bool
	TraversalPath []int
	Marked        map[int]bool
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

func (al *AdjacencyList) Solve(vert int) {
	tPaths := [][]int{}
	for _, av := range al.List[vert] {
		for i, v := range al.List[av] {
			if v == vert {
				newS := al.List[av][:i]
				newS = append(newS, al.List[av][i+1:]...)
				al.List[av] = newS
			}
		}
		al.TraversalPath = []int{}
		al.Visited = map[int]bool{}

		if _, ok := al.Marked[av]; !ok {
			al.DFS(av)
			al.Marked[av] = true
		}

		l := len(al.TraversalPath)
		if l > 0 {
			tPaths = append(tPaths, al.TraversalPath)
			if l%2 == 0 {
				Result++
			}
		}
	}

	for _, tPath := range tPaths {
		for _, v := range tPath {
			al.Solve(v)
		}
	}
}

func main() {
	reader := &Reader{
		scanner: bufio.NewScanner(os.Stdin),
		nextInt: []int{},
	}

	aList := &AdjacencyList{
		List:    map[int][]int{},
		Visited: map[int]bool{},
		Marked:  map[int]bool{},
	}

	reader.NextInt()
	M := reader.NextInt()

	for i := 0; i < M; i++ {
		u, v := reader.NextInt(), reader.NextInt()
		aList.List[u] = append(aList.List[u], v)
		aList.List[v] = append(aList.List[v], u)
	}

	aList.Solve(1)
	fmt.Println(Result)
}
