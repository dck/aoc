package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.ParseInt(os.Args[1], 10, 64)
	buffer := make([]string, n)
	i := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		buffer[i] = line
		i++

		if i == n {
			i = 0
			total += calculate(buffer)
			buffer = make([]string, n)
		}
	}

	fmt.Println(total)
}

func calculate(elves []string) int {
	result := []byte(elves[0])

	for i := 1; i < len(elves); i++ {
		result = intersect([]byte(result), []byte(elves[i]))
	}

	return value(result[0])
}

func value(char byte) int {
	value := int(char)

	if value >= int('a') {
		return value - int('a') + 1

	} else {
		return value - int('A') + 27
	}
}

func intersect[T comparable](a []T, b []T) []T {
	res := make([]T, 0)
	h := make(map[T]bool)

	for _, v := range a {
		h[v] = true
	}

	for _, v := range b {
		if _, ok := h[v]; ok {
			res = append(res, v)
		}
	}

	return res
}
