package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("/dev/stdin")
	if err != nil {
		fmt.Println("Error reading stdin:", err)
		os.Exit(1)
	}

	input := string(data)
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(memory string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(memory, -1)

	total := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		total += a * b
	}

	return total
}

func part2(memory string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)

	matches := re.FindAllStringSubmatch(memory, -1)

	total := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			total += a * b
		}
	}

	return total
}
