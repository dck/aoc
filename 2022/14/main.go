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

type Point struct {
	x int
	y int
}

type Sand struct {
	Point
	rest bool
}

func (s *Sand) Move(c *Cave) {
	moved := false

	if c.Canvas[s.y+1][s.x-400] == '.' {
		c.Canvas[s.y][s.x-400] = '.'
		c.Canvas[s.y+1][s.x-400] = '*'
		s.y += 1
		moved = true
		return
	}

	if c.Canvas[s.y+1][s.x-400-1] == '.' {
		c.Canvas[s.y][s.x-400] = '.'
		c.Canvas[s.y+1][s.x-400-1] = '*'
		s.y += 1
		s.x -= 1
		moved = true
		return
	}

	if c.Canvas[s.y+1][s.x-400+1] == '.' {
		c.Canvas[s.y][s.x-400] = '.'
		c.Canvas[s.y+1][s.x-400+1] = '*'
		s.y += 1
		s.x += 1
		moved = true
		return
	}

	if !moved {
		s.rest = true
	}
}

type Cave struct {
	Sands       []Sand
	Canvas      [][]rune
	width       int
	height      int
	currentSand *Sand

	step     time.Duration
	tickChan chan<- Refresh
}

type Refresh struct {
	quit bool
}

func CreateCave(width int, height int, step time.Duration, tickChan chan<- Refresh) *Cave {
	c := &Cave{
		width:    width,
		height:   height,
		step:     step,
		tickChan: tickChan,
	}

	c.Canvas = make([][]rune, c.height)

	for i := 0; i < c.height; i++ {
		c.Canvas[i] = make([]rune, c.width)
		for j := 0; j < c.width; j++ {
			c.Canvas[i][j] = '.'
		}
	}

	return c
}

func (c *Cave) DrawRock(from, to Point) {
	if from.x == to.x {
		step := map[bool]int{true: 1, false: -1}[from.y < to.y]

		i := from.y
		for {
			c.drawCell(from.x, i, '#')
			if i == to.y {
				break
			}
			i += step
		}
	}

	if from.y == to.y {
		step := map[bool]int{true: 1, false: -1}[from.x < to.x]
		i := from.x
		for {
			c.drawCell(i, from.y, '#')
			if i == to.x {
				break
			}
			i += step
		}
	}
}

func (c *Cave) AddSand() {
	sand := Sand{Point: Point{500, 0}}

	c.Sands = append(c.Sands, sand)

	c.currentSand = &sand

	c.drawCell(sand.x, sand.y, '*')
}

func (c *Cave) drawCell(x int, y int, m rune) {
	c.Canvas[y][x-400] = m
}

func (c *Cave) tick() {
	if c.currentSand == nil || c.currentSand.rest {
		c.AddSand()
	} else {
		c.currentSand.Move(c)
	}

	if c.currentSand.y >= c.height-1 {
		c.tickChan <- Refresh{quit: true}
	}

	if c.currentSand.rest && c.currentSand.x == 500 && c.currentSand.y == 0 {
		c.tickChan <- Refresh{quit: true}
	}

	c.tickChan <- Refresh{}
	time.Sleep(c.step)
}

func waitForActivity(sub chan Refresh) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

type Model struct {
	Cave *Cave
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

	for i := m.Cave.height - 50; i < m.Cave.height; i++ {
		for j := 0; j < m.Cave.width; j++ {
			fmt.Fprintf(&s, "%c", m.Cave.Canvas[i][j])
		}
		fmt.Fprint(&s, "\n")
	}

	return s.String()
}

func ParsePoint(line string) Point {
	coordsStr := strings.SplitN(line, ",", 2)
	x, _ := strconv.Atoi(strings.TrimSpace(coordsStr[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(coordsStr[1]))

	return Point{x, y}
}

func main() {
	refreshChan := make(chan Refresh)
	cave := CreateCave(500, 170, 1*time.Millisecond, refreshChan)

	maxY := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "->")

		for i := 0; i < len(parts)-1; i++ {
			from := ParsePoint(parts[i])
			to := ParsePoint(parts[i+1])

			if from.y > maxY {
				maxY = from.y
			}
			if to.y > maxY {
				maxY = to.y
			}
			cave.DrawRock(from, to)
		}

	}

	fmt.Println(Point{400, maxY + 2}, Point{400 + cave.width - 1, maxY + 2})
	cave.DrawRock(Point{400, maxY + 2}, Point{400 + cave.width - 1, maxY + 2})

	go func(c *Cave) {
		for {
			cave.tick()
		}
	}(cave)

	p := tea.NewProgram(&Model{
		Sub:  refreshChan,
		Cave: cave,
	})

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}

	fmt.Println(len(cave.Sands) - 1)
}
