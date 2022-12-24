package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [...]Point{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

type Point struct {
	x int
	y int
}

type Blizzard struct {
	Point
	Direction byte
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

type Board struct {
	grid      [][]byte
	blizzards []*Blizzard
	start     Point
	finish    Point
	width     int
	height    int
}

type Step struct {
	round  int
	player Point
}

func CreateBoard(input []string) *Board {
	b := &Board{}

	b.height = len(input) - 2
	b.width = len(input[0]) - 2

	grid := make([][]byte, b.height)
	blizzards := make([]*Blizzard, 0)

	for i, row := range input {
		if i == 0 {
			// find start
			for j, c := range row {
				if c == '.' {
					b.start = Point{j, i}
					break
				}
			}
			continue
		}

		if i == len(input)-1 {
			// find end
			for j, c := range row {
				if c == '.' {
					b.finish = Point{j, i}
					break
				}
			}
			continue
		}
		grid[i-1] = []byte(input[i][1 : b.width+1])

		for j := 0; j < len(grid[i-1]); j++ {
			if grid[i-1][j] != '.' {
				blizzards = append(blizzards, &Blizzard{Point: Point{j, i - 1}, Direction: grid[i-1][j]})
			}
		}
	}

	b.grid = grid
	b.blizzards = blizzards

	return b
}

func (board *Board) BlizzardMap() map[Point][]*Blizzard {
	bmap := map[Point][]*Blizzard{}

	for i := 0; i < len(board.blizzards); i++ {
		b := board.blizzards[i]

		key := Point{b.x, b.y}

		if _, ok := bmap[key]; ok {
			bmap[key] = append(bmap[key], b)
		} else {
			bmap[key] = []*Blizzard{b}
		}
	}

	return bmap
}

func (board *Board) Move() {
	for i := 0; i < len(board.blizzards); i++ {
		b := board.blizzards[i]

		if b.Direction == '>' {
			b.x++
			if b.x >= board.width {
				b.x = 0
			}
		}
		if b.Direction == '<' {
			b.x--
			if b.x < 0 {
				b.x = board.width - 1
			}
		}
		if b.Direction == '^' {
			b.y--
			if b.y < 0 {
				b.y = board.height - 1
			}
		}
		if b.Direction == 'v' {
			b.y++
			if b.y >= board.height {
				b.y = 0
			}
		}
	}
}

func (board *Board) Stringify(p Point) string {
	var s strings.Builder

	bmap := board.BlizzardMap()

	for i := 0; i < board.height+2; i++ {
		if i == 0 {
			for j := 0; j < board.width+2; j++ {
				if board.start.x == j && board.start.y == i {
					fmt.Fprint(&s, ".")
				} else {
					fmt.Fprint(&s, "#")
				}
			}
			fmt.Fprint(&s, "\n")
			continue
		}

		if i == board.height+1 {
			for j := 0; j < board.width+2; j++ {
				if board.finish.x == j && board.finish.y == i {
					fmt.Fprint(&s, ".")
				} else {
					fmt.Fprint(&s, "#")
				}
			}
			fmt.Fprint(&s, "\n")
			continue
		}

		fmt.Fprint(&s, "#")
		for j := 0; j < board.width; j++ {

			if p.x == j && p.y == i-1 {
				fmt.Fprint(&s, "E")
			} else if b, ok := bmap[Point{j, i - 1}]; ok {
				if len(b) > 1 {
					fmt.Fprintf(&s, "%d", len(b))

				} else {
					fmt.Fprintf(&s, "%c", b[0].Direction)

				}
			} else {
				fmt.Fprint(&s, ".")

			}
		}
		fmt.Fprint(&s, "#\n")
	}

	return s.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := []string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		grid = append(grid, line)
	}

	board := CreateBoard(grid)

	currentRound := 0
	steps := []Step{
		{player: Point{board.start.x - 1, board.start.y - 1}, round: currentRound},
	}
	totalSteps := 0
	found := false

	for len(steps) > 0 && !found {
		currentRoundSteps := map[Point]Step{}
		for len(steps) > 0 && steps[0].round == currentRound {
			step := steps[0]
			steps = steps[1:]
			currentRoundSteps[step.player] = step
		}

		if currentRound == 7 {
			for _, ste := range currentRoundSteps {
				fmt.Println(currentRound)
				fmt.Println(ste.player)
				fmt.Println(board.Stringify(ste.player))
			}
		}

		board.Move()
		bmap := board.BlizzardMap()

		for _, step := range currentRoundSteps {
			player := step.player

			if player.x+1 == board.finish.x && player.y+2 == board.finish.y {
				totalSteps = step.round
				found = true
				break
			}

			for _, d := range directions {
				newX := player.x + d.x
				newY := player.y + d.y

				if newX < 0 || newY < 0 || newX >= board.width || newY >= board.height {
					continue
				}

				if _, ok := bmap[Point{newX, newY}]; ok {
					continue
				}

				steps = append(steps, Step{player: Point{newX, newY}, round: step.round + 1})
			}

			_, overlapping := bmap[player]
			if !overlapping {
				steps = append(steps, Step{player: player, round: step.round + 1})
			}
		}
		currentRound++
	}

	fmt.Println(totalSteps + 1)
}
