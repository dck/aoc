package main

import (
	"fmt"
	"math"
	"os"
)

type Robot struct {
	x, y   int
	vx, vy int
}

func (r *Robot) move(seconds int) {
	r.x += r.vx * seconds
	r.y += r.vy * seconds
}

type Field struct {
	robots        []Robot
	width, height int
}

func (f *Field) move(seconds int) {
	for i := range f.robots {
		f.robots[i].move(seconds)

		f.robots[i].x = (f.robots[i].x%f.width + f.width) % f.width
		f.robots[i].y = (f.robots[i].y%f.height + f.height) % f.height
	}
}

func (f *Field) safetyFactor() int {
	horizontal := f.height / 2
	vertical := f.width / 2

	var quadrants [4]int // 0: top-left, 1: top-right, 2: bottom-left, 3: bottom-right

	for _, r := range f.robots {
		if r.x < vertical && r.y < horizontal {
			quadrants[0]++
		} else if r.x > vertical && r.y < horizontal {
			quadrants[1]++
		} else if r.x < vertical && r.y > horizontal {
			quadrants[2]++
		} else if r.x > vertical && r.y > horizontal {
			quadrants[3]++
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func (f *Field) print(printBorders bool) {
	grid := make([][]byte, f.height)
	for i := range grid {
		grid[i] = make([]byte, f.width)
	}

	for _, r := range f.robots {
		grid[r.y][r.x]++
	}

	for i := range grid {
		for j := range grid[i] {
			if printBorders && i == (len(grid))/2 {
				fmt.Print("-")
				continue
			}

			if printBorders && j == (len(grid[i]))/2 {
				fmt.Print("|")
				continue
			}

			if grid[i][j] > 0 {
				fmt.Print(string('0' + grid[i][j]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var original []Robot

	for {
		var x, y, vx, vy int
		if _, err := fmt.Fscanf(os.Stdin, "p=%d,%d v=%d,%d\n", &x, &y, &vx, &vy); err != nil {
			break
		}
		original = append(original, Robot{x, y, vx, vy})
	}

	robotsCopy := make([]Robot, len(original))
	copy(robotsCopy, original)
	field := Field{robotsCopy, 101, 103}
	fmt.Println("Part 1:", part1(field))

	robotsCopy = make([]Robot, len(original))
	copy(robotsCopy, original)
	field = Field{robotsCopy, 101, 103}
	fmt.Println("Part 2:", part2(field))
}

func part1(field Field) int {
	field.move(100)
	return field.safetyFactor()
}

func part2(field Field) int {
	min := math.MaxInt
	steps := 0
	for i := 1; i < 100000; i++ {
		field.move(1)
		sf := field.safetyFactor()

		if sf < min {
			min = sf
			steps = i
		}
	}

	return steps
}
