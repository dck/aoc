package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Network map[string]map[string]bool

func main() {
	network := make(Network)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, "-")

		if network[parts[0]] == nil {
			network[parts[0]] = make(map[string]bool)
		}

		if network[parts[1]] == nil {
			network[parts[1]] = make(map[string]bool)
		}
		network[parts[0]][parts[1]] = true
		network[parts[1]][parts[0]] = true
	}

	fmt.Println("Part 1:", part1(network))
	fmt.Println("Part 2:", part2(network))
}

func part1(network Network) int {
	res := map[[3]string]bool{}
	for a, subnetwork := range network {
		aStarts := strings.HasPrefix(a, "t")
		for b := range subnetwork {
			bStarts := strings.HasPrefix(b, "t")
			for c := range network[b] {
				startWithT := aStarts || bStarts || strings.HasPrefix(c, "t")

				if startWithT && network[a][b] && network[a][c] && network[b][c] {
					key := [3]string{a, b, c}
					sort.Strings(key[:])
					res[key] = true
				}
			}
		}

	}
	return len(res)
}

func part2(network Network) string {
	var maxSet Set
	for node := range network {
		group := make(Set)
		size := dfs(network, group, node)
		if size > len(maxSet) {
			maxSet = group
		}
	}

	keys := make([]string, 0, len(maxSet))
	for k := range maxSet {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return strings.Join(keys, ",")
}

type Set map[string]bool

func dfs(network Network, group Set, node string) int {
	group[node] = true
	subset := make(Set)
	for connection := range network[node] {
		if group[connection] {
			continue
		} else {
			subset[connection] = true
		}
	}

	for connection := range subset {
		all := true
		for n := range group {
			if !network[connection][n] {
				all = false
			}
		}

		if all {
			dfs(network, group, connection)
		}
	}

	return len(group)
}
