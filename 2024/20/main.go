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

func (m *Map) Print(path []PathPoint) {
	pathMap := map[Point]bool{}
	for _, p := range path {
		pathMap[p.pos] = true
	}

	for i := 0; i < len(m.grid); i++ {
		for j := 0; j < len(m.grid[i]); j++ {
			if m.start.i == i && m.start.j == j {
				fmt.Print("S")
			} else if m.end.i == i && m.end.j == j {
				fmt.Print("E")
			} else if pathMap[Point{i, j}] {
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
	cost     int
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

func main() {
	m := &Map{}
	grid := make([][]byte, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				m.start = Point{i, j}
			}
			if cell == 'E' {
				m.end = Point{i, j}
			}
		}
	}

	m.grid = grid

	fmt.Println("Part 1:", part1(m))
	fmt.Println("Part 2:", part2(m))
}

func part1(m *Map) int {
	path := astar(m)

	res := 0

	for i := 0; i < len(path)-2; i++ {
		for j := i + 2; j < len(path); j++ {
			left := path[i]
			right := path[j]

			if left.pos.i == right.pos.i || left.pos.j == right.pos.j {

				midI := (left.pos.i + right.pos.i) / 2
				midJ := (left.pos.j + right.pos.j) / 2

				if m.grid[midI][midJ] == '#' && manhattan(left.pos, right.pos) == 2 {
					if right.cost-left.cost-2 >= 100 {
						res++
					}
				}
			}
		}
	}
	return res
}

func part2(m *Map) int {
	path := astar(m)

	res := 0
	for i := 0; i < len(path)-2; i++ {
		for j := i + 2; j < len(path); j++ {
			left := path[i]
			right := path[j]

			dist := manhattan(left.pos, right.pos)
			if dist <= 20 {
				saved := right.cost - left.cost - dist
				if saved >= 100 {
					res++
				}
			}
		}
	}

	return res
}

func astar(m *Map) []PathPoint {
	pq := &PQ{}
	heap.Init(pq)

	start := PathPoint{0, 0, m.start}
	heap.Push(pq, start)

	costs := map[Point]int{}
	previous := map[Point]Point{}

	costs[m.start] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(PathPoint)

		if current.pos == m.end {
			path := []PathPoint{current}
			curr := current.pos
			for curr != m.start {
				curr = previous[curr]
				path = append(path, PathPoint{0, costs[curr], curr})
			}
			reverse(path)
			return path
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
				heap.Push(pq, PathPoint{priority, newCost, next})
				previous[next] = current.pos
			}
		}
	}

	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverse[A any](arr []A) {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func manhattan(a, b Point) int {
	return abs(a.i-b.i) + abs(a.j-b.j)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
