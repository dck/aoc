package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Computer struct {
	A, B, C uint64
	Program []byte
}

type Buffer []byte

func (buf Buffer) String() string {
	var s strings.Builder
	for _, b := range buf {
		if s.Len() > 0 {
			s.WriteString(",")
		}
		s.WriteString(fmt.Sprintf("%d", b))
	}
	return s.String()
}

func (c *Computer) combo(operand byte) uint64 {
	if operand < 4 {
		return uint64(operand)
	} else if operand == 4 {
		return c.A
	} else if operand == 5 {
		return c.B
	} else if operand == 6 {
		return c.C
	} else {
		panic("Invalid operand")
	}
}

func (c *Computer) Run() Buffer {
	buf := make([]byte, 0)
	pointer := 0
	for pointer < len(c.Program)-1 {
		op := c.Program[pointer]
		arg := c.Program[pointer+1]
		switch op {
		case 0: // adv
			c.A = c.A >> c.combo(arg)
		case 1: // bxl
			c.B = c.B ^ uint64(arg)
		case 2: // bst
			c.B = c.combo(arg) & 7
		case 3: // jnz
			if c.A != 0 {
				pointer = int(arg)
				continue
			}
		case 4: // bxc
			c.B = c.B ^ c.C
		case 5: // out
			buf = append(buf, byte(c.combo(arg)&7))
		case 6: // bdv
			c.B = c.A >> c.combo(arg)
		case 7: // cdv
			c.C = c.A >> c.combo(arg)
		}

		pointer += 2
	}

	return buf
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var c Computer

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Register A:") {
			fmt.Sscanf(line, "Register A: %d", &c.A)
		} else if strings.HasPrefix(line, "Register B:") {
			fmt.Sscanf(line, "Register B: %d", &c.B)
		} else if strings.HasPrefix(line, "Register C:") {
			fmt.Sscanf(line, "Register C: %d", &c.C)
		} else if strings.HasPrefix(line, "Program:") {
			programStr := strings.TrimPrefix(line, "Program: ")
			programParts := strings.Split(programStr, ",")
			for _, part := range programParts {
				var byteVal byte
				fmt.Sscanf(part, "%d", &byteVal)
				c.Program = append(c.Program, byteVal)
			}
		}
	}

	fmt.Println("Part 1:", part1(c))
	fmt.Println("Part 2:", part2(c))
}

func part1(c Computer) string {
	buf := c.Run()
	return buf.String()
}

func part2(c Computer) uint64 {
	var res uint64
	for i := len(c.Program) - 1; i >= 0; i-- {
		res <<= 3

		for j := uint64(0); j <= 7; j++ {
			c.A = res | j
			c.B = 0
			c.C = 0
			buf := c.Run()

			if slices.Equal(buf, c.Program[i:]) {
				res |= j
				break
			}
		}
	}
	return res
}
