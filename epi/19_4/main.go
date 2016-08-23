package main

import "fmt"

type Vertex string

type Edge struct {
	Start, End Vertex
}

type Graph struct {
	Vertices []Vertex
	Edges    []Edge
	AdjList  map[Vertex][]Vertex
}

func New(vertices []string, edges map[string]string) Graph {
	g := Graph{}
	for _, s := range vertices {
		g.Vertices = append(g.Vertices, Vertex(s))
	}

	g.AdjList = map[Vertex][]Vertex{}
	for start, end := range edges {
		g.Edges = append(g.Edges, Edge{Vertex(start), Vertex(end)})
		g.AdjList[Vertex(start)] = append(g.AdjList[Vertex(start)], Vertex(end))
	}
	return g
}

func (g Graph) GetStartVertices() []Vertex {
	ends := map[Vertex]bool{}
	for _, v := range g.AdjList {
		for _, vertex := range v {
			ends[vertex] = true
		}
	}
	starts := []Vertex{}
	for _, vertex := range g.Vertices {
		if _, ok := ends[vertex]; !ok {
			starts = append(starts, vertex)
		}
	}
	return starts
}

func (g Graph) DFS(start Vertex, visited map[Vertex]bool) bool {
	if _, ok := visited[start]; ok {
		return true
	}

	visited[start] = true
	for _, next := range g.AdjList[start] {
		if g.DFS(next, visited) {
			return true
		}
	}
	return false
}

func (g Graph) HasCircle() bool {
	visited := map[Vertex]bool{}
	for _, start := range g.Vertices {
		if g.DFS(start, visited) {
			return true
		}
	}
	return false
}

func main() {
	g := New([]string{"p1", "p2", "p3", "p4", "p5", "p6", "p7"},
		map[string]string{
			"p1": "p2",
			"p2": "p3",
			"p3": "p4",
			"p4": "p5",
			//"p5": "p3",
			"p6": "p7",
			"p7": "p6",
		})
	fmt.Printf("%#v\n", g)
	starts := g.GetStartVertices()
	fmt.Printf("the starts of the graph: %#v\n", starts)
	fmt.Printf("Has Circle? %v\n", g.HasCircle())
}
