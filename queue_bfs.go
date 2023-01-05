package main

import (
	"container/list"
	"fmt"
)

type QType = int

type QStruct struct {
	elem  QType
	level int
}

func NewQueue() *Queue {
	var q Queue
	q.Init()

	return &q
}

type Queue struct {
	queue *list.List
}

func (q *Queue) Init() {
	q.queue = list.New()

}

func (q *Queue) Enqueue(e QType, level int) {
	q.queue.PushBack(QStruct{e, level})
}

func (q *Queue) Dequeue() (QType, int) {

	front := q.queue.Front()
	q.queue.Remove(front)
	f := front.Value.(QStruct)
	return f.elem, f.level
}

func (q *Queue) GetFront() (QType, int) {
	front := q.queue.Front()
	f := front.Value.(QStruct)
	return f.elem, f.level
}

func (q *Queue) GetLen() int {
	return q.queue.Len()
}

func main() {
	// new linked list

	queue := Queue{}

	queue.Init()

	// Simply append to enqueue.
	queue.Enqueue(10, 0)
	queue.Enqueue(20, 1)
	queue.Enqueue(30, 2)

	for queue.GetLen() > 0 {
		fmt.Println("PEEK")
		fmt.Println(queue.GetFront())
		fmt.Println("REM")
		fmt.Println(queue.Dequeue())
	}
}
