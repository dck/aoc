package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rules := [][2]int{}
	pages := [][]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			nums := strings.Split(line, "|")
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])

			rules = append(rules, [2]int{a, b})
		}

		if strings.Contains(line, ",") {
			nums := strings.Split(line, ",")
			page := []int{}
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				page = append(page, n)
			}
			pages = append(pages, page)
		}
	}

	fmt.Println("Part 1:", part1(rules, pages))
	fmt.Println("Part 2:", part2(rules, pages))
}

func part1(rules [][2]int, pages [][]int) int {
	total := 0
	for _, page := range pages {
		if checkOrder(rules, page) {
			total += page[len(page)/2]
		}
	}

	return total
}

func part2(rules [][2]int, pages [][]int) int {
	total := 0

	set := map[int]map[int]bool{}
	for _, rule := range rules {
		if _, ok := set[rule[0]]; !ok {
			set[rule[0]] = map[int]bool{}
		}
		set[rule[0]][rule[1]] = true
	}

	for _, page := range pages {
		miniMatrix := map[int][]int{}
		for i := 0; i < len(page); i++ {
			for j := 0; j < len(page); j++ {
				if _, ok := set[page[i]]; ok {
					if _, ok := set[page[i]][page[j]]; ok {
						miniMatrix[page[i]] = append(miniMatrix[page[i]], page[j])
					}
				}
			}
		}

		sorted := TopSort(miniMatrix)

		revIndex := map[int]int{}
		for i, num := range sorted {
			revIndex[num] = i
		}

		copyPage := make([]int, len(page))
		copy(copyPage, page)

		sort.Slice(copyPage, func(i, j int) bool {
			return revIndex[copyPage[i]] < revIndex[copyPage[j]]
		})

		if !slices.Equal(page, copyPage) {
			total += copyPage[len(copyPage)/2]
		}
	}

	return total
}

func TopSort(matrix map[int][]int) []int {
	inbound := map[int]int{}
	for _, v := range matrix {
		for _, n := range v {
			inbound[n]++
		}
	}

	queue := []int{}
	for k := range matrix {
		if inbound[k] == 0 {
			queue = append(queue, k)
		}
	}

	sorted := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		sorted = append(sorted, node)
		for _, n := range matrix[node] {
			inbound[n]--
			if inbound[n] == 0 {
				queue = append(queue, n)
			}
		}
	}

	return sorted
}

func checkOrder(rules [][2]int, page []int) bool {
	for _, rule := range rules {
		aIndex, bIndex := -1, -1
		for i, num := range page {
			if num == rule[0] {
				aIndex = i
			}
			if num == rule[1] {
				bIndex = i
			}
		}
		if aIndex != -1 && bIndex != -1 && aIndex > bIndex {
			return false
		}
	}
	return true
}
