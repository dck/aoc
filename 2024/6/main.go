package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Direction struct {
	sym byte
	di  int
	dj  int
}

type Position struct {
	i int
	j int
}

var directions = []Direction{
	{sym: '>', di: 0, dj: 1},
	{sym: 'v', di: 1, dj: 0},
	{sym: '<', di: 0, dj: -1},
	{sym: '^', di: -1, dj: 0},
}

func main() {
	grid := [][]byte{}
	copygrid := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
		copygrid = append(copygrid, []byte(line))
	}

	fmt.Println("Part 1:", part1(grid))
	fmt.Println("Part 2:", part2(copygrid))
}

func part1(grid [][]byte) int {
	directionIdx, currPosition := getStartData(grid)
	currDirection := directions[directionIdx]

	for {
		grid[currPosition.i][currPosition.j] = 'X'
		ni := currPosition.i + currDirection.di
		nj := currPosition.j + currDirection.dj

		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
			break
		}

		if grid[ni][nj] == '#' {
			directionIdx = (directionIdx + 1) % len(directions)
			currDirection = directions[directionIdx]
			continue
		}

		currPosition.i = ni
		currPosition.j = nj
	}

	return countSyms(grid, 'X')
}

func part2(grid [][]byte) int {
	cycles := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' {
				grid[i][j] = '#'
				if isCycle(grid) {
					cycles++
				}
				grid[i][j] = '.'
			}
		}
	}

	return cycles
}

func isCycle(grid [][]byte) bool {
	directionIdx, currPosition := getStartData(grid)
	currDirection := directions[directionIdx]

	maxSteps := len(grid) * len(grid[0])
	steps := 0
	for {
		ni := currPosition.i + currDirection.di
		nj := currPosition.j + currDirection.dj
		steps++

		if steps > maxSteps {
			return true
		}

		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
			break
		}

		if grid[ni][nj] == '#' {
			directionIdx = (directionIdx + 1) % len(directions)
			currDirection = directions[directionIdx]
			continue
		}

		currPosition.i = ni
		currPosition.j = nj
	}

	return false
}

func print(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(string(grid[i]))
	}
	fmt.Println()
}

func countSyms(grid [][]byte, sym byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == sym {
				count++
			}
		}
	}
	return count
}

func getStartData(grid [][]byte) (directionIdx int, pos Position) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '#' && grid[i][j] != '.' {
				directionIdx = slices.IndexFunc(directions, func(d Direction) bool {
					return d.sym == grid[i][j]
				})

				pos = Position{i, j}
				break
			}
		}
	}
	return
}
