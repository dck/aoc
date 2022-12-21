package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	VAL = 0
	OP  = 1
)

type Monkey struct {
	kind  int
	value int
	op    string
	left  string
	right string
	x     bool
	cache int
}

//root: pppw + sjmn

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	monkeys := map[string]*Monkey{}
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		name := strings.TrimSpace(parts[0])
		fields := strings.Fields(strings.TrimSpace(parts[1]))

		if len(fields) > 2 {
			left := strings.TrimSpace(fields[0])
			op := strings.TrimSpace(fields[1])
			right := strings.TrimSpace(fields[2])

			monkeys[name] = &Monkey{kind: OP, left: left, op: op, right: right}
		} else {
			num, _ := strconv.Atoi(fields[0])
			monkeys[name] = &Monkey{kind: VAL, value: num}
		}
	}

	Calculate(monkeys, "root")
	Mark(monkeys, "root")

	fmt.Println(Solve(monkeys, "root", 0))
}

func Calculate(monkeys map[string]*Monkey, key string) int {
	m, ok := monkeys[key]
	if !ok {
		return 0
	}

	if m.kind == VAL {
		m.cache = m.value
	} else {
		switch m.op {
		case "+":
			m.cache = Calculate(monkeys, m.left) + Calculate(monkeys, m.right)
		case "-":
			m.cache = Calculate(monkeys, m.left) - Calculate(monkeys, m.right)
		case "*":
			m.cache = Calculate(monkeys, m.left) * Calculate(monkeys, m.right)
		case "/":
			m.cache = Calculate(monkeys, m.left) / Calculate(monkeys, m.right)
		}
	}

	return m.cache
}

func Mark(monkeys map[string]*Monkey, key string) bool {
	m := monkeys[key]

	if m.kind == VAL {
		m.x = key == "humn"
	} else {
		m.x = Mark(monkeys, m.left) || Mark(monkeys, m.right)
	}

	return m.x
}

func Solve(monkeys map[string]*Monkey, key string, desired int) int {
	m := monkeys[key]

	left := monkeys[m.left]
	right := monkeys[m.right]

	if key == "root" {
		if left.x {
			return Solve(monkeys, m.left, right.cache)
		} else {
			return Solve(monkeys, m.right, left.cache)
		}
	}

	if m.left == "humn" || m.right == "humn" {
		var known int
		if m.left == "humn" {
			known = right.cache
		} else {
			known = left.cache
		}

		switch m.op {
		case "+":
			return desired - known
		case "-":
			return desired + known
		case "*":
			return desired / known
		case "/":
			return desired * known
		}
	}

	if left.x {
		switch m.op {
		case "+":
			desired = desired - right.cache
		case "-":
			desired = desired + right.cache
		case "*":
			desired = desired / right.cache
		case "/":
			desired = desired * right.cache
		}
		return Solve(monkeys, m.left, desired)
	} else {
		switch m.op {
		case "+":
			desired = desired - left.cache
		case "-":
			desired = left.cache - desired
		case "*":
			desired = desired / left.cache
		case "/":
			desired = left.cache / desired
		}
		return Solve(monkeys, m.right, desired)
	}
}
