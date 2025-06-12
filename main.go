package main

import "fmt"

type Graph struct {
	n     int
	edges [][2]int
}

func k33() Graph {
	k3 := Graph{}
	k3.n = 6

	k3.edges = [][2]int{
		[2]int{0, 3},
		[2]int{0, 4},
		[2]int{0, 5},
		[2]int{1, 3},
		[2]int{1, 4},
		[2]int{1, 5},
		[2]int{2, 3},
		[2]int{2, 4},
		[2]int{2, 5},
	}
	return k3
}

func k(n int) Graph {
	g := Graph{}
	g.n = n
	edges := [][2]int{}
	for i := range n {
		for j := range n {
			if i != j && j > i {
				edges = append(edges, [2]int{i, j})
			}
		}

	}
	g.edges = edges
	return g

}

func main() {
	//g := k33()
	g := k(4)
	fmt.Println(g.dot())
	fmt.Println(g.Complement().dot())
	fmt.Println(g.L().dot())
	fmt.Println(g.L().Complement().dot())
}

func (g Graph) Complement() Graph {
	h := k(g.n)
	left := map[[2]int]bool{}
	for _, v := range h.edges {
		left[v] = true
	}
	for _, v := range g.edges {
		delete(left, v)
	}
	edges := [][2]int{}
	for k, _ := range left {
		edges = append(edges, k)
	}
	h.edges = edges
	return h
}

func (g Graph) L() Graph {
	h := Graph{}
	h.n = len(g.edges)
	// eek
	// find all edges that are adjacent to node i and put them in adj
	adj := map[int][]int{}
	for i, e := range g.edges {
		// if two edges share a node in G
		// -- that is the node is adjacent to these two edges
		// in the Line Graph that is edge between the new vertices (those edges became)
		current := adj[e[0]]
		if current == nil {
			current = []int{}
		}
		current = append(current, i)
		adj[e[0]] = current
		current = adj[e[1]]
		if current == nil {
			current = []int{}
		}
		current = append(current, i)
		adj[e[1]] = current

	}

	edges := [][2]int{}
	for _, v := range adj {

		for _, i := range v {
			for _, j := range v {
				if i != j && j > i {
					edges = append(edges, [2]int{i, j})
				}
			}
		}
	}
	h.edges = edges

	return h
}

func (g Graph) dot() string {
	fmt.Println("graph example {")
	for i, e := range g.edges {
		fmt.Printf(" %d -- %d [label=%d]\n", e[0], e[1], i)
	}
	fmt.Println("}")

	return ""
}
