package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Command struct {
	Instruction string
	Arg         int
}

type VM struct {
	Pixels   [][]rune
	width    int
	height   int
	x        int
	cycle    int
	step     time.Duration
	tickChan chan<- Refresh
}

type Refresh struct {
	quit bool
}

func CreateVM(width int, height int, step time.Duration, tickChan chan<- Refresh) *VM {
	vm := &VM{
		width:    width,
		height:   height,
		x:        1,
		cycle:    0,
		step:     step,
		tickChan: tickChan,
	}

	vm.Pixels = make([][]rune, vm.height)

	for i := 0; i < vm.height; i++ {
		vm.Pixels[i] = make([]rune, vm.width)
		for j := 0; j < vm.width; j++ {
			vm.Pixels[i][j] = '.'
		}
	}

	return vm
}

func (vm *VM) tick() {
	if math.Abs(float64(vm.cycle%vm.width-vm.x)) < 2 {
		i := vm.cycle / vm.width
		j := vm.cycle % vm.width
		vm.Pixels[i][j] = '#'
	}

	vm.cycle++
	vm.tickChan <- Refresh{}
	time.Sleep(vm.step)
}

func (vm *VM) Execute(commands []Command) {
	for _, c := range commands {
		if c.Instruction == "noop" {
			vm.tick()
		}
		if c.Instruction == "addx" {
			vm.tick()
			vm.tick()
			vm.x += c.Arg
		}
	}

	vm.tickChan <- Refresh{quit: true}
}

func waitForActivity(sub chan Refresh) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

type Model struct {
	Vm       *VM
	Sub      chan Refresh
	quitting bool
}

func (m *Model) Init() tea.Cmd {
	return waitForActivity(m.Sub)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
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

	for i := 0; i < m.Vm.height; i++ {
		fmt.Fprintf(&s, "%d [", i+1)
		for j := 0; j < m.Vm.width; j++ {

			if i*m.Vm.width+j == m.Vm.cycle {
				fmt.Fprintf(&s, "\033[1m%c\033[0m", m.Vm.Pixels[i][j])
			} else {
				fmt.Fprintf(&s, "%c", m.Vm.Pixels[i][j])
			}

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

	refreshChan := make(chan Refresh)
	vm := CreateVM(40, 6, 20*time.Millisecond, refreshChan)

	p := tea.NewProgram(&Model{
		Sub: refreshChan,
		Vm:  vm,
	})

	go vm.Execute(commands)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}

}
