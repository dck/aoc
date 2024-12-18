package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	UP byte = iota
	DOWN
	LEFT
	RIGHT
)

const Reset string = "\033[0m"
const Red string = "\033[31m"
const Blue string = "\033[34m"

var directions = map[byte][2]int{
	UP:    {-1, 0},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
	RIGHT: {0, 1},
}

type Point struct {
	i, j int
}

type Map struct {
	grid  [][]byte
	start Point
	end   Point
}

func (m *Map) Print(shortestPath map[Point]bool) {
	for i := 0; i < len(m.grid); i++ {
		for j := 0; j < len(m.grid[i]); j++ {
			if m.start.i == i && m.start.j == j {
				fmt.Print("S")
			} else if m.end.i == i && m.end.j == j {
				fmt.Print("E")
			} else if shortestPath[Point{i, j}] {
				fmt.Print(Blue + "O" + Reset)
			} else if m.grid[i][j] == '#' {
				fmt.Print(Red + "#" + Reset)
			} else {
				fmt.Print(string(m.grid[i][j]))
			}
		}
		fmt.Println()
	}
}

type PathPoint struct {
	priority int
	pos      Point
}

type PQ []PathPoint

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	item := x.(PathPoint)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Result struct {
	score int
	paths map[Point]bool
}

func main() {
	m := &Map{}
	bytes := []Point{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var x, y int
		line := scanner.Text()

		fmt.Sscanf(line, "%d,%d\n", &x, &y)
		bytes = append(bytes, Point{y, x})
	}

	fmt.Println("Part 1:", part1(m, bytes))
	fmt.Println("Part 2:", part2(m, bytes))
}

func part1(m *Map, bytes []Point) int {
	m.grid = buildGrid(bytes, 1024)
	m.start = Point{0, 0}
	m.end = Point{len(m.grid) - 1, len(m.grid[0]) - 1}
	result := astar(m)
	return result.score
}

func part2(m *Map, bytes []Point) Point {
	left, right := 0, len(bytes)-1

	for left < right {
		mid := (left + right) / 2
		m.grid = buildGrid(bytes, mid)
		m.start = Point{0, 0}
		m.end = Point{len(m.grid) - 1, len(m.grid[0]) - 1}
		result := astar(m)

		if result.score == -1 {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return bytes[right]
}

func astar(m *Map) Result {
	manhattan := func(a, b Point) int {
		return abs(a.i-b.i) + abs(a.j-b.j)
	}

	pq := &PQ{}
	heap.Init(pq)

	start := PathPoint{0, m.start}
	heap.Push(pq, start)

	costs := map[Point]int{}
	previous := map[Point]Point{}

	costs[m.start] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(PathPoint)

		if current.pos == m.end {
			paths := map[Point]bool{}
			score := 0
			curr := m.end
			for curr != m.start {
				curr = previous[curr]
				paths[curr] = true
				score++
			}
			return Result{score, paths}
		}

		for _, dir := range []byte{UP, DOWN, LEFT, RIGHT} {
			next := Point{current.pos.i + directions[dir][0], current.pos.j + directions[dir][1]}

			if next.i < 0 || next.i >= len(m.grid) || next.j < 0 || next.j >= len(m.grid[next.i]) {
				continue
			}

			if m.grid[next.i][next.j] == '#' {
				continue
			}

			newCost := costs[current.pos] + 1
			if _, ok := costs[next]; !ok || newCost < costs[next] {
				costs[next] = newCost
				priority := newCost + manhattan(next, m.end)
				heap.Push(pq, PathPoint{priority, next})
				previous[next] = current.pos
			}
		}
	}

	return Result{-1, nil}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func buildGrid(bytes []Point, howMany int) [][]byte {
	grid := make([][]byte, 71)
	for i := range grid {
		grid[i] = make([]byte, 71)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for i, b := range bytes {
		grid[b.i][b.j] = '#'
		if i >= howMany {
			break
		}
	}
	return grid
}
