package graph

import (
	"fmt"

	"github.com/hmeng-19/goutils/queue"
)

type Edge struct {
	V1, V2 int
}

type Graph struct {
	Vertices       map[int]int
	Edges          map[Edge]int
	adjList        map[int][]int
	adjMatrix      [][]int
	Vertices2Index map[int]int
}

func New(vertices map[int]int, edges map[Edge]int) Graph {
	return Graph{
		Vertices: vertices,
		Edges:    edges,
	}
}

func (g *Graph) InitAdjList() {
	g.adjList = map[int][]int{}
	for edge := range g.Edges {
		if _, ok := g.adjList[edge.V1]; !ok {
			g.adjList[edge.V1] = []int{}
		}
		if _, ok := g.adjList[edge.V2]; !ok {
			g.adjList[edge.V2] = []int{}
		}
		g.adjList[edge.V1] = append(g.adjList[edge.V1], edge.V2)
		g.adjList[edge.V2] = append(g.adjList[edge.V2], edge.V1)
	}
}

func (g *Graph) InitAdjMatrix() {
	n := len(g.Vertices)
	index := 0
	g.Vertices2Index = map[int]int{}
	for v := range g.Vertices {
		g.Vertices2Index[v] = index
		index++
	}
	g.adjMatrix = make([][]int, n)
	for i := 0; i < n; i++ {
		g.adjMatrix[i] = make([]int, n)
	}
	for edge := range g.Edges {
		g.adjMatrix[g.Vertices2Index[edge.V1]][g.Vertices2Index[edge.V2]] = 1
		g.adjMatrix[g.Vertices2Index[edge.V2]][g.Vertices2Index[edge.V1]] = 1
	}
}

func (g *Graph) BFS(start int, visited *map[int]bool, result *[]int) {
	q := queue.New()
	fmt.Printf("enqueue %v\n", start)
	q.EnQueue(start)
	*result = append(*result, start)
	(*visited)[start] = true
	for {
		x, err := q.DeQueue()
		if err != nil {
			return
		}

		for _, next := range g.adjList[x] {
			if _, ok := (*visited)[next]; ok {
				continue
			}

			fmt.Printf("enqueue %v\n", next)
			q.EnQueue(next)
			*result = append(*result, next)
			(*visited)[next] = true
		}
	}
}

func (g *Graph) DFS(start int, visited *map[int]bool, result *[]int) {
	if _, ok := (*visited)[start]; ok {
		return
	}

	fmt.Printf("visiting %v\n", start)
	*result = append(*result, start)
	(*visited)[start] = true
	for _, next := range g.adjList[start] {
		g.DFS(next, visited, result)
	}
}
