package main

import (
	"bufio"
	"fmt"
	"os"
)

const SLICE = 14

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		l := len(line)

		if l < SLICE {
			total += 0
			continue
		}

		for i := SLICE - 1; i < l; i++ {
			h := map[byte]bool{}

			for j := i; j > i-SLICE; j-- {
				h[line[j]] = true
			}

			if len(h) == SLICE {
				total += i + 1
				break
			}
		}
	}

	fmt.Println(total)
}
