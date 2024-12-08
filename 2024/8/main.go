package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	i int
	j int
}

func main() {
	grid := [][]byte{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}
	fmt.Println("Part 1:", part1(grid))
	fmt.Println("Part 2:", part2(grid))
}

func part1(grid [][]byte) int {
	points := map[byte][]Point{}
	for i, line := range grid {
		for j, char := range line {
			if char != '.' {
				points[char] = append(points[char], Point{i, j})
			}
		}
	}

	maxI := len(grid)
	maxJ := len(grid[0])

	uniq := map[Point]bool{}
	for _, ps := range points {
		for i := 0; i < len(ps); i++ {
			for j := i + 1; j < len(ps); j++ {
				p1 := ps[i]
				p2 := ps[j]

				r1, r2 := ResonatingPoints(p1, p2)
				if r1.i >= 0 && r1.i < maxI && r1.j >= 0 && r1.j < maxJ {
					uniq[r1] = true
				}
				if r2.i >= 0 && r2.i < maxI && r2.j >= 0 && r2.j < maxJ {
					uniq[r2] = true
				}
			}
		}
	}

	return len(uniq)
}

func part2(grid [][]byte) int {
	points := map[byte][]Point{}
	for i, line := range grid {
		for j, char := range line {
			if char != '.' {
				points[char] = append(points[char], Point{i, j})
			}
		}
	}

	maxI := len(grid)
	maxJ := len(grid[0])
	uniq := map[Point]bool{}
	for _, ps := range points {
		for i := 0; i < len(ps); i++ {
			for j := i + 1; j < len(ps); j++ {
				p1 := ps[i]
				p2 := ps[j]

				di := p2.i - p1.i
				dj := p2.j - p1.j

				for k := 1; ; k++ {
					newP1 := Point{p1.i + k*di, p1.j + k*dj}
					newP2 := Point{p2.i - k*di, p2.j - k*dj}
					out := 0

					if newP1.i >= 0 && newP1.i < maxI && newP1.j >= 0 && newP1.j < maxJ {
						uniq[newP1] = true
					} else {
						out++
					}

					if newP2.i >= 0 && newP2.i < maxI && newP2.j >= 0 && newP2.j < maxJ {
						uniq[newP2] = true
					} else {
						out++
					}

					if out == 2 {
						break
					}
				}
			}
		}
	}

	return len(uniq)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ResonatingPoints(p1, p2 Point) (Point, Point) {
	di := p2.i - p1.i
	dj := p2.j - p1.j

	signI := 0
	signJ := 0

	if di != 0 {
		signI = di / Abs(di)
	}

	if dj != 0 {
		signJ = dj / Abs(dj)
	}

	return Point{p1.i - signI*Abs(di), p1.j - signJ*Abs(dj)}, Point{p1.i + 2*signI*Abs(di), p1.j + 2*signJ*Abs(dj)}
}

func PrintGrid(grid [][]byte) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
