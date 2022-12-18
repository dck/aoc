package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	name      string
	rate      int
	neighbors []string
}

type ValveMap map[int]Valve

var nodes map[int]Valve = map[int]Valve{}
var distMap [][]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		valve := parseInput(line)
		nodes[i] = valve
		i++
	}

	distMap = shortestPathsMap(nodes)

	valves := []int{}
	var start int
	for i, n := range nodes {
		if n.name == "AA" {
			start = i
		}
		valves = append(valves, i)
	}
	result := maxPressure(30, start, valves)

	fmt.Println(result)
}

func maxPressure(time int, current int, valves []int) int {
	max := 0
	for _, n := range valves {
		d := distMap[current][n]

		if d < time {
			left := time - d - 1

			newValves := make([]int, 0)
			for _, ind := range valves {
				if ind != n {
					newValves = append(newValves, ind)
				}
			}

			pressure := nodes[n].rate*left + maxPressure(left, n, newValves)

			if pressure > max {
				max = pressure
			}
		}
	}

	return max
}

func parseInput(input string) Valve {
	r := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([\w, ]+)`)

	res := r.FindStringSubmatch(input)

	name := res[1]
	rate, _ := strconv.Atoi(res[2])
	neighbors := strings.Split(strings.TrimSpace(res[3]), ", ")

	return Valve{
		rate:      rate,
		name:      name,
		neighbors: neighbors,
	}
}

func shortestPathsMap(nodes ValveMap) [][]int {
	result := make([][]int, len(nodes))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(nodes))
		for j := 0; j < len(result[i]); j++ {
			result[i][j] = math.MaxInt
		}
	}

	for n, valve := range nodes {
		result[n][n] = 0

		for m, next := range nodes {
			if contains(valve.neighbors, next.name) {
				result[n][m] = 1
			}

		}
	}

	for key := range result {
		for i := range result {
			for j := range result {

				if result[i][key] == math.MaxInt || result[key][j] == math.MaxInt {
					continue
				}
				sum := result[i][key] + result[key][j]

				if sum < result[i][j] {
					result[i][j] = sum
				}
			}
		}
	}

	return result
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
