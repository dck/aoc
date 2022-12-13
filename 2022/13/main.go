package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	items  []Item
	number int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	index := 0
	pairs := [2]string{}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if calculate(pairs[0], pairs[1]) {
				total += index / 2
				fmt.Println(index / 2)
			}
			continue
		}

		pairs[index%2] = line
		index += 1
	}

	if calculate(pairs[0], pairs[1]) {
		total += index / 2
	}

	fmt.Println(total)
}

func calculate(left string, right string) bool {
	leftPart := parse(left)
	rightPart := parse(left)

	return compare(leftPart, rightPart) <= 0
}

func compare(left any, right any) int {
	a, okA := left.(float64)
	b, okB := right.(float64)
	if okA && okB {
		return int(a) - int(b)
	}

	var leftList []any
	var rightList []any

	switch left.(type) {
	case []any, []float64:
		leftList = left.([]any)
	case float64:
		leftList = []any{left}
	}

	switch right.(type) {
	case []any, []float64:
		rightList = right.([]any)
	case float64:
		rightList = []any{right}
	}

	if len(leftList) > len(rightList) {
		return 1
	}

	for i := range leftList {
		if r := compare(leftList[i], rightList[i]); r != 0 {
			return r
		}
	}

	if len(rightList) == len(leftList) {
		return 0
	}
	return -1

}

func parse(line string) any {
	var res any
	if err := json.Unmarshal([]byte(line), &res); err != nil {
		panic(err)
	}

	return res
}
