package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Point struct {
	x int
	y int
}

type Rock struct {
	Point
	shape [][]int8
	rest  bool

	width  int
	height int
}

type Chamber struct {
	rocks        []Rock
	width        int
	height       int
	currentRock  *Rock
	nextRockType int
}

type Refresh struct {
	quit bool
}

func CreateChamber(width int) *Chamber {
	c := &Chamber{
		width:  width,
		height: 0,
	}

	c.rocks = make([]Rock, 0)

	return c
}

func (c *Chamber) AddNextRocks(x, y int) {
	rock := Rock{}

	if c.nextRockType == 0 {
		rock.shape = [][]int8{
			{1, 1, 1, 1},
		}
		rock.height = 1
		rock.width = 4
	} else if c.nextRockType == 1 {
		rock.shape = [][]int8{
			{0, 1, 0},
			{1, 1, 1},
			{0, 1, 0},
		}
		rock.height = 3
		rock.width = 3
	} else if c.nextRockType == 2 {
		rock.shape = [][]int8{
			{0, 0, 1},
			{0, 0, 1},
			{1, 1, 1},
		}
		rock.height = 3
		rock.width = 3
	} else if c.nextRockType == 3 {
		rock.shape = [][]int8{
			{1},
			{1},
			{1},
			{1},
		}
		rock.height = 4
		rock.width = 1
	} else if c.nextRockType == 4 {
		rock.shape = [][]int8{
			{1, 1},
			{1, 1},
		}
		rock.height = 2
		rock.width = 2
	}

	c.nextRockType++
	if c.nextRockType == 5 {
		c.nextRockType = 0
	}

	rock.x = x
	rock.y = y + 2 + rock.height

	c.rocks = append(c.rocks, rock)
	c.currentRock = &c.rocks[len(c.rocks)-1]
}

func (c *Chamber) MoveTopRock(command rune) {
	if command == '<' {
		if c.currentRock.x > 0 {
			c.currentRock.x--
		}

	} else if command == '>' {
		if c.currentRock.x < c.width-c.currentRock.width {
			c.currentRock.x++
		}
	}
}

func (c *Chamber) FallTopRock() bool {
	if c.currentRock.y > 0 {
		c.currentRock.y--
		return true
	} else {
		c.currentRock.rest = true
		return false
	}
}

func (c *Chamber) tick(command rune) {
	if c.currentRock == nil || c.currentRock.rest {
		c.AddNextRocks(2, c.height)
		c.MoveTopRock(command)
	} else {

		success := c.FallTopRock()
		if success {
			c.MoveTopRock(command)
		}
	}

	c.height = c.currentRock.y

}

func waitForActivity(sub chan Refresh) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

type Model struct {
	Chamber  *Chamber
	Sub      chan Refresh
	step     time.Duration
	tickChan chan<- Refresh
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
	field := make([][]int, 15)
	for i := 0; i < len(field); i++ {
		field[i] = make([]int, m.Chamber.width)
	}

	for _, r := range m.Chamber.rocks {
		for i := 0; i < r.height; i++ {
			for j := 0; j < r.width; j++ {
				x := r.x + j
				y := r.y - i
				field[y][x] = int(r.shape[i][j])
			}
		}
	}

	var s strings.Builder

	for i := len(field) - 1; i >= 0; i-- {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == 1 {
				fmt.Fprint(&s, "#")
			} else {
				fmt.Fprint(&s, ".")
			}
		}
		fmt.Fprint(&s, "\n")
	}

	for j := 0; j < m.Chamber.width; j++ {
		fmt.Fprint(&s, "#")
	}
	fmt.Fprint(&s, "\n")

	return s.String()
}

func main() {
	var input string
	fmt.Scanln(&input)

	refreshChan := make(chan Refresh)
	chamber := CreateChamber(7)

	model := &Model{
		Sub:      refreshChan,
		Chamber:  chamber,
		tickChan: refreshChan,
		step:     500 * time.Millisecond,
	}
	p := tea.NewProgram(model)

	go func(commands string, m *Model) {
		for _, c := range commands {
			m.Chamber.tick(c)
			m.tickChan <- Refresh{}
			time.Sleep(m.step)
		}

		m.tickChan <- Refresh{quit: true}
	}(input, model)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}
