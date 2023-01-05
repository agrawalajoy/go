// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

type any interface{}

// An MinIntHeap is a min-heap of ints.
type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxIntHeap) Peek() any {
	// Check for Len before calling this
	return h[0]
}

// An MaxIntHeap is a min-heap of ints.
type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ====== findKthLargest from Heap ====

func findKthLargest(nums []int, k int) int {

	hmax := MaxIntHeap(nums)
	// copy
	// hmax := make([]MaxIntHeap, len(nums))
	// copy(cpy, orig)

	heap.Init(&hmax)

	if k > hmax.Len() {
		k = hmax.Len()
	}
	for i := 0; i < k-1; i++ {
		heap.Pop(&hmax)
	}
	return heap.Pop(&hmax).(int)

}

// This example inserts several ints into an MinIntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	hmin := &MinIntHeap{2, 1, 5}
	heap.Init(hmin)
	heap.Push(hmin, 3)
	fmt.Printf("minimum: %d\n", (*hmin)[0])
	for hmin.Len() > 0 {
		fmt.Printf("MIN %d ", heap.Pop(hmin))
	}

	hmax := &MaxIntHeap{2, 1, 5}
	heap.Init(hmax)
	heap.Push(hmax, 3)
	fmt.Printf("maximum: %d\n", (*hmax)[0])
	for hmax.Len() > 0 {
		fmt.Printf("MAX %d ", heap.Pop(hmax))
	}

}
