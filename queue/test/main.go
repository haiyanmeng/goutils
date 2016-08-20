package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/queue"
)

func main() {
	q := queue.New()
	for i := 0; i < 10; i++ {
		q.EnQueue(i)
	}

	for i := 0; i < 10; i++ {
		x, err := q.DeQueue()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(x)
	}

}
