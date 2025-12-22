package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
)

type DisjointSet []int

func NewDisjointSet(size int) DisjointSet {
	ds := make(DisjointSet, size)
	for i := range ds {
		ds[i] = -1
	}
	return ds
}

func (ds DisjointSet) Find(i int) int {
	if ds[i] < 0 {
		return i
	}
	ds[i] = ds.Find(ds[i])
	return ds[i]
}

func (ds DisjointSet) Union(i, j int) int {
	ri := ds.Find(i)
	rj := ds.Find(j)
	if ri == rj {
		return -ds[ri]
	}
	if ds[ri] < ds[rj] {
		ds[ri] += ds[rj]
		ds[rj] = ri
		return -ds[ri]
	} else {
		ds[rj] += ds[ri]
		ds[ri] = rj
		return -ds[rj]
	}
}

type Pair struct {
	I, J     int
	Distance int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].Distance > h[j].Distance }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type Point struct {
	x, y, z int
}

func main() {
	points := []Point{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		points = append(points, Point{x, y, z})
	}
	fmt.Println("Part 1:", part1(points))
	fmt.Println("Part 2:", part2(points))
}

func part1(points []Point) int {
	size := len(points)
	ds := NewDisjointSet(size)

	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d := distance(points[i], points[j])

			if h.Len() < 1000 {
				heap.Push(h, Pair{i, j, d})
			} else if d < (*h)[0].Distance {
				heap.Pop(h)
				heap.Push(h, Pair{i, j, d})
			}
		}
	}

	for h.Len() > 0 {
		pair := heap.Pop(h).(Pair)
		ds.Union(pair.I, pair.J)
	}

	groups := make(map[int]int)
	for i := 0; i < size; i++ {
		root := ds.Find(i)
		groups[root] += 1
	}

	sizes := make([]int, len(groups))
	for size := range maps.Values(groups) {
		sizes = append(sizes, size)
	}
	slices.Sort(sizes)
	largestSizes := sizes[len(sizes)-3:]
	return largestSizes[0] * largestSizes[1] * largestSizes[2]
}

func part2(points []Point) int {
	size := len(points)
	ds := NewDisjointSet(size)

	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d := distance(points[i], points[j])
			heap.Push(h, Pair{i, j, -d})
		}
	}

	for h.Len() > 0 {
		pair := heap.Pop(h).(Pair)
		groupSize := ds.Union(pair.I, pair.J)

		if groupSize == size {
			return points[pair.I].x * points[pair.J].x
		}
	}

	return -1
}

func distance(p1, p2 Point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return int(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}
