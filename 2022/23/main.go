package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var directions [4][2]int = [4][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Elf struct {
	x  int
	y  int
	pX int
	pY int
}

func GenDirections(index int) [3][2]int {
	direction := directions[index]

	if direction[0] == 0 {
		return [3][2]int{
			{-1, direction[1]}, {0, direction[1]}, {1, direction[1]},
		}

	} else {
		return [3][2]int{
			{direction[0], -1}, {direction[0], 0}, {direction[0], 1},
		}
	}
}

type Board struct {
	elves  []*Elf
	coords map[[2]int]bool
	round  int

	top    int
	bottom int
	left   int
	right  int
}

type Refresh struct {
	quit bool
}

func CreateBoard(grid []string) *Board {
	elves := make([]*Elf, 0)
	coord := make(map[[2]int]bool)

	for i, row := range grid {
		for j := 0; j < len(row); j++ {
			if grid[i][j] == '#' {
				elf := &Elf{x: j, y: i, pX: j, pY: i}

				coord[[2]int{j, i}] = true
				elves = append(elves, elf)
			}
		}
	}

	b := &Board{
		elves:  elves,
		coords: coord,
		top:    math.MaxInt,
		bottom: math.MinInt,
		left:   math.MaxInt,
		right:  math.MinInt,
	}
	return b
}

func (b *Board) Score() int {
	return (b.bottom-b.top+1)*(b.right-b.left+1) - len(b.elves)
}

func (b *Board) Propose() {
	for _, elf := range b.elves {

		needMove := false
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				x := elf.x + i
				y := elf.y + j

				key := [2]int{x, y}

				if b.coords[key] {
					needMove = true
				}
			}
		}

		if !needMove {
			elf.pX = elf.x
			elf.pY = elf.y
			continue
		}

		for i := 0; i < len(directions); i++ {
			directions := GenDirections((i + b.round) % len(directions))

			canMove := true
			for _, d := range directions {
				x := elf.x + d[0]
				y := elf.y + d[1]

				key := [2]int{x, y}

				if b.coords[key] {
					canMove = false
				}
			}

			if canMove {
				elf.pX = elf.x + directions[1][0]
				elf.pY = elf.y + directions[1][1]
				break
			} else {
				elf.pX = elf.x
				elf.pY = elf.y
			}
		}
	}
}

func (b *Board) Move() {
	elves := map[[2]int]*Elf{}

	for _, elf := range b.elves {
		coords := [2]int{elf.pX, elf.pY}

		if pair, ok := elves[coords]; ok {
			delete(elves, coords)

			elf.pX = elf.x
			elf.pY = elf.y

			pair.pX = pair.x
			pair.pY = pair.y
		} else {
			elves[coords] = elf
		}
	}

	for _, elf := range elves {
		delete(b.coords, [2]int{elf.x, elf.y})
		elf.x = elf.pX
		elf.y = elf.pY
		b.coords[[2]int{elf.x, elf.y}] = true

		b.top = Min(b.top, elf.y)
		b.bottom = Max(b.bottom, elf.y)

		b.left = Min(b.left, elf.x)
		b.right = Max(b.right, elf.x)
	}

	b.round++
}

type Model struct {
	Board    *Board
	Sub      chan Refresh
	step     time.Duration
	tickChan chan<- Refresh
}

func waitForActivity(sub chan Refresh) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

func (m *Model) Init() tea.Cmd {
	return waitForActivity(m.Sub)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.Board.Propose()
			m.Board.Move()
		}
	case Refresh:
		if msg.quit {
			return m, tea.Quit
		} else {
			return m, waitForActivity(m.Sub)
		}
	}

	return m, nil
}

func (m *Model) View() string {
	var s strings.Builder

	for i := m.Board.top - 5; i <= m.Board.bottom+5; i++ {
		for j := m.Board.left - 5; j <= m.Board.right+5; j++ {

			key := [2]int{j, i}

			if m.Board.coords[key] {
				fmt.Fprint(&s, "#")
			} else {
				fmt.Fprint(&s, ".")
			}
		}
		fmt.Fprintln(&s)
	}

	fmt.Fprintf(&s, "Round: %d    Score: %d", m.Board.round, m.Board.Score())
	return s.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := []string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		grid = append(grid, line)
	}

	refreshChan := make(chan Refresh)
	board := CreateBoard(grid)

	model := &Model{
		Sub:      refreshChan,
		Board:    board,
		tickChan: refreshChan,
		step:     250 * time.Millisecond,
	}
	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}

}
