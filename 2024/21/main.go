package main

import (
	"bufio"
	"fmt"
	"iter"
	"math"
	"os"
	"strings"
)

type Point [2]int
type Keypad [][]byte

func (k *Keypad) Coordinates(val byte) Point {
	for i, row := range *k {
		for j, key := range row {
			if key == val {
				return Point{i, j}
			}
		}
	}
	return Point{-1, -1}
}

type Path struct {
	Node     Point
	Steps    string
	Switches int
}

func (k *Keypad) Paths(start Point, target Point) []string {
	r1, c1 := start[0], start[1]
	r2, c2 := target[0], target[1]
	gap := k.Coordinates(' ')
	dr, dc := r2-r1, c2-c1

	rowMoves := strings.Repeat("v", abs(dr))
	if dr < 0 {
		rowMoves = strings.Repeat("^", abs(dr))
	}

	colMoves := strings.Repeat(">", abs(dc))
	if dc < 0 {
		colMoves = strings.Repeat("<", abs(dc))
	}

	if dr == 0 && dc == 0 {
		return []string{""}
	}

	if dr == 0 {
		return []string{colMoves}
	}

	if dc == 0 {
		return []string{rowMoves}
	}

	if r1 == gap[0] && c2 == gap[1] {
		return []string{rowMoves + colMoves}
	}

	if r2 == gap[0] && c1 == gap[1] {
		return []string{colMoves + rowMoves}
	}

	return []string{rowMoves + colMoves, colMoves + rowMoves}
}

var numericKeypad Keypad = Keypad{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
}

var directionalKeypad Keypad = Keypad{
	{' ', '^', 'A'},
	{'<', 'v', '>'},
}

func main() {
	codes := []string{}
	fh, _ := os.Open("input.txt")
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, strings.TrimSpace(line))
	}

	fmt.Println("Part 1:", part1(codes))
	fmt.Println("Part 2:", part2(codes))
}

func part1(codes []string) uint64 {
	res := uint64(0)

	for _, code := range codes {
		cache := make(map[CacheEntry]uint64)
		moves := Movements(code, 3, cache)

		var numericPart uint64
		fmt.Sscanf(code, "%dA", &numericPart)
		res += moves * numericPart
	}

	return res
}

func part2(codes []string) uint64 {
	res := uint64(0)

	for _, code := range codes {
		cache := make(map[CacheEntry]uint64)
		moves := Movements(code, 26, cache)

		var numericPart uint64
		fmt.Sscanf(code, "%dA", &numericPart)
		res += moves * numericPart
	}
	return res
}

type CacheEntry struct {
	Depth   int
	Command string
}

func Movements(command string, depth int, cache map[CacheEntry]uint64) uint64 {
	if depth == 0 {
		return uint64(len(command))
	}

	if cached, ok := cache[CacheEntry{depth, command}]; ok {
		return cached
	}

	keypad := numericKeypad
	if strings.ContainsAny(command, "<>^v") {
		keypad = directionalKeypad
	}

	res := uint64(0)
	for _, shortestPaths := range buildPaths(command, keypad) {
		var shortestMovement uint64 = math.MaxUint64
		for _, sp := range shortestPaths {
			possibleMovement := Movements(sp, depth-1, cache)
			if possibleMovement < shortestMovement {
				shortestMovement = possibleMovement
			}
		}
		res += shortestMovement
	}

	cache[CacheEntry{depth, command}] = res
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func buildPaths(commands string, keypad Keypad) [][]string {
	res := [][]string{}

	for key1, key2 := range Zip([]byte("A"+commands), []byte(commands)) {
		start := keypad.Coordinates(key1)
		target := keypad.Coordinates(key2)
		paths := keypad.Paths(start, target)
		for i := 0; i < len(paths); i++ {
			paths[i] = paths[i] + "A"
		}
		res = append(res, paths)
	}

	return res
}

func Zip[T any](a []T, b []T) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		for i := 0; i < len(a) && i < len(b); i++ {
			if !yield(a[i], b[i]) {
				break
			}
		}
	}
}
