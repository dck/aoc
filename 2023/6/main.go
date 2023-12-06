package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var times []int
	var distances []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "Time") {
			for _, field := range strings.Fields(line)[1:] {
				value, _ := strconv.Atoi(field)
				times = append(times, value)
			}
			continue
		}
		if strings.Contains(line, "Distance") {
			for _, field := range strings.Fields(line)[1:] {
				value, _ := strconv.Atoi(field)
				distances = append(distances, value)
			}
			continue
		}
	}

	result := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]

		result *= waysToWin(time, distance)
	}

	fmt.Println("Part 1:", result)

	time, _ := strconv.Atoi(joinInts(times))
	distance, _ := strconv.Atoi(joinInts(distances))
	result = waysToWin(time, distance)
	fmt.Println("Part 2:", result)
}

func waysToWin(time int, distance int) int {
	res := 0
	for wait := 1; wait < time; wait++ {
		if wait*(time-wait) > distance {
			res += 1
		}
	}
	return res
}

func joinInts(ints []int) string {
	var res []string
	for _, i := range ints {
		res = append(res, strconv.Itoa(i))
	}
	return strings.Join(res, "")
}
