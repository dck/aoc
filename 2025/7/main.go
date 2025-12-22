package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	i, j int
}

func main() {
	grid := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	fmt.Println("Part 1:", part1(grid))
	fmt.Println("Part 2:", part2(grid))
}

func part1(grid []string) int {
	var start Point
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				start = Point{i, j}
			}
		}
	}

	queue := []Point{start}
	visited := make(map[Point]bool)
	splits := 0

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.i < 0 || item.i >= len(grid) || item.j < 0 || item.j >= len(grid[item.i]) {
			continue
		}

		if visited[item] {
			continue
		}
		visited[item] = true

		current := grid[item.i][item.j]

		switch current {
		case '.', 'S':
			queue = append(queue, Point{item.i + 1, item.j})
		case '^':
			queue = append(queue, Point{item.i, item.j - 1})
			queue = append(queue, Point{item.i, item.j + 1})
			splits++
		}
	}

	return splits
}

func part2(grid []string) int {
	var start Point
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				start = Point{i, j}
			}
		}
	}

	memo := make(map[Point]int)

	var dfs func(p Point) int
	dfs = func(p Point) int {
		if p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[p.i]) {
			return 0
		}
		if val, ok := memo[p]; ok {
			return val
		}
		if p.i == len(grid)-1 {
			memo[p] = 1
			return 1
		}

		current := grid[p.i][p.j]
		var res int
		switch current {
		case '.', 'S':
			res = dfs(Point{p.i + 1, p.j})
		case '^':
			res = dfs(Point{p.i, p.j - 1}) + dfs(Point{p.i, p.j + 1})
		default:
			res = 0
		}
		memo[p] = res
		return res
	}

	return dfs(start)
}
