package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	ranges := [][2]uint{}
	ids := []uint{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			var start, end uint
			fmt.Sscanf(line, "%d-%d", &start, &end)
			ranges = append(ranges, [2]uint{start, end})
		} else if strings.TrimSpace(line) != "" {
			var id uint
			fmt.Sscanf(line, "%d", &id)
			ids = append(ids, id)
		}
	}
	fmt.Println("Part 1:", part1(ranges, ids))
	fmt.Println("Part 2:", part2(ranges, ids))
}

func part1(ranges [][2]uint, ids []uint) int {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	merged := [][2]uint{}
	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1][1] < r[0]-1 {
			merged = append(merged, r)
		} else {
			if merged[len(merged)-1][1] < r[1] {
				merged[len(merged)-1][1] = r[1]
			}
		}
	}

	fresh := 0
	for _, id := range ids {
		for _, m := range merged {
			if id >= m[0] && id <= m[1] {
				fresh++
				break
			}
		}
	}

	return fresh
}

func part2(ranges [][2]uint, _ids []uint) int {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	merged := [][2]uint{}
	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1][1] < r[0]-1 {
			merged = append(merged, r)
		} else {
			if merged[len(merged)-1][1] < r[1] {
				merged[len(merged)-1][1] = r[1]
			}
		}
	}

	total := 0
	for _, m := range merged {
		total += int(m[1] - m[0] + 1)
	}

	return total
}
