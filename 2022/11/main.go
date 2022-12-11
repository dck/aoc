package main

import (
	"fmt"
	"sort"
	"strings"
)

type Monkey struct {
	Index         int
	Items         []uint64
	Operation     func(uint64) uint64
	TestDivisible int
	TrueTarget    int
	FalseTarget   int
	inspections   int
}

func (m *Monkey) String() string {
	var s strings.Builder

	fmt.Fprintf(&s, "Monkey %d (Inspections: %d)\n", m.Index, m.inspections)
	fmt.Fprintf(&s, "  Starting items: %v\n", m.Items)

	return s.String()
}

func main() {
	// monkeys := []*Monkey{

	// 	{
	// 		Index:         0,
	// 		Items:         []uint64{79, 98},
	// 		Operation:     func(a uint64) uint64 { return a * 19 },
	// 		TestDivisible: 23,
	// 		TrueTarget:    2,
	// 		FalseTarget:   3,
	// 	},
	// 	{
	// 		Index:         1,
	// 		Items:         []uint64{54, 65, 75, 74},
	// 		Operation:     func(a uint64) uint64 { return a + 6 },
	// 		TestDivisible: 19,
	// 		TrueTarget:    2,
	// 		FalseTarget:   0,
	// 	},

	// 	{
	// 		Index:         2,
	// 		Items:         []uint64{79, 60, 97},
	// 		Operation:     func(a uint64) uint64 { return a * a },
	// 		TestDivisible: 13,
	// 		TrueTarget:    1,
	// 		FalseTarget:   3,
	// 	},

	// 	{
	// 		Index:         2,
	// 		Items:         []uint64{74},
	// 		Operation:     func(a uint64) uint64 { return a + 3 },
	// 		TestDivisible: 17,
	// 		TrueTarget:    0,
	// 		FalseTarget:   1,
	// 	},
	// }
	monkeys := []*Monkey{
		{
			Index:         0,
			Items:         []uint64{77, 69, 76, 77, 50, 58},
			Operation:     func(a uint64) uint64 { return a * 11 },
			TestDivisible: 5,
			TrueTarget:    1,
			FalseTarget:   5,
		},
		{
			Index:         1,
			Items:         []uint64{75, 70, 82, 83, 96, 64, 62},
			Operation:     func(a uint64) uint64 { return a + 8 },
			TestDivisible: 17,
			TrueTarget:    5,
			FalseTarget:   6,
		},
		{
			Index:         2,
			Items:         []uint64{53},
			Operation:     func(a uint64) uint64 { return a * 3 },
			TestDivisible: 2,
			TrueTarget:    0,
			FalseTarget:   7,
		},
		{
			Index:         3,
			Items:         []uint64{85, 64, 93, 64, 99},
			Operation:     func(a uint64) uint64 { return a + 4 },
			TestDivisible: 7,
			TrueTarget:    7,
			FalseTarget:   2,
		},

		{
			Index:         4,
			Items:         []uint64{61, 92, 71},
			Operation:     func(a uint64) uint64 { return a * a },
			TestDivisible: 3,
			TrueTarget:    2,
			FalseTarget:   3,
		},

		{
			Index:         5,
			Items:         []uint64{79, 73, 50, 90},
			Operation:     func(a uint64) uint64 { return a + 2 },
			TestDivisible: 11,
			TrueTarget:    4,
			FalseTarget:   6,
		},
		{
			Index:         6,
			Items:         []uint64{50, 89},
			Operation:     func(a uint64) uint64 { return a + 3 },
			TestDivisible: 13,
			TrueTarget:    4,
			FalseTarget:   3,
		},
		{
			Index:         7,
			Items:         []uint64{83, 56, 64, 58, 93, 91, 56, 65},
			Operation:     func(a uint64) uint64 { return a + 5 },
			TestDivisible: 19,
			TrueTarget:    1,
			FalseTarget:   0,
		},
	}

	product := uint64(1)
	for _, m := range monkeys {
		product *= uint64(m.TestDivisible)
	}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {

			for len(m.Items) > 0 {
				item := m.Items[0]
				m.Items = m.Items[1:]
				worryLevel := m.Operation(item) % product

				var newMonkey *Monkey = nil
				if worryLevel%uint64(m.TestDivisible) == 0 {
					newMonkey = monkeys[m.TrueTarget]
				} else {
					newMonkey = monkeys[m.FalseTarget]
				}

				newMonkey.Items = append(newMonkey.Items, worryLevel)
				m.inspections++
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	for _, m := range monkeys {
		fmt.Println(m)
		fmt.Println()
	}

	fmt.Println(monkeys[0].inspections * monkeys[1].inspections)
}
