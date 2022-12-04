package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		total += calculate(line)
	}

	fmt.Println(total)
}

func calculate(line string) int {
	parts := strings.Split(line, ",")
	a1, b1 := parseRange(parts[0])
	a2, b2 := parseRange(parts[1])

	// if a1 <= a2 && b1 >= b2 {
	// 	return 1
	// }

	// if a2 <= a1 && b2 >= b1 {
	// 	return 1
	// }

	// return 0

	if b1 < a2 || b2 < a1 {
		return 0
	}
	return 1
}

func parseRange(r string) (int, int) {
	parts := strings.Split(r, "-")
	a, _ := strconv.ParseInt(parts[0], 10, 64)
	b, _ := strconv.ParseInt(parts[1], 10, 64)

	return int(a), int(b)
}
