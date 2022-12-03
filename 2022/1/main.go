package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	elves := [][]uint{}
	elf := []uint{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" && len(elf) > 0 {
			elves = append(elves, elf)
			elf = nil
			continue
		}

		calory, _ := strconv.ParseUint(line, 10, 64)
		elf = append(elf, uint(calory))
	}

	if len(elf) > 0 {
		elves = append(elves, elf)
	}

	total := calculate(elves, 3)
	fmt.Printf("%d\n", total)
}

func calculate(elves [][]uint, n uint) uint {
	max := make([]uint, n)

	for _, v := range elves {
		local_sum := sum(v)

		j := uint(n - 1)
		for j > 0 && local_sum > max[j-1] {
			max[j] = max[j-1]
			j--
		}
		if local_sum > max[j] {
			max[j] = local_sum
		}
	}

	return sum(max)
}

func sum(ary []uint) uint {
	result := uint(0)
	for _, v := range ary {
		result += v
	}

	return result
}
