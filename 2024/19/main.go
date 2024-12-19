package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TowelSet map[string]bool

func (ts TowelSet) Variants(design string) int {
	dp := make([]int, len(design)+1)
	dp[0] = 1

	for i := 1; i <= len(design); i++ {
		for k := range ts {
			n := len(k)

			if i >= n && design[i-n:i] == k && dp[i-n] > 0 {
				dp[i] += dp[i-n]
			}
		}
	}

	return dp[len(design)]
}

func main() {
	fh, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	towelSet := make(TowelSet)
	var designs []string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.Contains(line, ", ") {
			towels := strings.Split(line, ", ")
			for _, towel := range towels {
				towelSet[towel] = true
			}
		} else {
			designs = append(designs, line)
		}

	}

	fmt.Println("Part 1:", part1(towelSet, designs))
	fmt.Println("Part 2:", part2(towelSet, designs))
}

func part1(ts TowelSet, designs []string) int {
	res := 0
	for _, design := range designs {
		if ts.Variants(design) > 0 {
			res++
		}
	}
	return res
}

func part2(ts TowelSet, designs []string) int {
	res := 0
	for _, design := range designs {
		res += ts.Variants(design)
	}
	return res
}
