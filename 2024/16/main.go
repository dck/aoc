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
	score     int
	direction byte
	pos       Point
}

type PQ []PathPoint

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].score < pq[j].score
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
	grid := [][]byte{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	m.grid = grid

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				m.start = Point{i, j}
			} else if grid[i][j] == 'E' {
				m.end = Point{i, j}
			}
		}
	}

	fmt.Println("Part 1:", part1(m))
	fmt.Println("Part 2:", part2(m))
}

func part1(m *Map) int {
	return findPaths(m).score
}

func part2(m *Map) int {
	return len(findPaths(m).paths)
}

func findPaths(m *Map) Result {
	var allPathPoints map[Point]bool

	pq := &PQ{}
	heap.Init(pq)

	previous := make(map[PathPoint][]PathPoint)
	scores := make(map[PathPoint]int)
	startPoint := PathPoint{0, RIGHT, m.start}
	scores[startPoint] = 0

	heap.Push(pq, startPoint)
	shortestScore := -1
	for pq.Len() > 0 {
		pathPoint := heap.Pop(pq).(PathPoint)

		if pathPoint.pos == m.end {
			if shortestScore == -1 || pathPoint.score == shortestScore {
				shortestScore = pathPoint.score
				allPathPoints = reconstructPaths(previous, pathPoint)
			} else if pathPoint.score > shortestScore {
				break
			}
			continue
		}

		for k, v := range directions {
			newPos := Point{pathPoint.pos.i + v[0], pathPoint.pos.j + v[1]}

			if m.grid[newPos.i][newPos.j] == '#' {
				continue
			}
			if pathPoint.direction == UP && k == DOWN {
				continue
			}
			if pathPoint.direction == DOWN && k == UP {
				continue
			}
			if pathPoint.direction == LEFT && k == RIGHT {
				continue
			}
			if pathPoint.direction == RIGHT && k == LEFT {
				continue
			}

			newScore := pathPoint.score + 1
			if pathPoint.direction != k {
				newScore += 1000
			}

			newPathPoint := PathPoint{newScore, k, newPos}
			if prevScore, ok := scores[newPathPoint]; !ok || newScore < prevScore {
				scores[newPathPoint] = newScore
				previous[newPathPoint] = []PathPoint{pathPoint}
				heap.Push(pq, newPathPoint)
			} else if newScore == prevScore {
				previous[newPathPoint] = append(previous[newPathPoint], pathPoint)
			}
		}
	}

	return Result{score: shortestScore, paths: allPathPoints}
}

func reconstructPaths(previous map[PathPoint][]PathPoint, start PathPoint) map[Point]bool {
	path := make(map[Point]bool)
	queue := []PathPoint{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		path[curr.pos] = true
		for _, prev := range previous[curr] {
			queue = append(queue, prev)
		}
	}

	return path
}
