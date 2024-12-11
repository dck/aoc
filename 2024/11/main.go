package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head  *Node
	tail  *Node
	count int
}

type Entry struct {
	number int
	level  int
}

func (l *List) add(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	l.count++
}

func main() {
	nums := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		for _, n := range strings.Fields(line) {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}
	}
	fmt.Println("Part 1:", part1(nums))
	fmt.Println("Part 2:", part2(nums))
}

func part1(nums []int) int {
	list := List{}
	for _, n := range nums {
		list.add(n)
	}

	for range 25 {
		curr := list.head
		for curr != nil {
			if curr.value == 0 {
				curr.value = 1
			} else if digits := numberOfDigits(curr.value); digits%2 == 0 {
				a := curr.value / int(math.Pow(10, float64(digits/2)))
				b := curr.value % int(math.Pow(10, float64(digits/2)))

				next := &Node{value: b}
				next.next = curr.next
				curr.value = a
				curr.next = next
				list.count++
				curr = next
			} else {
				curr.value *= 2024
			}
			curr = curr.next
		}
	}

	return list.count
}

func part2(nums []int) int {
	total := 0
	cache := make(map[Entry]int)
	for _, n := range nums {
		total += blink(n, 75, cache)
	}
	return total
}

func numberOfDigits(n int) int {
	digits := 0
	for n > 0 {
		n /= 10
		digits++
	}
	return digits
}

func blink(n int, level int, cache map[Entry]int) int {
	if cached, ok := cache[Entry{n, level}]; ok {
		return cached
	}

	if level == 0 {
		return 1
	}

	if n == 0 {
		cache[Entry{n, level}] = blink(1, level-1, cache)
	} else if digits := numberOfDigits(n); digits%2 == 0 {
		a := n / int(math.Pow(10, float64(digits/2)))
		b := n % int(math.Pow(10, float64(digits/2)))
		cache[Entry{n, level}] = blink(a, level-1, cache) + blink(b, level-1, cache)
	} else {
		cache[Entry{n, level}] = blink(n*2024, level-1, cache)
	}

	return cache[Entry{n, level}]
}
