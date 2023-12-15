package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Len struct {
	Label string
	Focus int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := strings.Split(line, ",")

		fmt.Println("Focal:", focal(fields))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func focal(fields []string) int {
	var boxes [256][]Len

	for _, field := range fields {
		var focusStr string

		if field[len(field)-1] == '-' {
			field = field[:len(field)-1]
		} else {
			field, focusStr = strings.Split(field, "=")[0], strings.Split(field, "=")[1]
		}

		boxId := hash(field)

		if focusStr != "" {
			focus, _ := strconv.Atoi(focusStr)
			idx := slices.IndexFunc(boxes[boxId], func(c Len) bool { return c.Label == field })

			if idx == -1 {
				boxes[boxId] = append(boxes[boxId], Len{Label: field, Focus: focus})
			} else {
				boxes[boxId][idx].Focus = focus
			}
		} else {
			idx := slices.IndexFunc(boxes[boxId], func(c Len) bool { return c.Label == field })
			if idx != -1 {
				boxes[boxId] = append(boxes[boxId][:idx], boxes[boxId][idx+1:]...)
			}
		}
	}

	total := 0
	for i, box := range boxes {
		if len(box) == 0 {
			continue
		}
		for j, len := range box {
			total += (i + 1) * (j + 1) * len.Focus
		}
	}

	return total
}

func hash(s string) uint8 {
	var h int
	for i := 0; i < len(s); i++ {
		h += int(s[i])
		h *= 17
		h = h % 256
	}
	return uint8(h)
}
