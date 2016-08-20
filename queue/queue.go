package queue

import "fmt"

type Queue struct {
	arr []int
	n   int
}

func New() Queue {
	return Queue{
		arr: []int{},
	}
}

func (q *Queue) EnQueue(x int) {
	q.arr = append(q.arr, x)
	q.n++
}

func (q *Queue) DeQueue() (int, error) {
	if q.n == 0 {
		return 0, fmt.Errorf("the queue is empty")
	}

	x := q.arr[0]
	q.arr = q.arr[1:]
	q.n--
	return x, nil
}
