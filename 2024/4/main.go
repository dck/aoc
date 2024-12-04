package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var directions = [][]int{
	{0, 1}, {1, 0}, {1, 1}, {1, -1},
	{0, -1}, {-1, 0}, {-1, -1}, {-1, 1},
}

func main() {
	field := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, line)
	}

	fmt.Println("Part 1:", part1(field))
	fmt.Println("Part 2:", part2(field))
}

func part1(field []string) int {
	word := "XMAS"
	count := 0
	wordLen := len(word)

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == word[0] {
				for _, dir := range directions {
					if checkWord(field, word, i, j, dir, wordLen) {
						count++
					}
				}
			}
		}
	}

	return count
}

func part2(field []string) int {
	count := 0
	for i := 1; i < len(field)-1; i++ {
		for j := 1; j < len(field[0])-1; j++ {
			if field[i][j] == 'A' && checkPattern(field, i, j) {
				count++
			}
		}
	}

	return count
}

func checkWord(field []string, word string, x, y int, dir []int, wordLen int) bool {
	for k := 0; k < wordLen; k++ {
		ni := x + k*dir[0]
		nj := y + k*dir[1]
		if ni < 0 || ni >= len(field) || nj < 0 || nj >= len(field[0]) || field[ni][nj] != word[k] {
			return false
		}
	}
	return true
}

func checkPattern(field []string, i, j int) bool {
	diag1 := []byte{field[i-1][j-1], field[i][j], field[i+1][j+1]}
	diag2 := []byte{field[i-1][j+1], field[i][j], field[i+1][j-1]}

	sort.Slice(diag1, func(i, j int) bool { return diag1[i] < diag1[j] })
	sort.Slice(diag2, func(i, j int) bool { return diag2[i] < diag2[j] })

	return string(diag1) == "AMS" && string(diag2) == "AMS"
}
