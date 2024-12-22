package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}

	fmt.Println("Part 1:", part1(numbers))
	fmt.Println("Part 2:", part2(numbers))
}

func part1(numbers []int) uint64 {
	res := uint64(0)
	for _, n := range numbers {
		for range 2000 {
			n = next(n)
		}
		res += uint64(n)
	}
	return res
}

func part2(numbers []int) int {
	global := make(map[[4]int]int)
	for _, n := range numbers {
		localSequences := make(map[[4]int]int)

		a, b, c, d := 0, 0, 0, n
		for i := 0; i < 2000; i++ {
			nextNum := next(n)
			delta := nextNum%10 - n%10
			n = nextNum
			a, b, c, d = b, c, d, delta

			if i < 3 {
				continue
			}

			key := [4]int{a, b, c, d}
			if _, ok := localSequences[key]; !ok {
				localSequences[key] = n % 10
			}
		}

		for key, localValue := range localSequences {
			global[key] += localValue
		}
	}
	max := 0
	for v := range maps.Values(global) {
		max = Max(max, v)
	}

	return max
}

func next(n int) int {
	n = ((n * 64) ^ n) % 16777216
	n = ((n / 32) ^ n) % 16777216
	n = ((n * 2048) ^ n) % 16777216
	return n
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
