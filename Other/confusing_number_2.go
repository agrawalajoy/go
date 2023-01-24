package main

import (
	"container/list"
)

func confusingNumber(n int) bool {
	hash := map[int]int{9: 6, 6: 9, 0: 0, 1: 1, 8: 8}
	orig := n
	rev := 0

	for orig > 0 {

		o := orig % 10

		if v, ok := hash[o]; ok {
			rev = rev*10 + v
		} else {

			return false
		}
		orig = orig / 10

	}
	return n != rev
}

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

func confusingNumberII(n int) int {

	digits := getDigitCount(n)

	numbers := []int{9, 6, 0, 1, 8}
	numbers_0 := []int{9, 6, 1, 8}
	count := 0
	if n == 9 {
		return 1
	}

	queue := Queue{}

	queue.Init()

	for _, j := range numbers_0 {
		if j <= n {

			if j == 9 || j == 6 {
				count++
			}
			queue.Enqueue(j, 1)
		}
	}

	for queue.GetLen() > 0 {

		val, d := queue.Dequeue()
		if d < digits {

			for _, j := range numbers {

				if val*10+j <= n {
					queue.Enqueue(val*10+j, d+1)

					if confusingNumber(val*10 + j) {
						count++
					}

				}
			}
		}
	}

	return count
}

func getDigitCount(n int) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}
