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

type Point struct {
	x int
	y int
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

func GenDirections(index int) [3]Point {
	direction := directions[index]

	if direction[0] == 0 {
		return [3]Point{
			{-1, direction[1]}, {0, direction[1]}, {1, direction[1]},
		}

	} else {
		return [3]Point{
			{direction[0], -1}, {direction[0], 0}, {direction[0], 1},
		}
	}
}

type Board struct {
	elves  []*Elf
	coords map[Point]bool
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
	coord := make(map[Point]bool)

	for i, row := range grid {
		for j := 0; j < len(row); j++ {
			if grid[i][j] == '#' {
				elf := &Elf{x: j, y: i, pX: j, pY: i}

				coord[Point{j, i}] = true
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

				x := elf.x + j
				y := elf.y + i

				key := Point{x, y}

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
				x := elf.x + d.x
				y := elf.y + d.y

				key := Point{x, y}

				if b.coords[key] {
					canMove = false
					break
				}
			}

			if canMove {
				elf.pX = elf.x + directions[1].x
				elf.pY = elf.y + directions[1].y
				break
			}
		}
	}
}

func (b *Board) Move() {
	elves := map[Point]*Elf{}

	for _, elf := range b.elves {
		coords := Point{elf.pX, elf.pY}

		if _, ok := elves[coords]; ok {
			delete(elves, coords)
		} else {
			elves[coords] = elf
		}
	}

	for _, elf := range elves {
		delete(b.coords, Point{elf.x, elf.y})
		elf.x = elf.pX
		elf.y = elf.pY
		b.coords[Point{elf.x, elf.y}] = true

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

			key := Point{j, i}

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
		step:     10 * time.Millisecond,
	}
	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}

}
