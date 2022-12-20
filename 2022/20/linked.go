package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	value    int
	position int
	next     *Number
	prev     *Number
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	encoded := []*Number{}
	i := 0

	var buf *Number
	var root *Number
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)

		num := &Number{value: n, position: i}
		encoded = append(encoded, num)
		if buf == nil {
			buf = num
			root = num
		} else {
			num.prev = buf
			buf.next = num
			buf = num
		}
		i++
	}

	buf.next = root
	root.prev = buf

	p := root
	p.value *= 811589153
	p = p.next

	for p != root {
		p.value *= 811589153
		p = p.next
	}

	for i := 0; i < 10; i++ {
		decrypt(encoded, root)
	}

	index := 1
	p = root.next
	for p != nil && p != root && p.value != 0 {
		index++
		p = p.next
	}
	index = index % len(encoded)
	l := len(encoded)
	a := getByIndex(root, (index+1000)%l).value
	b := getByIndex(root, (index+2000)%l).value
	c := getByIndex(root, (index+3000)%l).value
	fmt.Println(a + b + c)
}

func decrypt(order []*Number, root *Number) *Number {
	for _, n := range order {
		id := n.position
		target := root
		for target.position != id {
			target = target.next
		}

		steps := target.value % (len(order) - 1)

		if steps < 0 {
			for i := 0; i > steps; i-- {
				next := target.next
				prev := target.prev

				next.prev = prev
				prev.next = next

				target.next = prev
				target.prev = prev.prev

				prev.prev.next = target
				prev.prev = target
			}
		} else if steps > 0 {
			for i := 0; i < steps; i++ {
				next := target.next
				prev := target.prev

				prev.next = next
				next.prev = prev

				target.prev = next
				target.next = next.next

				next.next = target
				target.next.prev = target
			}
		} else {
			continue
		}
	}

	return root
}

func getByIndex(root *Number, index int) *Number {
	if root == nil {
		return nil
	}

	p := root
	for i := 0; i < index; i++ {
		p = p.next
	}

	return p
}

func pretty(root *Number) string {
	var s strings.Builder

	if root == nil {
		return ""
	}

	p := root
	fmt.Fprint(&s, "[")
	fmt.Fprintf(&s, "%d ", p.value)
	p = p.next
	for p != nil && p != root {
		fmt.Fprintf(&s, "%d ", p.value)
		p = p.next
	}
	fmt.Fprint(&s, "]")

	return s.String()
}
