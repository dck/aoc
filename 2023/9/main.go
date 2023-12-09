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

	total1 := 0
	total2 := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		numbers := convert(strings.Fields(line))
		total1 += calculateNext(numbers)
		total2 += calculatePrev(numbers)
	}

	fmt.Println("Part 1: ", total1)
	fmt.Println("Part 2: ", total2)
}

func convert(fields []string) []int {
	res := make([]int, len(fields))
	for i, v := range fields {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func calculateNext(ary []int) int {
	res := make([]int, len(ary)-1)

	allZeros := true
	for i := 0; i < len(ary)-1; i++ {
		res[i] = ary[i+1] - ary[i]
		if res[i] != 0 {
			allZeros = false
		}
	}

	if allZeros {
		return ary[len(ary)-1]
	} else {
		last := calculateNext(res)
		return last + ary[len(ary)-1]
	}
}

func calculatePrev(ary []int) int {
	res := make([]int, len(ary)-1)

	allZeros := true
	for i := 0; i < len(ary)-1; i++ {
		res[i] = ary[i+1] - ary[i]
		if res[i] != 0 {
			allZeros = false
		}
	}

	if allZeros {
		return ary[0]
	} else {
		first := calculatePrev(res)
		return ary[0] - first
	}
}
