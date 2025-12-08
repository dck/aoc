package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ranges []string
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range strings.Split(line, ",") {
			if r == "" {
				continue
			}
			ranges = append(ranges, r)
		}
	}

	fmt.Println("Part 1:", part1(ranges))
	fmt.Println("Part 2:", part2(ranges))

}

func part1(ranges []string) int {
	total := 0
	for _, r := range ranges {
		parts := strings.SplitN(r, "-", 2)
		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])

		for i := low; i <= high; i++ {
			d := digits(i)
			if d%2 != 0 {
				continue
			}

			half := d / 2
			left := i / int(math.Pow10(half))
			right := i % int(math.Pow10(half))

			if left == right {
				total += i
			}
		}
	}

	return total
}

func part2(ranges []string) int {
	total := 0
	for _, r := range ranges {
		parts := strings.SplitN(r, "-", 2)
		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])

		for i := low; i <= high; i++ {
			d := digits(i)

			for step := 1; step <= d/2; step++ {
				if d%step != 0 {
					continue
				}

				matches := true
				numParts := d / step
				firstPart := i % int(math.Pow10(step))
				for p := 1; p < numParts; p++ {
					nextPart := (i / int(math.Pow10(p*step))) % int(math.Pow10(step))
					if nextPart != firstPart {
						matches = false
						break
					}
				}

				if matches {
					total += i
					break
				}
			}
		}
	}

	return total
}

func digits(x int) int {
	t := 0

	for x != 0 {
		t++
		x /= 10
	}
	return t
}
