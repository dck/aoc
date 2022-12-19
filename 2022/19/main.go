package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	ORE      = "ore"
	CLAY     = "clay"
	OBSIDIAN = "obsidian"
	GEODE    = "geode"
)

var Priorities []string = []string{GEODE, OBSIDIAN, CLAY, ORE}

type Resource map[string]int
type Blueprint struct {
	index      int
	Generation int
	Robots     []string
	Resources  Resource
	Costs      map[string]Resource
}

func CreateBlueprint(i int) *Blueprint {
	b := &Blueprint{
		index:      i,
		Generation: 1,
		Robots:     make([]string, 0),
		Resources:  map[string]int{ORE: 0, CLAY: 0, OBSIDIAN: 0, GEODE: 0},
		Costs:      make(map[string]Resource),
	}
	return b
}

func Evolute(b *Blueprint, steps int) int {
	maxGeodes := 0
	queue := []*Blueprint{b}

	for len(queue) > 0 {
		blueprint := queue[0]
		queue = queue[1:]

		maxGeodes = Max(maxGeodes, blueprint.Resources[GEODE])

		if blueprint.Generation == steps {
			continue
		}

		possibleRobots := []string{}
		for _, t := range Priorities {
			if b.CanAddRobot(t) {
				possibleRobots = append(possibleRobots, t)
			}
		}

		for _, t := range possibleRobots {
			newBlueprint := blueprint

			newBlueprint.GatherResources()
			newBlueprint.AddRobot(t)
			newBlueprint.Generation += 1
			queue = append(queue, newBlueprint)
		}

		newBlueprint := blueprint
		newBlueprint.GatherResources()
		newBlueprint.Generation += 1
		queue = append(queue, newBlueprint)
	}

	return maxGeodes
}

func (b *Blueprint) GatherResources() {
	for _, r := range b.Robots {
		b.Resources[r] += 1
	}
}

func (b *Blueprint) CanAddRobot(t string) bool {
	cost := b.Costs[t]
	for unit, price := range cost {
		left := b.Resources[unit]
		if left < price {
			return false
		}
	}

	return true
}

func (b *Blueprint) AddRobot(t string) error {
	cost := b.Costs[t]
	for unit, price := range cost {
		if b.Resources[unit] < price {
			panic("You can't build such robot")
		}
		b.Resources[unit] -= price
	}

	b.Robots = append(b.Robots, t)
	return nil
}

func (b Blueprint) AddRobotForFree(t string) {
	b.Robots = append(b.Robots, t)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		blueprint := parseLine(line)

		blueprint.AddRobotForFree(ORE)

		m := Evolute(blueprint, 24)
		fmt.Println(m)
	}
}

func parseLine(line string) *Blueprint {
	parts := strings.Split(line, ".")

	numStr := strings.Fields(parts[0])[1]
	index, _ := strconv.Atoi(numStr[:len(numStr)-1])

	blueprint := CreateBlueprint(index)

	robotRegexp := regexp.MustCompile(`Each (\w+) robot`)
	r := regexp.MustCompile(`(\d+) (\w+)`)
	for i := 0; i < len(parts)-1; i++ {
		robotStatement := parts[i]

		m := robotRegexp.FindStringSubmatch(robotStatement)
		t := m[1]

		cost := make(map[string]int)
		matches := r.FindAllString(robotStatement, -1)
		for i := 0; i < len(matches); i++ {
			costs := strings.Fields(matches[i])
			unit := costs[1]
			price, _ := strconv.Atoi(costs[0])

			cost[unit] = price
		}

		blueprint.Costs[t] = cost
	}

	return blueprint
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
