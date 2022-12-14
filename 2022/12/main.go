package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func (p Point) AddPoint(a Point) Point {
	return Point{x: p.x + a.x, y: p.y + a.y}
}

type Grid [][]int

func (g Grid) GetValue(p Point) int {
	return g[p.x][p.y]
}

func (g Grid) SetValue(p Point, v int) {
	g[p.x][p.y] = v
}

func (g Grid) IsInside(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(g) && p.y < len(g[0])
}

func main() {
	grid := Grid{}
	//var start Point
	var goal Point
	scanner := bufio.NewScanner(os.Stdin)

	x := 0
	for scanner.Scan() {
		line := scanner.Text()

		row := []int{}
		for y, c := range line {
			if c == 'S' {
				row = append(row, 1)
				// start = Point{x, y}

			} else if c == 'E' {
				row = append(row, 'z'-'a'+2)
				goal = Point{x, y}
			} else {
				row = append(row, int(c)-'a'+1)
			}
		}

		grid = append(grid, row)
		x++
	}
	//fmt.Println(bfs(grid, start, 'z'-'a'+2))
	fmt.Println(bfs(grid, goal, 1))
}

func bfs(grid Grid, start Point, target int) (steps int) {
	directions := []Point{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	queue := []struct {
		p     Point
		score int
	}{{start, 0}}
	visited := map[Point]bool{}

	for len(queue) > 0 {
		current := queue[0]
		pos := current.p
		score := current.score
		queue = queue[1:]

		if grid.GetValue(pos) == target {
			return score
		}

		if _, ok := visited[pos]; ok {
			continue
		}
		visited[pos] = true

		for _, d := range directions {
			newPoint := pos.AddPoint(d)

			if grid.IsInside(newPoint) && grid.GetValue(pos)-grid.GetValue(newPoint) <= 1 {
				queue = append(queue, struct {
					p     Point
					score int
				}{newPoint, score + 1})
			}
		}
	}

	return -1
}
