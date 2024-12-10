package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := make([][]byte, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]byte, 0)
		for _, c := range line {
			row = append(row, byte(c)-'0')
		}
		grid = append(grid, row)
	}
	fmt.Println("Part 1:", part1(grid))
	fmt.Println("Part 2:", part2(grid))
}

func part1(grid [][]byte) int {
	total := 0
	for i, row := range grid {
		for j, num := range row {
			if num == 0 {
				total += scoreAt(i, j, grid, true)
			}
		}
	}
	return total
}

func part2(grid [][]byte) int {
	total := 0
	for i, row := range grid {
		for j, num := range row {
			if num == 0 {
				total += scoreAt(i, j, grid, false)
			}
		}
	}
	return total
}

func scoreAt(i, j int, grid [][]byte, checkUniq bool) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	visited[i][j] = true
	score := 0
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{i, j})
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if grid[pos[0]][pos[1]] == 9 {
			score++
		}

		for _, dir := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newI := pos[0] + dir[0]
			newJ := pos[1] + dir[1]
			if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) {
				continue
			}
			if checkUniq && visited[newI][newJ] {
				continue
			}

			if grid[newI][newJ]-grid[pos[0]][pos[1]] != 1 {
				continue
			}

			visited[newI][newJ] = true
			queue = append(queue, [2]int{newI, newJ})
		}
	}
	return score
}
