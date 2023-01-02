package main

import (
	"container/heap"
	"fmt"
	"github.com/mbark/advent-of-code-2022/maps"
	"github.com/mbark/advent-of-code-2022/maths"
	"math"
	"strconv"
)

func main() {
	m := maps.New[blizzard](Input, func(x, y int, b byte) blizzard {
		switch b {
		case '>':
			return []maps.Direction{maps.Right}
		case '<':
			return []maps.Direction{maps.Left}
		case 'v':
			return []maps.Direction{maps.Down}
		case '^':
			return []maps.Direction{maps.Up}

		default:
			return nil
		}
	})

	states = []maps.Map[blizzard]{m}
	fmt.Printf("first: %d\n", first(m))
	fmt.Printf("second: %d\n", second(m))
}

func first(m maps.Map[blizzard]) int {
	start, goal := maps.Coordinate{X: 1, Y: 0}, maps.Coordinate{Y: m.Rows - 1, X: m.Columns - 2}
	return traverse(start, goal, 0)
}

func second(m maps.Map[blizzard]) int {
	start, goal := maps.Coordinate{X: 1, Y: 0}, maps.Coordinate{Y: m.Rows - 1, X: m.Columns - 2}
	toEnd := traverse(start, goal, 0)
	backAgain := traverse(goal, start, toEnd)
	s := traverse(start, goal, backAgain)
	return s
}

type key struct {
	x, y, time int
}

func memokey(c maps.Coordinate, time, lcm int) key {
	return key{x: c.X, y: c.Y, time: time % lcm}
}

var states []maps.Map[blizzard]

func getState(time int) maps.Map[blizzard] {
	if time < len(states) {
		return states[time]
	}

	for i := len(states); i <= time; i++ {
		states = append(states, nextState(states[i-1]))
	}

	return states[time]
}

type item struct {
	c    maps.Coordinate
	time int
}

func traverse(start, goal maps.Coordinate, startTime int) int {
	lcm := maths.LCM(states[0].Rows, states[0].Columns)
	visited := make(map[key]bool)

	var pq maps.PriorityQueue[item]
	heap.Init(&pq)
	heap.Push(&pq, &maps.Item[item]{Value: item{c: start, time: startTime}, Priority: 0})

	doneAt := math.MaxInt
	pushIf := func(c maps.Coordinate, state maps.Map[blizzard], time int) {
		if b := state.At(c); b != nil {
			return
		}
		if time+c.ManhattanDistance(goal) > doneAt {
			return
		}
		if visited[memokey(c, time, lcm)] {
			return
		}

		heap.Push(&pq, &maps.Item[item]{Value: item{c: c, time: time}, Priority: c.ManhattanDistance(goal)})
	}

	for len(pq) > 0 {
		n := heap.Pop(&pq).(*maps.Item[item])
		val := n.Value
		at := n.Value.c
		visited[memokey(at, val.time, lcm)] = true

		time := val.time + 1
		if time > doneAt {
			continue
		}

		state := getState(time % lcm)
		for _, a := range at.Adjacent() {
			if a == goal {
				doneAt = maths.MinInt(doneAt, time)
				continue
			}

			if a.X <= 0 || a.Y <= 0 || a.X >= state.Columns-1 || a.Y >= state.Rows-1 {
				continue
			}

			pushIf(a, state, time)
		}

		pushIf(at, state, time)
	}

	return doneAt
}

func printMap(at maps.Coordinate, m maps.Map[blizzard]) {
	stringf := func(c maps.Coordinate, b blizzard) string {
		if c == at {
			return "E"
		}
		if c.Y == 0 && c.X == 1 {
			return "."
		}
		if c.Y == m.Rows-1 && c.X == m.Columns-2 {
			return "."
		}
		if c.Y == 0 {
			return "#"
		}
		if c.X == 0 {
			return "#"
		}
		if c.Y == m.Rows-1 {
			return "#"
		}
		if c.X == m.Columns-1 {
			return "#"
		}

		return b.String()
	}

	fmt.Println(m.Stringf(stringf))
}

func nextState(curr maps.Map[blizzard]) maps.Map[blizzard] {
	coords := make(map[maps.Coordinate]blizzard)
	rows := curr.Rows - 1
	cols := curr.Columns - 1
	for _, c := range curr.Coordinates() {
		for _, d := range curr.At(c) {
			coord := d.Apply(c)

			if coord.Y == rows {
				coord.Y = 1
			}
			if coord.X == cols {
				coord.X = 1
			}
			if coord.Y == 0 {
				coord.Y = rows - 1
			}
			if coord.X == 0 {
				coord.X = cols - 1
			}

			coords[coord] = append(coords[coord], d)
		}
	}

	for _, c := range curr.Coordinates() {
		if curr.At(c) == nil && coords[c] == nil {
			coords[c] = nil
		}
	}

	return maps.MapFromCoordinates[blizzard](coords)
}

type blizzard []maps.Direction

func (b blizzard) String() string {
	if b == nil {
		return "."
	}

	if len(b) == 1 {
		return b[0].String()
	}

	return strconv.Itoa(len(b))
}

const testInput = `
#E######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

const Input = `
#.######################################################################################################################################################
#>.^v.>^v>><>>^>>v>v^>.<>^.>v^v>..<<><>>>v.><^<.>^.<^>^v^>>v<^>><.^<^.<v>v><><^^<^>>^<<<>.>.>^^^<^<^>.>^v.>.>^>^v<<^^<>><vv<><>>>v^.>^<vvv>.^>v.>.v<v.<#
#<^^^><<vv<vv^^><>><>vv.vv<>vvvv<vv.>^>^>v^vvv>.<^>v<>>.<><<><>v<v^^<^.v^^><<><^><<.v.vv<<><.<<<<v^^>><^<><^<<<<.>v<^v<v>^>>>><.v>^>.>^<^>>v<v<^vv<v>v<#
#>>>>^.>.^v>v<.^<^vv^^^<>.>>.>><<.v<v^><^><.<<<><.^v^^v<><v^v><vv<<v<..<<^<.<>^<^v.^>^<v^v.^<v>v<><>>>^v<v.^^<><>><>v<^v>>>^v<v^>v<^vvv><v>^><<><^v>>^>#
#><<>^>v>.<^^v><v.<>v^<<<>^^v>^.v<<>><>v.v>v^.v>>>^<^><.>>v<<<>vv^.>^>^<><^.>.v^^<v<.^>^><>^<v>vv^v>v^<>><<vv>>>v^^<^v>>>v.<<>.v<><^^v<.^>v>.<^^><vv.^>#
#<.v^vv^vv>><v>^>>^.vvv^^<v<.v^>><<>v<v<<^<<<vv.>><>v.^>vvv<v<v>^<>>v><.>^.>.^<<<.^vvv>^..^>>.v<v.^v^<^^<<^<^>^v>^v<vvv.v^^v<<<>^v<<<v><>v^<<^><.v>v<v>#
#<^v^><>v^<.>vv>.^<<<vv^^<><^<^.^^<.^vv^>v^^.^<^>^.^<><><^<<>><vv<v^<>v<><>v^^v<.<v^<v>vvv>^>>^<^><>>^vv<^>><v><^.<<v><v.>^v^v>>>^<vv>v<^<><>.^^>.v><<>#
#<.><^v>^<.>>v<v>v<^v^v.^>v<.v>.<vv>^>v<v<<>^^.<>.<.vvv<^>v^.^<^v>^.>.v<.v>><<<^vv^.^<<^<^v<v^^^>.><>v.^^.^><<.<vvvv<v><<^vv.<^^^^>^<.<><<>>^>v>.^^v.<<#
#<<^.>v<><>.^>>^v<<^>vv<^v>><.<v^v^<<><v><^^.^^^>^.><^>v^v<<<<>>^>v>^v<^<>..v<.<v<><>v^>v^^v..<>^<><>^v<<v<<><<.v>...>^^v>.>v^.^>vv>^v<>v><<^.>^^..<<^<#
#><v^<^..>>v<v.<>vvv>v<v^>>>.v^<^v<>.>^<.<<^<>.<..>.^.>^^>.^.^^>v.v^>v<^>^v.<v^>v^^^vv><^>^>>>v><v<>>>..v.<^<.<<<>>><^<<v>>^<.>^^^<><^<.<^>^<v.^>vv>vv>#
#<><>v^..>v>^^^^<.v>v>..<>><^><<^v.v.<v^<v><><<<v>>>^^>>v<<>^^<^^^>v.<<>v^<v><v^<<v>>^v>>v><v^^>vv^<^v<^v^.><>vv>>>.>>v^<v^v>><^^v.^<v.><v^^<<<v^>>^v^>#
#>v><>><<^.>v^.v<<<v><^vv<.v^>v<.>vv>v>>^<^<<><>v>v<v..v<^<.vvv<^>^^v<^.>.^^^>^^v.^.>^>>>^><v<.v<^^^^.v.vv<<vv<.>v<^>v^<<vv^v><v>v^>>v.^.^<<.vv>^^^<^v<#
#<v^<>>^>>^..vv<^<<<v.v.^vv<.<^^<v^<.>^<.^.<<v^^^<^><>.>>>><^<>^>^v>><v^>v.<.^<v>^^>><><.>^v^>^<>^<^^.<^>^^^^^^vv<>vv^v^v>^v^<v<vv><<vv<^^^.>>^><<<.^>>#
#<.<v^.>>><vv^.<^^^>>>.><v<v>^<>>.^^vv>>^<<<^v><^^<v<>>^^^^<v^<>^^>><^>v^<<>>v<^.>vvvv>>^<.^>.v<^^<<v>>.vv.v^v<>>>><^v^>>>.v.^^v<<v>v^^>^v>^v<><.v><^v.#
#><<^v><<<.v<^^><^v<^v.<<<.^^.v>vvvv>^>><^^v<>.<>><<.v^>..^<>.><^>^>^v^^>v<<vv>^>vvv.>^^>^^<<>v>v^>^^^<^v>.v>^<><^<.v<^><<>^.^>>><<v^>v^v^v<^v>>^^^><^<#
#<^>v<>.>>^.v>^<v><>.v.^^v>^<..v.^>^>.v>>><vv^<>^^>v><>vv.<v<v^vv<.^vv<><v>>>^v><v^vvvvv.<vv>>^v<<.v>v^>v><<<^.<<<<>vv>^>>^<>^^<vvv^<vvvvvv><^^<<<>^^<>#
#>>.>>.^>..<^<^.<<v<><>..^^<..v<.><^^<v.v^v><>^vv<v<^v<^>.v<vvvv><<<>^>^v^>v<>><vv^v<>>>>>v<v<>^.>><v<v.>><v<v^^<v.<><.<>^^^^v<<^>>.><<.><.v^^..>>>.v>>#
#>v<v^^<^vvv.>^<>>v^v.vv^>v^vvvvv<>.v^.v<.><>^>>^<<v<^^v>v><^.>>^><>v<vv>v<<^>^>^.<v>vv<<>^^>^<<>v><v>.<<^.<>><<^<<.^<^<<<.^>.<.vv>v^v<^>vvv>^v<<<<v<<<#
#<^<>v^.^<><>^<>^>v<v.^v..v..vv^.<.^^>^>^><.vv>>.<^..^^^<^<.^^<^^v^<^.<<<<<v<^<<><v>^v.<<>.^>><^>^v>><<v<>^v^<<^>>v<<^.v<<^vv>>v>^.v^^^<><vv^^<<>.v.<><#
#..vv<v><^vv>.v<v<^v><vvv>.v>^v^>>vv^.^>v.v^^<<>^^<>^<^^>.<^v<v>^>^<v<<v<>v<^^^>v>><><<><.<<><^<v<<>v^^>>v><>^v^.v>^^v<^><><^v<^<vv.vvv><><<>v<<><>v<.>#
#>>.<^^>><<<^>v<^^>>^<<>><<.<^^^<^^^<^<v<^<<^><<v<v<^^>vv<<^v>.^<^^>><>^^^.>v>v^<^>>v><<^<^^^>v><<>^^v><>v..<.^^>..vv.<v>>>^<<v^<v^^v>>>>.>vv.>>^v>^v<.#
######################################################################################################################################################.#
`
