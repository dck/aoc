package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

type Network map[string]Node

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	instructions := scanner.Text()

	network := Network{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)

		fields := strings.Split(line, " = ")
		nodeName := fields[0]

		children := strings.Split(strings.Trim(fields[1], "()"), ", ")
		network[nodeName] = Node{children[0], children[1]}
	}

	fmt.Println(part1(instructions, network, "AAA", []string{"ZZZ"}))
	fmt.Println(part2(instructions, network))
}

func cycle(s string) <-chan byte {
	ch := make(chan byte)
	go func() {
		i := 0
		for i < len(s) {
			ch <- s[i]
			i++
			if i == len(s) {
				i = 0
			}
		}
	}()
	return ch
}

func part1(instructions string, network Network, start string, finish []string) int {
	curr := start
	steps := 0
	instructionCycle := cycle(instructions)
	for !slices.Contains(finish, curr) {
		direction := <-instructionCycle
		if direction == 'R' {
			curr = network[curr].Right
		} else {
			curr = network[curr].Left
		}
		steps++
	}

	return steps
}

func part2(instructions string, network Network) int {
	starts := []string{}
	finish := []string{}

	for label := range network {
		if strings.HasSuffix(label, "A") {
			starts = append(starts, label)
		}

		if strings.HasSuffix(label, "Z") {
			finish = append(finish, label)
		}
	}

	results := []int{}

	for _, label := range starts {
		results = append(results, part1(instructions, network, label, finish))
	}

	total := results[0]
	for i := 1; i < len(results); i++ {
		total = lcm(total, results[i])
	}

	return total
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
