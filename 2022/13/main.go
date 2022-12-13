package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	allSignals := []any{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if calculate(pairs[0], pairs[1]) {
				total += index / 2
			}
			continue
		}

		pairs[index%2] = line
		allSignals = append(allSignals, parse(line))
		index += 1
	}

	if calculate(pairs[0], pairs[1]) {
		total += index / 2
	}
	fmt.Println(total)

	allSignals = append(allSignals, parse("[[2]]"))
	allSignals = append(allSignals, parse("[[6]]"))

	sort.Slice(allSignals, func(i, j int) bool {
		return compare(allSignals[i], allSignals[j]) < 0
	})

	key := 1
	for i, signal := range allSignals {
		signalStr, _ := json.Marshal(signal)
		if string(signalStr) == "[[2]]" || string(signalStr) == "[[6]]" {
			key *= i + 1
		}
	}
	fmt.Println(key)
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

	for i := range leftList {
		if i >= len(rightList) {
			return 1
		}
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
