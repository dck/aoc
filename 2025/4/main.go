package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '@' {
				continue
			}

			c := 0
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					if i+x < 0 || i+x >= len(grid) || j+y < 0 || j+y >= len(grid[i]) {
						continue
					}

					if grid[i+x][j+y] == '@' {
						c++
					}
				}
			}
			if c < 4 {
				total++
			}
		}
	}
	return total
}

func part2(grid []string) int {
	gridChars := make([][]rune, len(grid))
	for i := range grid {
		gridChars[i] = []rune(grid[i])
	}

	total := 0
	removed := true
	for removed {
		removed = false
		for i := range gridChars {
			for j := range gridChars[i] {
				if gridChars[i][j] != '@' {
					continue
				}

				c := 0
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue
						}
						ni := i + x
						nj := j + y
						if ni < 0 || ni >= len(gridChars) || nj < 0 || nj >= len(gridChars[ni]) {
							continue
						}

						if gridChars[ni][nj] == '@' {
							c++
						}
					}
				}
				if c < 4 {
					total++
					gridChars[i][j] = 'x'
					removed = true
				}
			}
		}
	}
	return total
}
