package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/maps"
	"github.com/mbark/advent-of-code-2022/maths"
	"github.com/mbark/advent-of-code-2022/util"
	"strconv"
	"strings"
)

const testInput0 = `noop
addx 3
addx -5`

func main() {
	var instructions []instruction
	for _, l := range util.ReadInput(Input, "\n") {
		if l == "noop" {
			instructions = append(instructions, instruction{noop: true})
			continue
		}

		s := strings.Split(l, " ")
		instructions = append(instructions,
			instruction{add: util.ParseInt[int](s[1]), start: true},
			instruction{add: util.ParseInt[int](s[1])},
		)
	}

	fmt.Printf("first: %d\n", first(instructions))
	fmt.Printf("second:\n\n%s\n", second(instructions))
}

type instruction struct {
	start bool
	add   int
	noop  bool
}

func (i instruction) String() string {
	if i.noop {
		return "noop"
	}
	if i.start {
		return "addx " + strconv.Itoa(i.add) + " (start)"
	}

	return fmt.Sprintf("addx %d", i.add)
}

func (i instruction) do(x int) int {
	switch {
	case i.noop:
	case i.start:
	default:
		return x + i.add
	}

	return x
}

func first(instructions []instruction) int {
	x := 1
	var score int
	for c := 0; c <= 240; c++ {
		if c >= len(instructions) {
			break
		}

		cycle := c + 1
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			score += cycle * x
		}

		x = instructions[c].do(x)
	}

	return score
}

type pixel bool

func (p pixel) String() string {
	if p {
		return "#"
	} else {
		return " "
	}
}

func second(instructions []instruction) string {
	m := maps.Map[pixel]{}
	m = m.WithPadding(0, 40, 6, 0)

	x := 1
	for c := 0; c <= 240; c++ {
		if c >= len(instructions) {
			break
		}

		row := c / 40
		crt := c % 40

		if maths.AbsInt(x-crt) <= 1 {
			m.Set(maps.Coordinate{Y: row, X: crt}, true)
		}

		x = instructions[c].do(x)
	}

	return m.String()
}

const testInput = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`

const Input = `noop
noop
addx 5
addx 29
addx -28
addx 5
addx -1
noop
noop
addx 5
addx 12
addx -6
noop
addx 4
addx -1
addx 1
addx 5
addx -31
addx 32
addx 4
addx 1
noop
addx -38
addx 5
addx 2
addx 3
addx -2
addx 2
noop
addx 3
addx 2
addx 5
addx 2
addx 3
noop
addx 2
addx 3
noop
addx 2
addx -32
addx 33
addx -20
addx 27
addx -39
addx 1
noop
addx 5
addx 3
noop
addx 2
addx 5
noop
noop
addx -2
addx 5
addx 2
addx -16
addx 21
addx -1
addx 1
noop
addx 3
addx 5
addx -22
addx 26
addx -39
noop
addx 5
addx -2
addx 2
addx 5
addx 2
addx 23
noop
addx -18
addx 1
noop
noop
addx 2
noop
noop
addx 7
addx 3
noop
addx 2
addx -27
addx 28
addx 5
addx -11
addx -27
noop
noop
addx 3
addx 2
addx 5
addx 2
addx 27
addx -26
addx 2
addx 5
addx 2
addx 4
addx -3
addx 2
addx 5
addx 2
addx 3
addx -2
addx 2
noop
addx -33
noop
noop
noop
noop
addx 31
addx -26
addx 6
noop
noop
addx -1
noop
addx 3
addx 5
addx 3
noop
addx -1
addx 5
addx 1
addx -12
addx 17
addx -1
addx 5
noop
noop
addx 1
noop
noop
`
