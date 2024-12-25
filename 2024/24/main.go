package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Registry map[string]Node

func (r Registry) Value(num byte) int {
	bits := []string{}

	for _, node := range NodesRegistry {
		if node.Name()[0] == num {
			bits = append(bits, node.Name())
		}
	}

	sort.Strings(bits)

	res := 0
	for i, z := range bits {
		val := NodesRegistry[z].Value()
		res |= int(val) << i
	}

	return res
}

func (r Registry) NodesByType(t string) []string {
	res := []string{}
	for _, node := range NodesRegistry {
		if reflect.TypeOf(node).Name() == t {
			res = append(res, node.Name())
		}
	}
	return res
}

type Node interface {
	Value() uint8
	Name() string
}

type ValueNode struct {
	name  string
	value uint8
}

func (n ValueNode) Value() uint8 {
	return n.value
}

func (n ValueNode) Name() string {
	return n.name
}

type XorNode struct {
	name  string
	left  string
	right string
}

func (n XorNode) Value() uint8 {
	return NodesRegistry[n.left].Value() ^ NodesRegistry[n.right].Value()
}

func (n XorNode) Name() string {
	return n.name
}

type AndNode struct {
	name  string
	left  string
	right string
}

func (n AndNode) Value() uint8 {
	return NodesRegistry[n.left].Value() & NodesRegistry[n.right].Value()
}

func (n AndNode) Name() string {
	return n.name
}

type OrNode struct {
	name  string
	left  string
	right string
}

func (n OrNode) Value() uint8 {
	return NodesRegistry[n.left].Value() | NodesRegistry[n.right].Value()
}

func (n OrNode) Name() string {
	return n.name
}

var NodesRegistry = make(Registry)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.Contains(line, "->") {
			parts := strings.Split(line, " -> ")
			name := parts[1]
			declaration := strings.Fields(parts[0])
			left := declaration[0]
			operator := declaration[1]
			right := declaration[2]
			switch operator {
			case "AND":
				NodesRegistry[name] = AndNode{name, left, right}
			case "OR":
				NodesRegistry[name] = OrNode{name, left, right}
			case "XOR":
				NodesRegistry[name] = XorNode{name, left, right}
			}
		} else {
			parts := strings.Split(line, ": ")
			name := parts[0]
			ValueString := parts[1]
			value, _ := strconv.Atoi(ValueString)
			NodesRegistry[name] = ValueNode{name, uint8(value)}
		}
	}

	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {
	return NodesRegistry.Value('z')
}

func part2() string {
	wrongGates := []string{}
	zGates := []string{}
	for _, node := range NodesRegistry {
		if node.Name()[0] == 'z' {
			zGates = append(zGates, node.Name())
		}
	}
	sort.Strings(zGates)

	for i := 0; i < len(zGates)-1; i++ {
		if _, ok := NodesRegistry[zGates[i]].(XorNode); !ok {
			wrongGates = append(wrongGates, zGates[i])
		}
	}

	for _, middleGate := range NodesRegistry {
		if middleGate.Name()[0] == 'z' {
			continue
		}

		if _, ok := middleGate.(XorNode); ok {
			leftName := middleGate.(XorNode).left
			rightName := middleGate.(XorNode).right

			if leftName[0] == 'x' || leftName[0] == 'y' || rightName[0] == 'x' || rightName[0] == 'y' {
				continue
			}

			wrongGates = append(wrongGates, middleGate.Name())
		}
	}

	for _, xorGate := range NodesRegistry.NodesByType("XorNode") {
		node := NodesRegistry[xorGate].(XorNode)
		if (node.left[0] == 'x' || node.left[0] == 'y') && (node.right[0] == 'x' || node.right[0] == 'y') {
			if node.Name()[0] == 'z' {
				continue
			}

			found := false
			for _, nextXor := range NodesRegistry.NodesByType("XorNode") {
				leftName := NodesRegistry[nextXor].(XorNode).left
				rightName := NodesRegistry[nextXor].(XorNode).right

				if leftName == node.Name() || rightName == node.Name() {
					found = true
					break
				}
			}

			if !found {
				wrongGates = append(wrongGates, node.Name())
			}
		}
	}

	for _, andGate := range NodesRegistry.NodesByType("AndNode") {
		node := NodesRegistry[andGate].(AndNode)

		if (node.left == "x00" || node.left == "y00") && (node.right == "x00" || node.right == "y00") {
			continue
		}

		found := false
		for _, nextOr := range NodesRegistry.NodesByType("OrNode") {
			leftName := NodesRegistry[nextOr].(OrNode).left
			rightName := NodesRegistry[nextOr].(OrNode).right

			if leftName == node.Name() || rightName == node.Name() {
				found = true
				break
			}
		}

		if !found {
			wrongGates = append(wrongGates, node.Name())
		}
	}
	wrongGates = Uniq(wrongGates)
	sort.Strings(wrongGates)

	return strings.Join(wrongGates, ",")
}

func Uniq(s []string) []string {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}

	res := []string{}
	for k := range m {
		res = append(res, k)
	}

	return res
}
