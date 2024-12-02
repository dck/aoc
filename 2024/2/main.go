package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := [][]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var levels []int
		numStrs := strings.Fields(line)
		for _, numStr := range numStrs {
			n, _ := strconv.Atoi(numStr)
			levels = append(levels, n)
		}

		reports = append(reports, levels)
	}

	fmt.Println("Part 1:", part1(reports))
	fmt.Println("Part 2:", part2(reports))
}

func part1(reports [][]int) int {
	total := 0
	for _, levels := range reports {
		if isIncreasing(levels) == -1 || isDecreasing(levels) == -1 {
			total++
		}
	}
	return total
}

func part2(reports [][]int) int {
	total := 0

	for _, levels := range reports {
		inIdx := isIncreasing(levels)
		deIdx := isDecreasing(levels)

		if inIdx == -1 || deIdx == -1 {
			total++
		} else {
			for i := 0; i < len(levels); i++ {
				newLevels := make([]int, 0, len(levels)-1)
				newLevels = append(newLevels, levels[:i]...)
				newLevels = append(newLevels, levels[i+1:]...)

				if isIncreasing(newLevels) == -1 || isDecreasing(newLevels) == -1 {
					total++
					break
				}
			}
		}
	}

	return total
}

func isIncreasing(levels []int) int {
	for i := 0; i < len(levels)-1; i++ {
		if levels[i] >= levels[i+1] || levels[i+1]-levels[i] > 3 {
			return i
		}
	}
	return -1
}

func isDecreasing(levels []int) int {
	for i := 0; i < len(levels)-1; i++ {
		if levels[i] <= levels[i+1] || levels[i]-levels[i+1] > 3 {
			return i
		}
	}
	return -1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
