package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	LOCK = iota
	KEY
)

const SIZE = 5

type Seq [SIZE]uint8

type Device struct {
	Type uint8
	Seq  Seq
}

func main() {
	var input string
	var err error
	reader := bufio.NewReader(os.Stdin)
	for {
		var line string
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		input += line
	}

	blocks := strings.Split(input, "\n\n")

	devices := make([]Device, 0, len(blocks))
	for _, block := range blocks {
		grid := strings.Split(block, "\n")
		var kind uint8
		var seq Seq
		if grid[0][0] == '#' {
			kind = LOCK
		} else {
			kind = KEY
		}

		for i := 1; i <= SIZE; i++ {
			for j := 0; j < SIZE; j++ {
				if grid[i][j] == '#' {
					seq[j]++
				}
			}
		}

		devices = append(devices, Device{kind, seq})
	}

	fmt.Println("Part 1:", part1(devices))
}

func part1(devices []Device) int {
	res := 0

	sort.Slice(devices, func(i, j int) bool {
		return devices[i].Type < devices[j].Type
	})

	atLeastKeys := make([]Seq, 0, len(devices))
	for _, device := range devices {
		if device.Type == LOCK {
			inverted := Seq{}
			for i := 0; i < SIZE; i++ {
				inverted[i] = SIZE - device.Seq[i]
			}
			atLeastKeys = append(atLeastKeys, inverted)
		} else {
			res += findLess(atLeastKeys, device.Seq)
		}
	}

	return res
}

func findLess(ary []Seq, seq Seq) int {
	res := 0
	for _, s := range ary {
		found := false

		for i := 0; i < SIZE; i++ {
			if seq[i] > s[i] {
				found = true
				break
			}
		}

		if !found {
			res++
		}
	}
	return res
}
