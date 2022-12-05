package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node[T any] struct {
	value T
	prev  *Node[T]
}

type Stack[T any] struct {
	top  *Node[T]
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

func (s *Stack[T]) Push(value T) {
	node := &Node[T]{value, s.top}
	s.top = node
	s.size += 1
}

func (s *Stack[T]) Pop() T {
	node := s.top
	s.top = node.prev
	s.size -= 1

	return node.value
}

func (s *Stack[T]) Peek() T {
	if s.size == 0 {
		return *new(T)
	}
	return s.top.value
}

func (s *Stack[T]) PushMany(values []T) {
	for _, v := range values {
		s.Push(v)
	}
}

func (s *Stack[T]) String() string {
	var sb strings.Builder
	curr := s.top
	for curr != nil {
		delimiter := " -> "
		if curr.prev == nil {
			delimiter = ""
		}
		sb.WriteString(fmt.Sprintf("%c%s", curr.value, delimiter))
		curr = curr.prev
	}

	return fmt.Sprintf("Stack(%d): [%s]", s.size, sb.String())
}

const STACK_NUM int = 9

func main() {
	stacks := make([]*Stack[rune], STACK_NUM)
	for i := 0; i < STACK_NUM; i++ {
		stacks[i] = NewStack[rune]()
	}

	stacks[0].PushMany([]rune{'G', 'D', 'V', 'Z', 'J', 'S', 'B'})
	stacks[1].PushMany([]rune{'Z', 'S', 'M', 'G', 'V', 'P'})
	stacks[2].PushMany([]rune{'C', 'L', 'B', 'S', 'W', 'T', 'Q', 'F'})
	stacks[3].PushMany([]rune{'H', 'J', 'G', 'W', 'M', 'R', 'V', 'Q'})
	stacks[4].PushMany([]rune{'C', 'L', 'S', 'N', 'F', 'M', 'D'})
	stacks[5].PushMany([]rune{'R', 'G', 'C', 'D'})
	stacks[6].PushMany([]rune{'H', 'G', 'T', 'R', 'J', 'D', 'S', 'Q'})
	stacks[7].PushMany([]rune{'P', 'F', 'V'})
	stacks[8].PushMany([]rune{'D', 'R', 'S', 'T', 'J'})

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		steps, _ := strconv.ParseInt(fields[1], 10, 64)
		from, _ := strconv.ParseInt(fields[3], 10, 64)
		to, _ := strconv.ParseInt(fields[5], 10, 64)

		temp := NewStack[rune]()

		for i := 0; i < int(steps); i++ {
			temp.Push(stacks[from-1].Pop())
		}

		for temp.size != 0 {
			stacks[to-1].Push(temp.Pop())
		}
	}

	for i := 0; i < STACK_NUM; i++ {
		fmt.Printf("%c", stacks[i].Peek())
	}
	fmt.Printf("\n")
}
