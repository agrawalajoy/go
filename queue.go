package main

import (
	"container/list"
	"fmt"
)

type QType = int
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

func main() {
	// new linked list

	queue := Queue{}

	queue.Init()

	// Simply append to enqueue.
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	for queue.GetLen() > 0 {
		fmt.Println("PEEK", queue.GetFront())
		fmt.Println("REM", queue.Dequeue())
	}
}
