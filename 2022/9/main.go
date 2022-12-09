package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Point struct {
	X int
	Y int
}

type Field struct {
	Head    *Point
	Tail    *Point
	body    []Point
	length  int
	visited map[Point]bool
}

func initField(length int) *Field {
	f := &Field{visited: map[Point]bool{}, body: make([]Point, length), length: length}
	for i := 0; i < length; i++ {
		f.body[i] = Point{}
	}

	f.Head = &f.body[0]
	f.Tail = &f.body[length-1]

	return f
}

func (f *Field) Move(direction string) {
	switch direction {
	case "up", "U":
		f.Head.Y--
	case "down", "D":
		f.Head.Y++
	case "left", "L":
		f.Head.X--
	case "right", "R":
		f.Head.X++
	}

	for i := 1; i < f.length; i++ {
		d_x := f.body[i-1].X - f.body[i].X
		d_y := f.body[i-1].Y - f.body[i].Y

		if math.Abs(float64(d_x)) >= 2 || math.Abs(float64(d_y)) >= 2 {
			if d_x > 0 {
				f.body[i].X += 1
			} else if d_x < 0 {
				f.body[i].X += -1
			} else {
				f.body[i].X += 0
			}
			if d_y > 0 {
				f.body[i].Y += 1
			} else if d_y < 0 {
				f.body[i].Y += -1
			} else {
				f.body[i].Y += 0
			}
		}
	}

	f.visited[*f.Tail] = true
}

func (f *Field) Visited() int {
	return len(f.visited)
}

func (f *Field) Init() tea.Cmd {
	return nil
}

func (f *Field) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return f, tea.Quit
		case "up", "down", "left", "right":
			f.Move(msg.String())
		}
	}

	return f, nil
}

func (f *Field) View() string {
	var s strings.Builder

	snake := map[Point]string{}
	for i := 1; i < f.length-1; i++ {
		snake[f.body[i]] = "*"
	}

	snake[*f.Head] = "H"
	snake[*f.Tail] = "T"

	for y := -10; y <= 10; y++ {
		for x := -25; x <= 25; x++ {

			p := Point{X: x, Y: y}

			sign, ok := snake[p]

			if ok {
				fmt.Fprint(&s, sign)
			} else {
				fmt.Fprint(&s, ".")
			}
		}

		fmt.Fprint(&s, "\n")
	}

	return s.String()
}

func main() {
	// field := initField(10)

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	line := scanner.Text()

	// 	parts := strings.Fields(line)

	// 	direction := parts[0]
	// 	steps, _ := strconv.Atoi(parts[1])

	// 	for i := 0; i < steps; i++ {
	// 		field.Move(direction)
	// 	}
	// }
	// fmt.Println(field.Visited())

	p := tea.NewProgram(initField(10))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occured: %v", err)
		os.Exit(1)
	}
}
