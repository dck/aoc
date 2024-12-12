package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 4
	RIGHT = 8
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
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	cost := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			area, perimeter := dfs(grid, i, j, visited)
			cost += area * perimeter
		}
	}

	return cost
}

func part2(grid []string) int {
	sideMap := make([][]uint8, len(grid))
	for i := 0; i < len(sideMap); i++ {
		sideMap[i] = make([]uint8, len(grid[i]))
	}

	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			buildSideMaps(grid, visited, i, j, sideMap)
		}
	}

	visited = make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	cost := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] {
				items, corners := countCorners(grid, visited, sideMap, i, j)
				cost += items * corners
			}
		}
	}

	return cost
}

func dfs(grid []string, i, j int, visited [][]bool) (int, int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || visited[i][j] {
		return 0, 0
	}

	visited[i][j] = true
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	area := 1
	perimeter := 0
	for _, d := range directions {
		ni, nj := i+d[0], j+d[1]

		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[i]) || grid[ni][nj] != grid[i][j] {
			perimeter++
			continue
		}

		if grid[ni][nj] == grid[i][j] {
			newArea, newPerimeter := dfs(grid, ni, nj, visited)
			area += newArea
			perimeter += newPerimeter
		}
	}

	return area, perimeter
}

func buildSideMaps(grid []string, visited [][]bool, i, j int, sideMap [][]uint8) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || visited[i][j] {
		return
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	sides := [4]int{UP, DOWN, LEFT, RIGHT}
	visited[i][j] = true

	for idx, d := range directions {
		ni, nj := i+d[0], j+d[1]

		if ni < 0 {
			sideMap[i][j] |= UP
		}
		if ni >= len(grid) {
			sideMap[i][j] |= DOWN
		}
		if nj < 0 {
			sideMap[i][j] |= LEFT
		}
		if nj >= len(grid[i]) {
			sideMap[i][j] |= RIGHT
		}

		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[i]) {
			if grid[ni][nj] != grid[i][j] {
				sideMap[i][j] |= uint8(sides[idx])
			} else {
				buildSideMaps(grid, visited, ni, nj, sideMap)
			}
		}
	}
}

func countCorners(grid []string, visited [][]bool, sideMap [][]uint8, i, j int) (int, int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || visited[i][j] {
		return 0, 0
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	items := 1
	corners := Corners(sideMap, i, j)
	visited[i][j] = true

	for _, d := range directions {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[i]) {
			continue
		}

		if grid[ni][nj] == grid[i][j] {
			newItems, newCorners := countCorners(grid, visited, sideMap, ni, nj)
			items += newItems
			corners += newCorners
		}
	}

	return items, corners
}

func Corners(sideMap [][]uint8, i, j int) int {
	corners := 0
	if sideMap[i][j]&UP > 0 && sideMap[i][j]&LEFT > 0 {
		corners++
	} else if sideMap[i][j]&UP == 0 && sideMap[i][j]&LEFT == 0 && i > 0 && j > 0 && sideMap[i-1][j-1]&DOWN > 0 && sideMap[i-1][j-1]&RIGHT > 0 {
		corners++
	}

	if sideMap[i][j]&UP > 0 && sideMap[i][j]&RIGHT > 0 {
		corners++
	} else if sideMap[i][j]&UP == 0 && sideMap[i][j]&RIGHT == 0 && i > 0 && j < len(sideMap[0])-1 && sideMap[i-1][j+1]&DOWN > 0 && sideMap[i-1][j+1]&LEFT > 0 {
		corners++
	}

	if sideMap[i][j]&DOWN > 0 && sideMap[i][j]&LEFT > 0 {
		corners++
	} else if sideMap[i][j]&DOWN == 0 && sideMap[i][j]&LEFT == 0 && i < len(sideMap)-1 && j > 0 && sideMap[i+1][j-1]&UP > 0 && sideMap[i+1][j-1]&RIGHT > 0 {
		corners++
	}

	if sideMap[i][j]&DOWN > 0 && sideMap[i][j]&RIGHT > 0 {
		corners++
	} else if sideMap[i][j]&DOWN == 0 && sideMap[i][j]&RIGHT == 0 && i < len(sideMap)-1 && j < len(sideMap[0])-1 && sideMap[i+1][j+1]&UP > 0 && sideMap[i+1][j+1]&LEFT > 0 {
		corners++
	}
	return corners
}
