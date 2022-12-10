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

type CRT struct {
	commands []Command
	Pixels   [][]rune
	width    int
	height   int
	x        int
	cycle    int
}

func CreateCRT(commands []Command) *CRT {
	crt := &CRT{width: 40, height: 6, commands: commands, x: 1, cycle: 0}

	crt.Pixels = make([][]rune, crt.height)

	for i := 0; i < crt.height; i++ {
		crt.Pixels[i] = make([]rune, crt.width)
		for j := 0; j < crt.width; j++ {
			crt.Pixels[i][j] = '.'
		}
	}

	for _, c := range commands {
		if math.Abs(float64(crt.cycle%40-crt.x)) < 2 {
			i := crt.cycle / crt.width
			j := crt.cycle % crt.width
			crt.Pixels[i][j] = '#'
		}

		if c.Instruction == "noop" {
			crt.cycle++
			continue
		}
		if c.Instruction == "addx" {
			crt.cycle++
			if math.Abs(float64(crt.cycle%40-crt.x)) < 2 {
				i := crt.cycle / crt.width
				j := crt.cycle % crt.width
				crt.Pixels[i][j] = '#'
			}
			crt.cycle++
			crt.x += c.Arg
		}
	}

	return crt
}

func (crt *CRT) String() string {
	var s strings.Builder

	for i := 0; i < crt.height; i++ {
		fmt.Fprintf(&s, "%d [", i+1)
		for j := 0; j < crt.width; j++ {
			fmt.Fprintf(&s, "%c", crt.Pixels[i][j])
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

	crt := CreateCRT(commands)
	fmt.Println(crt)
}
