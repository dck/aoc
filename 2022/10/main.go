package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Instruction string
	Arg         int
}

type VM struct {
	Pixels [][]rune
	width  int
	height int
	x      int
	cycle  int
}

func CreateVM(width int, height int) *VM {
	vm := &VM{width: width, height: height, x: 1, cycle: 0}

	vm.Pixels = make([][]rune, vm.height)

	for i := 0; i < vm.height; i++ {
		vm.Pixels[i] = make([]rune, vm.width)
		for j := 0; j < vm.width; j++ {
			vm.Pixels[i][j] = '.'
		}
	}

	return vm
}

func (vm *VM) tick() {
	if math.Abs(float64(vm.cycle%vm.width-vm.x)) < 2 {
		i := vm.cycle / vm.width
		j := vm.cycle % vm.width
		vm.Pixels[i][j] = '#'
	}

	vm.cycle++
}

func (vm *VM) Execute(commands []Command) {
	for _, c := range commands {
		if c.Instruction == "noop" {
			vm.tick()
		}
		if c.Instruction == "addx" {
			vm.tick()
			vm.tick()
			vm.x += c.Arg
		}
	}
}

func (vm *VM) String() string {
	var s strings.Builder

	for i := 0; i < vm.height; i++ {
		fmt.Fprintf(&s, "%d [", i+1)
		for j := 0; j < vm.width; j++ {
			fmt.Fprintf(&s, "%c", vm.Pixels[i][j])
		}
		fmt.Fprint(&s, "]\n")
	}

	return s.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := make([]Command, 0)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		instruction := parts[0]
		arg := 0
		if len(parts) > 1 {
			arg, _ = strconv.Atoi(parts[1])
		}

		commands = append(commands, Command{Instruction: instruction, Arg: arg})
	}

	vm := CreateVM(40, 6)
	vm.Execute(commands)
	fmt.Println(vm)
}
