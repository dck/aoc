package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rotations := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var dir rune
		var num int
		_, _ = fmt.Sscanf(line, "%c%d", &dir, &num)
		if dir == 'L' {
			rotations = append(rotations, -num)
		} else {
			rotations = append(rotations, num)
		}
	}

	fmt.Println("Part 1:", part1(rotations))
	fmt.Println("Part 2:", part2(rotations))
}

func part1(rotations []int) int {
	total := 50
	zeros := 0
	for _, r := range rotations {
		total += r

		if total < 0 {
			total += 100
		}
		total %= 100

		if total == 0 {
			zeros++
		}
	}
	return zeros
}

func part2(rotations []int) int {
	crosses := 0
	total := 50

	for _, r := range rotations {
		distance := Abs(r)

		crosses += distance / 100
		distance %= 100

		if r > 0 {
			if total != 0 && total+distance > 100 {
				crosses++
			}
			total = (total + distance) % 100
		} else {
			if total != 0 && total-distance < 0 {
				crosses++
			}
			total = (total - distance + 100) % 100
		}

		if total == 0 {
			crosses++
		}
	}
	return crosses
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
