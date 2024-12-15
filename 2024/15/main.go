package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Deque[T comparable] struct {
	items []T
	set   map[T]bool
}

func NewDeque[T comparable]() *Deque[T] {
	return &Deque[T]{
		items: make([]T, 0),
		set:   make(map[T]bool),
	}
}

func (d *Deque[T]) Add(item T) {
	if _, exists := d.set[item]; !exists {
		d.items = append(d.items, item)
		d.set[item] = true
	}
}

func (d *Deque[T]) At(idx int) T {
	return d.items[idx]
}

func (d *Deque[T]) Length() int {
	return len(d.items)
}

var directions = map[rune][2]int{
	'<': {0, -1},
	'>': {0, 1},
	'^': {-1, 0},
	'v': {1, 0},
}
var Reset = "\033[0m"
var Red = "\033[31m"
var Blue = "\033[34m"

type Robot struct {
	i, j int
}

type Game struct {
	version     uint8
	isManual    bool
	robot       Robot
	grid        [][]byte
	moves       string
	currentMove int
}

func (g *Game) Run(duration time.Duration, sub chan Refresh) {
	for i, move := range g.moves {
		g.currentMove = i
		sub <- Refresh{movement: move, quit: false}
		time.Sleep(duration)
	}
	sub <- Refresh{quit: true}
}

func (g *Game) GPSSum() int {
	sum := 0
	for i, row := range g.grid {
		for j, cell := range row {
			if cell == 'O' || cell == '[' {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func (g *Game) Move(move rune) {
	deltas := directions[move]

	newI := g.robot.i + deltas[0]
	newJ := g.robot.j + deltas[1]

	if g.grid[newI][newJ] == '#' {
		return
	}

	if g.grid[newI][newJ] == '.' {
		g.robot.i = newI
		g.robot.j = newJ
		return
	}

	if move == '>' || move == '<' {
		steps := 1
		for {
			if g.grid[newI][newJ] == '#' {
				break
			}

			if g.grid[newI][newJ] == '.' {
				for k := 0; k < steps; k++ {
					g.grid[newI][newJ] = g.grid[newI-deltas[0]][newJ-deltas[1]]
					newI -= deltas[0]
					newJ -= deltas[1]
				}
				g.robot.i = newI + deltas[0]
				g.robot.j = newJ + deltas[1]
				break
			}

			newI += deltas[0]
			newJ += deltas[1]
			steps++
		}
	}

	if move == '^' || move == 'v' {
		deque := NewDeque[[2]int]()
		deque.Add([2]int{g.robot.i, g.robot.j})

		idx := 0
		canMove := true
		for {
			point := deque.At(idx)
			idx++
			newI := point[0] + deltas[0]
			newJ := point[1] + deltas[1]
			if g.grid[newI][newJ] == '#' {
				canMove = false
				break
			}
			if g.grid[newI][newJ] == ']' {
				deque.Add([2]int{newI, newJ})
				deque.Add([2]int{newI, newJ - 1})
			} else if g.grid[newI][newJ] == '[' {
				deque.Add([2]int{newI, newJ})
				deque.Add([2]int{newI, newJ + 1})
			}

			if idx == deque.Length() {
				break
			}
		}

		if canMove {
			for idx := deque.Length() - 1; idx > 0; idx-- {
				point := deque.At(idx)
				g.grid[point[0]+deltas[0]][point[1]+deltas[1]], g.grid[point[0]][point[1]] = g.grid[point[0]][point[1]], g.grid[point[0]+deltas[0]][point[1]+deltas[1]]
			}

			g.robot.i = deque.At(0)[0] + deltas[0]
			g.robot.j = deque.At(0)[1] + deltas[1]
		}
	}
}

type Refresh struct {
	movement rune
	quit     bool
}

func waitForActivity(sub chan Refresh) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

type Model struct {
	game *Game
	Sub  chan Refresh
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
			m.game.Move('>')
		case "left":
			m.game.Move('<')
		case "up":
			m.game.Move('^')
		case "down":
			m.game.Move('v')
		}
	case Refresh:
		if msg.quit {
			return m, tea.Quit
		} else {
			m.game.Move(msg.movement)
			return m, waitForActivity(m.Sub)
		}
	}

	return m, nil
}

func (m *Model) View() string {
	var s strings.Builder

	game := m.game

	for i, row := range game.grid {
		for j, cell := range row {
			if i == game.robot.i && j == game.robot.j {
				fmt.Fprint(&s, Blue+"@"+Reset)
			} else if cell == '#' {
				fmt.Fprint(&s, Red+"#"+Reset)
			} else {
				fmt.Fprint(&s, string(cell))
			}
		}
		fmt.Fprintln(&s)
	}

	fmt.Fprintf(&s, "\nCurrent Move: \033[1m%s"+Reset+"\n", string(game.moves[game.currentMove]))
	fmt.Fprintf(&s, "GPS Score: \033[1m%d\n"+Reset, game.GPSSum())

	return s.String()
}

func main() {
	var game Game

	if len(os.Args) > 1 && os.Args[1] == "-2" {
		game.version = 2
	} else {
		game.version = 1
	}

	if len(os.Args) > 2 && os.Args[2] == "-m" {
		game.isManual = true
	}

	scanner := bufio.NewScanner(os.Stdin)
	consumeMoves := false
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			consumeMoves = true
		}

		if consumeMoves {
			game.moves += line
		} else {
			if game.version == 1 {
				game.grid = append(game.grid, []byte(line))
			} else {
				row := []byte{}
				for _, cell := range line {
					switch cell {
					case '@':
						row = append(row, '@', '.')
					case 'O':
						row = append(row, '[', ']')
					case '#':
						row = append(row, '#', '#')
					case '.':
						row = append(row, '.', '.')
					}
				}
				game.grid = append(game.grid, row)
			}
		}
	}

	for i, row := range game.grid {
		for j, cell := range row {
			if cell == '@' {
				game.robot.i = i
				game.robot.j = j
				game.grid[i][j] = '.'
			}
		}
	}

	refreshChan := make(chan Refresh)
	model := &Model{game: &game, Sub: refreshChan}
	p := tea.NewProgram(model)

	if !game.isManual {
		go game.Run(10*time.Millisecond, refreshChan)
	}

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}
