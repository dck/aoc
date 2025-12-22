package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := []string{}
	operators := []rune{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "+") || strings.Contains(line, "*") {
			for _, char := range strings.Fields(line) {
				operators = append(operators, []rune(char)[0])
			}
		} else {
			lines = append(lines, line)
		}

	}
	fmt.Println("Part 1:", part1(lines, operators))
	fmt.Println("Part 2:", part2(lines, operators))
}

func part1(lines []string, operators []rune) int {
	numbers := [][]int{}

	for _, line := range lines {
		fields := strings.Fields(line)
		row := []int{}
		for _, field := range fields {
			var num int
			fmt.Sscanf(field, "%d", &num)
			row = append(row, num)
		}
		numbers = append(numbers, row)
	}

	total := 0

	for i, op := range operators {
		switch op {
		case '+':
			localSum := 0
			for j := 0; j < len(numbers); j++ {
				localSum += numbers[j][i]
			}
			total += localSum
		case '*':
			localProd := 1
			for j := 0; j < len(numbers); j++ {
				localProd *= numbers[j][i]
			}
			total += localProd
		}
	}

	return total
}

func part2(lines []string, operators []rune) int {
	total := 0
	maxLen := 0
	for _, line := range lines {
		l := len(line)
		if l > maxLen {
			maxLen = l
		}
	}

	opIndex := len(operators) - 1
	localTotal := 0
	for pos := maxLen - 1; pos >= 0; pos-- {
		localNumber := 0
		for _, line := range lines {
			if pos >= len(line) {
				continue
			}
			char := line[pos]
			if char == ' ' {
				continue
			}
			localNumber = 10*localNumber + int(char-'0')
		}
		if localNumber == 0 {
			opIndex--
			total += localTotal
			localTotal = 0
			continue
		}

		if operators[opIndex] == '+' {
			localTotal += localNumber
		} else {
			if localTotal == 0 {
				localTotal = 1
			}
			localTotal *= localNumber
		}

		if pos == 0 {
			total += localTotal
		}
	}

	return total
}
