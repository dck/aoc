package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var num1, num2 int

		_, err := fmt.Sscanf(line, "%d %d", &num1, &num2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		left = append(left, num1)
		right = append(right, num2)
	}

	fmt.Println("Part 1:", part1(left, right))
	fmt.Println("Part 2:", part2(left, right))
}

func part1(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i := 0; i < len(left); i++ {
		total += Abs(left[i] - right[i])
	}

	return total
}

func part2(left, right []int) int {
	freq := make(map[int]int)

	for _, n := range right {
		freq[n]++
	}

	total := 0
	for _, n := range left {
		total += freq[n] * n
	}

	return total
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
