package main

import (
	"fmt"
	"sort"
)

type interval struct {
	start, end int
}

type point struct {
	time       int
	start      bool
	intervalID int
}

type points []point

func (p points) Len() int {
	return len(p)
}

func (p points) Less(i, j int) bool {
	if p[i].time < p[j].time {
		return true
	} else if p[i].time > p[j].time {
		return false
	} else {
		// p[i].time == p[j].time
		if p[i].start == true {
			return true
		}
		return false
	}
}

func (p points) Swap(i, j int) {
	t := p[i]
	p[i] = p[j]
	p[j] = t
}

func convert2Points(intervals []interval) points {
	var p points = []point{}
	for i, interval := range intervals {
		p = append(p, point{
			time:       interval.start,
			start:      true,
			intervalID: i,
		})

		p = append(p, point{
			time:       interval.end,
			start:      false,
			intervalID: i,
		})
	}
	return p
}

func converage(p points) []int {
	result := []int{}
	visiting := []int{}
	visited := make(map[int]bool)
	for _, item := range p {
		if item.start {
			visiting = append(visiting, item.intervalID)
			continue
		}
		if _, ok := visited[item.intervalID]; ok {
			continue
		}

		visiting = append(visiting, item.intervalID)
		for _, x := range visiting {
			visited[x] = true
		}
		visiting = []int{}
		result = append(result, item.time)
	}
	return result
}

func main() {
	intervals := []interval{
		interval{0, 3},
		interval{2, 6},
		interval{3, 4},
		interval{6, 9},
	}

	p := convert2Points(intervals)
	sort.Sort(p)
	fmt.Printf("The sorted endpoints: %v\n", p)
	a := converage(p)
	fmt.Println(a)
}
