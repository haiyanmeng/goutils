package main

import "fmt"

type Node struct {
	data                int
	left, right, parent *Node
}

func pathToRoot(a *Node) []*Node {
	cur := a
	result := []*Node{}
	for cur != nil {
		result = append(result, cur)
		cur = cur.parent
	}
	return result
}

func LCA(a, b *Node) *Node {
	p1, p2 := pathToRoot(a), pathToRoot(b)
	n1, n2 := len(p1), len(p2)
	var node *Node
	for n1 > 0 && n2 > 0 && (p1[n1-1] == p2[n2-1]) {
		node = p1[n1-1]
		n1--
		n2--
	}
	return node
}

func distToRoot(a *Node) int {
	d := 0
	for a != nil {
		d++
		a = a.parent
	}
	return d
}

func LCA1(a, b *Node) *Node {
	d1, d2 := distToRoot(a), distToRoot(b)
	diff := d1 - d2
	lower, higher := a, b
	if diff < 0 {
		lower, higher = b, a
		diff = -diff
	}
	// let the lower node get to the same level with the higher node
	for i := 0; i < diff; i++ {
		lower = lower.parent
	}
	// find the common ancestor
	for lower != higher {
		lower = lower.parent
		higher = higher.parent
	}
	return lower
}

func main() {
	nodes := []*Node{}
	for i := 0; i < 11; i++ {
		nodes = append(nodes, &Node{i + 1, nil, nil, nil})
	}
	nodes[0].left, nodes[0].right, nodes[0].parent = nodes[1], nodes[2], nil
	nodes[1].left, nodes[1].right, nodes[1].parent = nodes[3], nodes[4], nodes[0]
	nodes[2].left, nodes[2].right, nodes[2].parent = nodes[5], nodes[6], nodes[0]
	nodes[3].left, nodes[3].right, nodes[3].parent = nodes[7], nodes[8], nodes[1]
	nodes[4].left, nodes[4].right, nodes[4].parent = nil, nil, nodes[1]
	nodes[5].left, nodes[5].right, nodes[5].parent = nil, nil, nodes[2]
	nodes[6].left, nodes[6].right, nodes[6].parent = nodes[9], nodes[10], nodes[2]
	nodes[7].left, nodes[7].right, nodes[7].parent = nil, nil, nodes[3]
	nodes[8].left, nodes[8].right, nodes[8].parent = nil, nil, nodes[3]
	nodes[9].left, nodes[9].right, nodes[9].parent = nil, nil, nodes[6]
	nodes[10].left, nodes[10].right, nodes[10].parent = nil, nil, nodes[6]
	fmt.Println(LCA(nodes[6], nodes[10]))
	fmt.Println(LCA1(nodes[6], nodes[10]))
	fmt.Println(LCA(nodes[7], nodes[10]))
	fmt.Println(LCA1(nodes[7], nodes[10]))

}
