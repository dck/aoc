package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var directions [4][2]int = [4][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type Player struct {
	x         int
	y         int
	direction [2]int
}

type Board struct {
	grid   [][]byte
	width  int
	height int
	player Player
}

type Refresh struct {
	quit bool
}

func CreateBoard(grid []string) *Board {
	b := &Board{}

	b.height = len(grid)
	for _, line := range grid {
		b.width = Max(b.width, len(line))
	}

	b.grid = make([][]byte, b.height)
	playerX := 0

	for i, line := range grid {
		b.grid[i] = make([]byte, b.width)

		j := 0
		for j < len(line) {
			if i == 0 && line[j] == ' ' {
				playerX++
			}

			b.grid[i][j] = line[j]
			j++
		}
		for j < b.width {
			b.grid[i][j] = ' '
			j++
		}
	}

	b.player = Player{x: playerX, y: 0, direction: [2]int{1, 0}}

	return b
}

func (b *Board) Password() int {
	row := b.player.y + 1
	column := b.player.x + 1
	facing := map[[2]int]int{
		directions[0]: 0,
		directions[1]: 1,
		directions[2]: 2,
		directions[3]: 3,
	}[b.player.direction]

	return 1000*row + 4*column + facing
}

func (b *Board) Move() {
	if b.player.direction[0] == 1 {
		// right edge
		x := b.player.x + 1
		line := b.grid[b.player.y]
		if x >= len(line) || line[x] == ' ' {
			x = 0
			for x < len(line) {
				if line[x] == ' ' {
					x++
				} else {
					break
				}
			}
		}
		if line[x] == '.' {
			b.player.x = x
		}
	}

	if b.player.direction[0] == -1 {
		// left edge
		x := b.player.x - 1
		line := b.grid[b.player.y]

		if x < 0 || line[x] == ' ' {
			x = len(line) - 1
			for x >= 0 {
				if line[x] == ' ' {
					x--
				} else {
					break
				}
			}
		}

		if line[x] == '.' {
			b.player.x = x
		}
	}

	if b.player.direction[1] == 1 {
		// bottom edge
		y := b.player.y + 1
		if y >= len(b.grid) || b.grid[y][b.player.x] == ' ' {
			y = 0
			for y < len(b.grid) {
				if b.grid[y][b.player.x] == ' ' {
					y++
				} else {
					break
				}
			}
		}

		if b.grid[y][b.player.x] == '.' {
			b.player.y = y
		}
	}

	if b.player.direction[1] == -1 {
		// bottom edge
		y := b.player.y - 1
		if y < 0 || b.grid[y][b.player.x] == ' ' {
			y = len(b.grid) - 1
			for y > 0 {
				if b.grid[y][b.player.x] == ' ' {
					y--
				} else {
					break
				}
			}
		}

		if b.grid[y][b.player.x] == '.' {
			b.player.y = y
		}
	}
}

func (b *Board) ChangeDirection(where byte) {
	pos := -1
	for i, d := range directions {
		if d == b.player.direction {
			pos = i
		}
	}

	if where == 'R' {
		pos++
		if pos >= len(directions) {
			pos = 0
		}

		b.player.direction = directions[pos]
	} else if where == 'L' {
		pos--
		if pos < 0 {
			pos = len(directions) - 1
		}

		b.player.direction = directions[pos]
	} else {
		panic("Unknown direction")
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
		case "right":
			m.Board.Move()
		case "up":
			m.Board.ChangeDirection('R')
		case "down":
			m.Board.ChangeDirection('L')
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

	for y, line := range m.Board.grid {
		for x, letter := range line {
			if m.Board.player.x == x && m.Board.player.y == y {
				fmt.Fprint(&s, "X")
			} else {
				fmt.Fprintf(&s, "%c", letter)
			}
		}
		fmt.Fprintln(&s)
	}

	c := map[[2]int]byte{
		directions[0]: '>',
		directions[1]: 'v',
		directions[2]: '<',
		directions[3]: '^',
	}[m.Board.player.direction]

	fmt.Fprintf(&s, "Player (%d, %d) %c    Password: %d", m.Board.player.x, m.Board.player.y, c, m.Board.Password())

	return s.String()
}

func splitCommands(s string) []string {
	var a []string
	var j int
	for i, r := range s {
		if r == 'R' || r == 'L' {
			a = append(a, s[j:i])
			a = append(a, s[i:i+1])
			j = i + 1
		}
	}
	a = append(a, s[j:])
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	steps := []string{}
	grid := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, line)

		if line == "" {
			scanner.Scan()
			line = strings.TrimSpace(scanner.Text())
			steps = splitCommands(line)
			break
		}
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

	go func(steps []string) {
		for _, s := range steps {
			if s == "R" || s == "L" {
				model.Board.ChangeDirection(byte(s[0]))
				continue
			}

			total, _ := strconv.Atoi(s)

			for i := 0; i < total; i++ {
				model.Board.Move()
				model.tickChan <- Refresh{}
				time.Sleep(model.step)
			}

		}
		model.Sub <- Refresh{quit: true}

	}(steps)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}

	fmt.Printf("Password: %d\n", board.Password())
}
