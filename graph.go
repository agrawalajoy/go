package main

import (
	"container/list"
	"fmt"
)

type QType = int

type Graph struct {
	nodeCount int
	adjMatrix [][]int
}

func NewGraph(nodes int) Graph {

	adjMatrix := make([][]int, nodes)
	return Graph{nodeCount: nodes, adjMatrix: adjMatrix}
}
func (g *Graph) AddEdge(a, b int) {

	g.AddDirectedEdge(a, b)
	g.AddDirectedEdge(b, a)
}

func (g *Graph) AddDirectedEdge(a, b int) {

	// Add a->b
	if g.adjMatrix[a] == nil {
		g.adjMatrix[a] = make([]int, 0)
	}

	//remove this if want faster and unsafe
	if !g.IsDirectlyConnected(a, b) {
		g.adjMatrix[a] = append(g.adjMatrix[a], b)
	}
}

func (g *Graph) IsDirectlyConnected(a, b int) bool {

	// Add a->b
	if g.adjMatrix[a] == nil {
		return false
	}

	for _, v := range g.adjMatrix[a] {
		if v == b {
			return true
		}
	}
	return false

	//Add b->a
}

func (g *Graph) IsConnected(a, b int) bool {

	if a == b {
		return true
	}
	queue := Queue{}

	queue.Init()

	visited := map[int]struct{}{}
	queue.Enqueue(a)
	visited[a] = struct{}{}

	for queue.GetLen() > 0 {
		val := queue.Dequeue()
		for _, v := range g.adjMatrix[val] {
			if v == b {
				return true
			}
			if _, ok := visited[v]; !ok {
				queue.Enqueue(v)
				visited[v] = struct{}{}
			}
		}
	}
	return false

}

type Queue struct {
	queue *list.List
}

func (q *Queue) Init() {
	q.queue = list.New()

}

func (q *Queue) Enqueue(e QType) {
	q.queue.PushBack(e)
}

func (q *Queue) Dequeue() QType {

	front := q.queue.Front()
	q.queue.Remove(front)
	return front.Value.(QType)
}

func (q *Queue) GetFront() QType {
	front := q.queue.Front()
	return front.Value.(QType)
}

func (q *Queue) GetLen() int {
	return q.queue.Len()
}

////////// TEST

// Graph
func validPath(n int, edges [][]int, source int, destination int) bool {
	// n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
	graph := NewGraph(n)
	for _, e := range edges {
		graph.AddEdge(e[0], e[1])
	}
	return graph.IsConnected(source, destination)

}

func main() {
	n := 50
	edges := [][]int{{31, 5}, {10, 46}, {19, 31}, {5, 1}, {31, 28}, {28, 29}, {8, 26}, {13, 23}, {16, 34}, {30, 1}, {16, 18}, {33, 46}, {27, 35}, {2, 25}, {49, 33}, {44, 19}, {22, 26}, {30, 13}, {27, 12}, {8, 16}, {42, 13}, {18, 3}, {21, 20}, {2, 17}, {5, 48}, {41, 37}, {39, 37}, {2, 11}, {20, 26}, {19, 43}, {45, 7}, {0, 21}, {44, 23}, {2, 39}, {27, 36}, {41, 48}, {17, 42}, {40, 32}, {2, 28}, {35, 38}, {3, 9}, {41, 30}, {5, 11}, {24, 22}, {39, 5}, {40, 31}, {18, 35}, {23, 39}, {20, 24}, {45, 12}}

	source := 29
	destination := 46

	fmt.Println("validPath", validPath(n, edges, source, destination))
}
