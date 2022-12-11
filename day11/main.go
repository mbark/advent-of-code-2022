package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const testInput = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

var testMonkeys = []monkey{
	{
		items:       []int{79, 98},
		inspect:     func(old int) int { return old * 19 },
		divisibleBy: 23,
		ifTrue:      2,
		ifFalse:     3,
	},
	{
		items:       []int{54, 65, 75, 74},
		inspect:     func(old int) int { return old + 6 },
		divisibleBy: 19,
		ifTrue:      2,
		ifFalse:     0,
	},
	{
		items:       []int{79, 60, 97},
		inspect:     func(old int) int { return old * old },
		divisibleBy: 13,
		ifTrue:      1,
		ifFalse:     3,
	},
	{
		items:       []int{74},
		inspect:     func(old int) int { return old + 3 },
		divisibleBy: 17,
		ifTrue:      0,
		ifFalse:     1,
	},
}

func main() {
	monkeys := realMonkeys
	fmt.Printf("first: %d\n", first(monkeyCopy(monkeys)))
	fmt.Printf("second: %d\n", second(monkeyCopy(monkeys)))
}

type monkey struct {
	items       []int
	inspect     func(i int) int
	divisibleBy int
	ifTrue      int
	ifFalse     int
}

func (m monkey) copy() monkey {
	var items []int
	for _, i := range m.items {
		items = append(items, i)
	}

	return monkey{
		items:       items,
		inspect:     m.inspect,
		divisibleBy: m.divisibleBy,
		ifTrue:      m.ifTrue,
		ifFalse:     m.ifFalse,
	}
}

func monkeyCopy(monkeys []monkey) []monkey {
	var out []monkey
	for _, m := range monkeys {
		out = append(out, m.copy())
	}

	return out
}

func (m monkey) String() string {
	var s []string
	for _, item := range m.items {
		s = append(s, strconv.Itoa(item))
	}

	return strings.Join(s, ", ")
}

func first(monkeys []monkey) int {
	inspections := make(map[int]int)

	for round := 0; round < 20; round++ {
		for mi, m := range monkeys {
			for _, item := range m.items {
				inspections[mi] += 1
				item = m.inspect(item)
				item = item / 3

				var to int
				if item%m.divisibleBy == 0 {
					to = m.ifTrue
				} else {
					to = m.ifFalse
				}

				monkeys[to].items = append(monkeys[to].items, item)
			}

			monkeys[mi].items = nil
		}
	}

	var counts []int
	for _, i := range inspections {
		counts = append(counts, i)
	}
	sort.Ints(counts)
	max1, max2 := counts[len(counts)-1], counts[len(counts)-2]

	return max1 * max2
}

func second(monkeys []monkey) int {
	prime := 1
	for _, m := range monkeys {
		prime *= m.divisibleBy
	}

	inspections := make(map[int]int)
	for round := 0; round < 10000; round++ {
		for mi, m := range monkeys {
			for _, item := range m.items {
				inspections[mi] += 1
				item = m.inspect(item)
				item = item % prime

				var to int
				if item%m.divisibleBy == 0 {
					to = m.ifTrue
				} else {
					to = m.ifFalse
				}

				monkeys[to].items = append(monkeys[to].items, item)
			}

			monkeys[mi].items = nil
		}
	}

	var counts []int
	for _, i := range inspections {
		counts = append(counts, i)
	}
	sort.Ints(counts)
	max1, max2 := counts[len(counts)-1], counts[len(counts)-2]

	return max1 * max2
}

var realMonkeys = []monkey{
	{
		items:       []int{54, 89, 94},
		inspect:     func(old int) int { return old * 7 },
		divisibleBy: 17,
		ifTrue:      5,
		ifFalse:     3,
	},
	{
		items:       []int{66, 71},
		inspect:     func(old int) int { return old + 4 },
		divisibleBy: 3,
		ifTrue:      0,
		ifFalse:     3,
	},
	{
		items:       []int{76, 55, 80, 55, 55, 96, 78},
		inspect:     func(old int) int { return old + 2 },
		divisibleBy: 5,
		ifTrue:      7,
		ifFalse:     4,
	},
	{
		items:       []int{93, 69, 76, 66, 89, 54, 59, 94},
		inspect:     func(old int) int { return old + 7 },
		divisibleBy: 7,
		ifTrue:      5,
		ifFalse:     2,
	},
	{
		items:       []int{80, 54, 58, 75, 99},
		inspect:     func(old int) int { return old * 17 },
		divisibleBy: 11,
		ifTrue:      1,
		ifFalse:     6,
	},
	{
		items:       []int{69, 70, 85, 83},
		inspect:     func(old int) int { return old + 8 },
		divisibleBy: 19,
		ifTrue:      2,
		ifFalse:     7,
	},
	{
		items:       []int{89},
		inspect:     func(old int) int { return old + 6 },
		divisibleBy: 2,
		ifTrue:      0,
		ifFalse:     1,
	},
	{
		items:       []int{62, 80, 58, 57, 93, 56},
		inspect:     func(old int) int { return old * old },
		divisibleBy: 13,
		ifTrue:      6,
		ifFalse:     4,
	},
}
