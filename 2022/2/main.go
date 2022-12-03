package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

var pairs = map[string]int{
	"A X": 3 + 0,
	"A Y": 1 + 3,
	"A Z": 2 + 6,
	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,
	"C X": 2 + 0,
	"C Y": 3 + 3,
	"C Z": 1 + 6,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		score, ok := pairs[line]
		if !ok {
			log.Fatal(errors.New("Impossible pair combination"))
		}

		total += score
	}

	fmt.Println(total)
}
