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
}

//root: pppw + sjmn

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	monkeys := map[string]Monkey{}
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		name := strings.TrimSpace(parts[0])
		fields := strings.Fields(strings.TrimSpace(parts[1]))

		if len(fields) > 2 {
			left := strings.TrimSpace(fields[0])
			op := strings.TrimSpace(fields[1])
			right := strings.TrimSpace(fields[2])

			monkeys[name] = Monkey{kind: OP, left: left, op: op, right: right}
		} else {
			num, _ := strconv.Atoi(fields[0])
			monkeys[name] = Monkey{kind: VAL, value: num}
		}
	}

	fmt.Println(Calculate(monkeys, "root"))
}

func Calculate(monkeys map[string]Monkey, key string) int {
	m, ok := monkeys[key]
	if !ok {
		return 0
	}
	if m.kind == VAL {
		return m.value
	} else {
		switch m.op {
		case "+":
			return Calculate(monkeys, m.left) + Calculate(monkeys, m.right)
		case "-":
			return Calculate(monkeys, m.left) - Calculate(monkeys, m.right)
		case "*":
			return Calculate(monkeys, m.left) * Calculate(monkeys, m.right)
		case "/":
			return Calculate(monkeys, m.left) / Calculate(monkeys, m.right)
		}
	}

	return 0
}
