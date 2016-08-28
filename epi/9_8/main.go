package main

import "fmt"

type Queue struct {
	data       []int
	n          int
	start, end int
}

func New(n int) *Queue {
	return &Queue{
		data:  make([]int, n, n),
		n:     0,
		start: 0,
		end:   0,
	}
}

func (q *Queue) Count() int {
	return q.n
}

func (q *Queue) Dequeue() (int, error) {
	if q.n == 0 {
		return 0, fmt.Errorf("the queue is empty")
	}
	x := q.data[q.start]
	q.n--
	q.start++
	return x, nil
}

func (q *Queue) Enqueue(x int) {
	if q.n < cap(q.data) {
		if q.end >= cap(q.data) {
			q.end = 0
		}
	} else {
		// enlarge the underlying alice
		// copy the data from the old one to the new one
		newData := make([]int, 2*q.n, 2*q.n)
		for i := 0; i < q.n; i++ {
			newData[i] = q.data[(q.start+i)%q.n]
		}
		q.data = newData
		q.start = 0
		q.end = q.n
	}
	fmt.Printf("q.data: %v; q.end: %v\n", q.data, q.end)
	q.data[q.end] = x
	q.end++
	q.n++
}

func (q *Queue) Print() {
	fmt.Printf("the queue storage: len = %v; q.n = %v; q.data = %v\n", cap(q.data), q.n, q.data)
}

func main() {
	q := New(5)
	fmt.Println(q.data)
	for i := 0; i < 5; i++ {
		q.Enqueue(i + 1)
	}

	for i := 0; i < 2; i++ {
		x, _ := q.Dequeue()
		fmt.Println(x)
	}
	fmt.Printf("the queue storage: len = %v; q.n = %v; q.data = %v\n", cap(q.data), q.n, q.data)
	q.Enqueue(6)
	q.Enqueue(7)
	fmt.Printf("the queue storage: len = %v; q.n = %v; q.data = %v\n", cap(q.data), q.n, q.data)
	q.Enqueue(8)
	fmt.Printf("the queue storage: len = %v; q.n = %v; q.data = %v\n", cap(q.data), q.n, q.data)
}
