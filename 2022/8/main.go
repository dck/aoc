package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := [][]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		row := []int{}
		for _, n := range line {
			row = append(row, int(n-'0'))
		}
		grid = append(grid, row)
	}

	fmt.Println(calculateScenic(grid))
}

func calculateVisible(grid [][]int) int {
	total := 0

	for i := 1; i < len(grid)-1; i++ {
		row := grid[i]

		for j := 1; j < len(row)-1; j++ {
			tree := row[j]

			visible := true
			for k := i - 1; k >= 0; k-- {
				if grid[k][j] >= tree {
					visible = false
					break
				}
			}

			if visible {
				total += 1
				continue
			}

			visible = true
			for k := i + 1; k < len(grid); k++ {
				if grid[k][j] >= tree {
					visible = false
					break
				}
			}

			if visible {
				total += 1
				continue
			}

			visible = true
			for k := j - 1; k >= 0; k-- {
				if grid[i][k] >= tree {
					visible = false
					break
				}
			}

			if visible {
				total += 1
				continue
			}

			visible = true
			for k := j + 1; k < len(row); k++ {
				if grid[i][k] >= tree {
					visible = false
					break
				}
			}

			if visible {
				total += 1
				continue
			}
		}
	}
	return total + len(grid)*4 - 4
}

func calculateScenic(grid [][]int) int {
	highest := 0

	for i := 1; i < len(grid)-1; i++ {
		row := grid[i]

		for j := 1; j < len(row)-1; j++ {
			tree := row[j]

			viewsUp := 0
			for k := i - 1; k >= 0; k-- {
				if grid[k][j] >= tree {
					viewsUp += 1
					break
				}
				viewsUp += 1
			}

			viewsDown := 0
			for k := i + 1; k < len(grid); k++ {
				if grid[k][j] >= tree {
					viewsDown += 1
					break
				}

				viewsDown += 1
			}

			viewsLeft := 0
			for k := j - 1; k >= 0; k-- {
				if grid[i][k] >= tree {
					viewsLeft += 1
					break
				}
				viewsLeft += 1
			}

			viewsRight := 0
			for k := j + 1; k < len(row); k++ {
				if grid[i][k] >= tree {
					viewsRight += 1
					break
				}

				viewsRight += 1
			}

			scenicScore := viewsDown * viewsRight * viewsLeft * viewsUp
			if scenicScore > highest {
				highest = scenicScore
			}
		}
	}
	return highest
}
