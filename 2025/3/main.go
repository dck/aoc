package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	batteries := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		batteries = append(batteries, line)
	}
	fmt.Println("Part 1:", part1(batteries))
	fmt.Println("Part 2:", part2(batteries))
}

func part1(batteries []string) int {
	total := 0
	for i := 0; i < len(batteries); i++ {
		j := joltage(batteries[i], 2)
		total += j
	}
	return total
}

func part2(batteries []string) int {
	total := 0
	for i := 0; i < len(batteries); i++ {
		j := joltage(batteries[i], 12)
		total += j
	}
	return total
}

func joltage(b string, num int) int {
	stack := []int{}

	k := len(b) - num
	for i := 0; i < len(b); i++ {
		d := int(b[i] - '0')

		for k > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, d)
	}

	total := 0
	for i := 0; i < num; i++ {
		digit := stack[i]
		total = total*10 + digit
	}

	return total
}
