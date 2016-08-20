package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/graph"
)

func main() {
	g := graph.New(
		map[int]int{1: 5, 2: 6, 3: 7, 4: 8, 5: 9},
		map[graph.Edge]int{
			graph.Edge{1, 2}: 3,
			graph.Edge{2, 3}: 5,
			graph.Edge{3, 4}: 7,
			graph.Edge{4, 5}: 9,
			graph.Edge{1, 4}: 9,
		},
	)

	g.InitAdjList()
	g.InitAdjMatrix()

	fmt.Printf("%#v\n", g)

	visited := map[int]bool{}
	result := []int{}
	g.DFS(1, &visited, &result)
	fmt.Printf("dfs visited: %v\n", visited)
	fmt.Printf("dfs: %v\n", result)

	visited = map[int]bool{}
	result = []int{}
	g.BFS(1, &visited, &result)
	fmt.Printf("bfs visited: %v\n", visited)
	fmt.Printf("bfs: %v\n", result)
}
