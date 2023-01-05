package main

import (
	"container/heap"
	"fmt"
)

func main() {
	n := 4
	roads := [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}
	fmt.Println("Dijkstra")
	// Example
	graph := newGraph()

	for _, road := range roads {
		graph.addEdge(road[0], road[1], road[2])

	}

	// graph.addEdge("S", "B", 4)
	// graph.addEdge("S", "C", 2)
	// graph.addEdge("B", "C", 1)
	// graph.addEdge("B", "D", 5)
	// graph.addEdge("C", "D", 8)
	// graph.addEdge("C", "E", 10)
	// graph.addEdge("D", "E", 2)
	// graph.addEdge("D", "T", 6)
	// graph.addEdge("E", "T", 2)
	fmt.Println(graph.getPath(1, n))
}

type edge struct {
	node   int
	weight int
}

type graph struct {
	nodes map[int][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[int][]edge)}
}

func (g *graph) addEdge(origin, destiny int, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node int) []edge {
	return g.nodes[node]
}

func (g *graph) getPath(origin, destiny int) (int, []int) {
	h := newHeap()
	h.push(path{value: 0, nodes: []int{origin}})
	visited := make(map[int]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append([]int{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}

type path struct {
	value int
	nodes []int
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type myHeap struct {
	values *minPath
}

func newHeap() *myHeap {
	return &myHeap{values: &minPath{}}
}

func (h *myHeap) push(p path) {
	heap.Push(h.values, p)
}

func (h *myHeap) pop() path {
	i := heap.Pop(h.values)
	return i.(path)
}
