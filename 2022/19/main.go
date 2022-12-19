package main

import (
	"bufio"
	"fmt"
	"os"
)

type Blueprint struct {
	index        int
	OreCost      int
	ClayCost     int
	ObsidianCost [2]int
	GeodeCost    [2]int
}

type State struct {
	Generation     int
	GeodeRobots    int
	Geode          int
	ObsidianRobots int
	Obsidian       int
	ClayRobots     int
	Clay           int
	OreRobots      int
	Ore            int
}

func CreateState() State {
	s := State{
		Generation: 0,
		OreRobots:  1,
	}
	return s
}

func Evolute(b *Blueprint, steps int) int {
	maxGeodes := 0
	queue := []State{CreateState()}

	seen := map[State]bool{}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		fmt.Printf(
			"Generation: %d. Robots: Ore: %d, Clay: %d, Obsidian: %d, Geode: %d\n",
			state.Generation,
			state.OreRobots,
			state.ClayRobots,
			state.ObsidianRobots,
			state.GeodeRobots,
		)

		maxGeodes = Max(maxGeodes, state.Geode)

		if state.Generation == steps {
			continue
		}

		if seen[state] {
			continue
		}

		seen[state] = true

		if state.Ore >= b.GeodeCost[0] && state.Obsidian >= b.GeodeCost[1] {
			newState := state
			newState.GatherResources()
			newState.Ore -= b.GeodeCost[0]
			newState.Obsidian -= b.GeodeCost[1]
			newState.GeodeRobots += 1
			queue = append(queue, newState)
		}

		if state.Ore >= b.ObsidianCost[0] && state.Clay >= b.ObsidianCost[1] {
			newState := state
			newState.GatherResources()
			newState.Ore -= b.ObsidianCost[0]
			newState.Clay -= b.ObsidianCost[1]
			newState.ObsidianRobots += 1
			queue = append(queue, newState)
		}

		if state.Ore >= b.ClayCost {
			newState := state
			newState.GatherResources()
			newState.Ore -= b.ClayCost
			newState.ClayRobots += 1
			queue = append(queue, newState)
		}

		if state.Ore >= b.OreCost {
			newState := state
			newState.GatherResources()
			newState.Ore -= b.OreCost
			newState.OreRobots += 1
			queue = append(queue, newState)
		}

		newState := state
		newState.GatherResources()
		queue = append(queue, newState)
	}

	return maxGeodes
}

func (b *State) GatherResources() {
	b.Geode += b.GeodeRobots
	b.Obsidian += b.ObsidianRobots
	b.Clay += b.ClayRobots
	b.Ore += b.OreRobots
	b.Generation += 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		blueprint := parseLine(line)

		m := Evolute(blueprint, 24)
		fmt.Println(m)
	}
}

func parseLine(line string) *Blueprint {
	blueprint := &Blueprint{}

	fmt.Sscanf(
		line,
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&blueprint.index,
		&blueprint.OreCost,
		&blueprint.ClayCost,
		&blueprint.ObsidianCost[0],
		&blueprint.ObsidianCost[1],
		&blueprint.GeodeCost[0],
		&blueprint.GeodeCost[1],
	)

	return blueprint
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
