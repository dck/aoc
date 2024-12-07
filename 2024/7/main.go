package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	result int
	nums   []int
}

func main() {
	equations := []Equation{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		result, _ := strconv.Atoi(parts[0])
		numsStr := strings.Split(parts[1], " ")
		nums := make([]int, len(numsStr))
		for _, nStr := range numsStr {
			n, _ := strconv.Atoi(nStr)
			nums = append(nums, n)
		}

		equations = append(equations, Equation{result, nums})
	}
	fmt.Println("Part 1:", part1(equations))
	fmt.Println("Part 2:", part2(equations))
}

func part1(equations []Equation) int {
	total := 0
	for _, eq := range equations {
		reachableSums := map[int]bool{}
		findSums(eq.result, eq.nums, 0, reachableSums, false)
		if reachableSums[eq.result] {
			total += eq.result
		}
	}
	return total
}

func part2(equations []Equation) int {
	total := 0
	for _, eq := range equations {
		reachableSums := map[int]bool{}
		findSums(eq.result, eq.nums, 0, reachableSums, true)
		if reachableSums[eq.result] {
			total += eq.result
		}
	}
	return total
}

func findSums(target int, numbers []int, currentSum int, reachableSums map[int]bool, withConcat bool) {
	if currentSum > target {
		return
	}

	if len(numbers) == 0 {
		reachableSums[currentSum] = true
		return
	}

	if currentSum == 0 {
		findSums(target, numbers[1:], numbers[0], reachableSums, withConcat)
	} else {
		product := currentSum * numbers[0]
		sum := currentSum + numbers[0]
		findSums(target, numbers[1:], product, reachableSums, withConcat)
		findSums(target, numbers[1:], sum, reachableSums, withConcat)

		if withConcat {
			concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentSum, numbers[0]))
			findSums(target, numbers[1:], concatenated, reachableSums, true)
		}
	}
}
