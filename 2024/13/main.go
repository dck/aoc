package main

import (
	"fmt"
	"os"
)

type Machine struct {
	ax     int
	ay     int
	bx     int
	by     int
	prizex int
	prizey int
}

func (m *Machine) Solve() (int, int) {
	d := m.ax*m.by - m.bx*m.ay

	da := m.prizex*m.by - m.bx*m.prizey
	db := m.ax*m.prizey - m.prizex*m.ay

	if da%d != 0 || db%d != 0 {
		return -1, -1
	}

	a := da / d
	b := db / d

	return a, b
}

func main() {
	var machines []Machine

	for {
		var ax, ay, bx, by, prizex, prizey int

		if _, err := fmt.Fscanf(os.Stdin, "Button A: X+%d, Y+%d\n", &ax, &ay); err != nil {
			break
		}
		if _, err := fmt.Fscanf(os.Stdin, "Button B: X+%d, Y+%d\n", &bx, &by); err != nil {
			break
		}
		if _, err := fmt.Fscanf(os.Stdin, "Prize: X=%d, Y=%d\n", &prizex, &prizey); err != nil {
			break
		}

		machine := Machine{ax: ax, ay: ay, bx: bx, by: by, prizex: prizex, prizey: prizey}
		machines = append(machines, machine)
		fmt.Scanln()
	}
	fmt.Println("Part 1:", part1(machines))
	fmt.Println("Part 2:", part2(machines))
}

func part1(machines []Machine) int {
	coins := 0
	for _, m := range machines {
		a, b := m.Solve()
		if a == -1 || b == -1 {
			continue
		}
		coins += 3*a + b
	}
	return coins
}

func part2(machines []Machine) int {
	coins := 0
	for _, m := range machines {
		m.prizex += 10000000000000
		m.prizey += 10000000000000

		a, b := m.Solve()
		if a == -1 || b == -1 {
			continue
		}
		coins += 3*a + b
	}
	return coins
}
