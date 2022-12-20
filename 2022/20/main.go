package old

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	position int
	value    int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	encoded := []Number{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)

		encoded = append(encoded, Number{value: n, position: i})
		i++
	}

	decoded := make([]Number, len(encoded))
	copy(decoded, encoded)

	for i, n := range encoded {
		currentIndex := -1
		for j, m := range decoded {
			if n.position == m.position {
				currentIndex = j
			}
		}

		if math.Abs(float64(n.value%len(decoded))) == float64(len(decoded)) || math.Abs(float64(n.value%len(decoded))) == 0 {
			continue
		}

		newIndex := (currentIndex + n.value) % len(decoded)

		if newIndex < 0 {
			newIndex = len(decoded) + newIndex
		} else if newIndex == 0 && n.value < 0 {
			newIndex = len(decoded) - 1
		} else if newIndex > 0 && currentIndex+n.value >= len(decoded) && newIndex != currentIndex {
			newIndex++
		}

		shiftItem(decoded, currentIndex, newIndex)
		fmt.Printf("%d: Value: %d, Index: %d. %s\n", i, n.value, newIndex, pretty(decoded))
	}

	index := -1
	for i, n := range decoded {
		if n.value == 0 {
			index = i
		}
	}

	l := len(decoded)
	a := decoded[(index+1000)%l].value
	b := decoded[(index+2000)%l].value
	c := decoded[(index+3000)%l].value
	fmt.Println(a + b + c)
}

func removeItem[T comparable](array []T, index int) {
	copy(array[index:], array[index+1:])
	array = array[:len(array)-1]
}

func insertItem[T comparable](array []T, index int, item T) {
	array = append(array[:index+1], array[index:]...)
	array[index] = item
}

func shiftItem[T comparable](array []T, from int, to int) {
	val := array[from]
	to = to % len(array)

	sign := 0
	if from == to {
		return
	} else if from < to {
		sign = 1
	} else {
		sign = -1
	}

	for from != to {
		if from+sign < 0 {
			from = len(array) - 1
		}
		if from+sign >= len(array) {
			from = 0
		}

		array[from] = array[from+sign]
		from += sign
	}

	array[from] = val
}

func pretty(numbers []Number) string {
	var s strings.Builder

	fmt.Fprint(&s, "[")
	for _, n := range numbers {
		fmt.Fprintf(&s, "%d ", n.value)

	}
	fmt.Fprint(&s, "]")

	return s.String()
}
