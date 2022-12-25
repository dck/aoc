package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := uint64(0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		total += SnafuToDec(line)
	}

	fmt.Println("Total: ", total)
	fmt.Println("SNAFU: ", DecToSnafu(total))
}

func SnafuToDec(s string) uint64 {
	total := uint64(0)

	m := 1
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]

		if c == '=' {
			total -= uint64(2 * m)
		} else if c == '-' {
			total -= uint64(m)
		} else {
			n, _ := strconv.Atoi(string(c))
			total += uint64(n * m)
		}

		m = m * 5
	}

	return total
}

func DecToSnafu(n uint64) string {
	var s strings.Builder

	for n > 0 {
		rem := n % 5
		n = n / 5

		if rem == 4 {
			fmt.Fprint(&s, "-")
			n += 1
		} else if rem == 3 {
			fmt.Fprint(&s, "=")
			n += 1
		} else {
			fmt.Fprintf(&s, "%d", rem)
		}
	}

	return reversed(s.String())
}

func reversed(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
