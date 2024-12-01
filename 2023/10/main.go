package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func (p Point) Is(q Point) bool {
	return p.x == q.x && p.y == q.y
}

type Diff struct {
	dx, dy int
}

var Pipes = map[byte][]Diff{
	'|': {{0, -1}, {0, 1}},
	'-': {{-1, 0}, {1, 0}},
	'L': {{0, -1}, {1, 0}},
	'J': {{0, -1}, {-1, 0}},
	'7': {{0, 1}, {-1, 0}},
	'F': {{0, 1}, {1, 0}},
	'S': {{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}},
}

var Connections = map[byte][]map[byte]bool{
	'|': {{'|': true, '7': true, 'F': true, 'S': true}, {'|': true, 'L': true, 'J': true, 'S': true}},
	'-': {{'-': true, 'L': true, 'F': true, 'S': true}, {'-': true, '7': true, 'J': true, 'S': true}},
	'L': {{'|': true, '7': true, 'F': true, 'S': true}, {'-': true, '7': true, 'J': true, 'S': true}},
	'J': {{'|': true, '7': true, 'F': true, 'S': true}, {'-': true, 'L': true, 'F': true, 'S': true}},
	'7': {{'|': true, 'L': true, 'J': true, 'S': true}, {'-': true, 'L': true, 'F': true, 'S': true}},
	'F': {{'|': true, 'L': true, 'J': true, 'S': true}, {'-': true, '7': true, 'J': true, 'S': true}},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := make([]string, 0)
	var start Point
	i := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		field = append(field, line)

		for j, c := range line {
			if c == 'S' {
				start = Point{j, i}
			}
		}
		i++
	}

	visited := drawCycle(field, start)

	cycleLength := countValues(visited, 1)
	fmt.Println("Part 1:", (cycleLength+1)/2)

	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if i == 0 || i == len(visited)-1 || j == 0 || j == len(visited[i])-1 {
				flood(&visited, Point{j, i})
			}
		}
	}

	print2DIntSlice(visited)
	fmt.Println("Part 2:", countValues(visited, 0))

}

func drawCycle(field []string, start Point) [][]int {
	visited := make([][]int, len(field))

	for i := range visited {
		visited[i] = make([]int, len(field[i]))
	}

	var dfs func(Point)
	dfs = func(p Point) {
		if visited[p.y][p.x] > 0 {
			return
		}

		visited[p.y][p.x]++

		current := field[p.y][p.x]

		for i, d := range Pipes[current] {
			x, y := p.x+d.dx, p.y+d.dy

			if y < 0 || y >= len(field) || x < 0 || x >= len(field[y]) {
				continue
			}

			if field[y][x] != '.' && (current == 'S' || Connections[current][i][field[y][x]]) {
				dfs(Point{x, y})
			}
		}
	}

	dfs(start)
	return visited
}

func print2DIntSlice(slice [][]int) {
	for _, row := range slice {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}

func countValues(schema [][]int, value int) int {
	total := 0
	for i := range schema {
		for j := range schema[i] {
			if schema[i][j] == value {
				total++
			}
		}
	}

	return total
}

func flood(schema *[][]int, point Point) {
	s := *schema

	if s[point.y][point.x] > 0 {
		return
	}

	(*schema)[point.y][point.x] = 2

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			y, x := point.y+i, point.x+j

			if y < 0 || y >= len(s) || x < 0 || x >= len(s[y]) {
				continue
			}

			flood(schema, Point{x, y})
		}
	}
}
