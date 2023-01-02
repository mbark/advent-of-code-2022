package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/maps"
	"github.com/mbark/advent-of-code-2022/maths"
	"math"
	"strings"
)

func main() {
	m := maps.New[terrain](Input, func(x, y int, b byte) terrain {
		switch b {
		case '.':
			return terrain{}

		case '#':
			return terrain{isElf: true}
		}

		return terrain{}
	})

	elves := make(map[maps.Coordinate]*terrain)
	for _, c := range m.Coordinates() {
		if e := m.At(c); e.isElf {
			elves[c] = &e
		}
	}

	elves2 := make(map[maps.Coordinate]*terrain)
	for k, v := range elves {
		elves2[k] = v
	}

	fmt.Printf("first: %d\n", first(elves))
	fmt.Printf("second: %d\n", second(elves2))
}

type proposition struct {
	move  maps.Direction
	check []maps.Direction
}

func (p proposition) String() string {
	var s []string
	for _, d := range p.check {
		s = append(s, d.String())
	}
	return fmt.Sprintf("move to %s and check [%s]", p.move, strings.Join(s, ", "))
}

func first(elves map[maps.Coordinate]*terrain) int {
	propositions := []proposition{
		{move: maps.North, check: []maps.Direction{maps.North, maps.NorthEast, maps.NorthWest}},
		{move: maps.South, check: []maps.Direction{maps.South, maps.SouthEast, maps.SouthWest}},
		{move: maps.West, check: []maps.Direction{maps.West, maps.SouthWest, maps.NorthWest}},
		{move: maps.East, check: []maps.Direction{maps.East, maps.SouthEast, maps.NorthEast}},
	}

	elfAt := make(map[*terrain]maps.Coordinate, len(elves))
	for c, e := range elves {
		elfAt[e] = c
	}

	var pidx int
	for i := 0; i < 10; i++ {
		proposes := make(map[maps.Coordinate][]*terrain)

		for c, e := range elves {
			shouldMove := false
			for _, c := range c.Surrounding() {
				if _, ok := elves[c]; ok {
					shouldMove = true
					break
				}
			}

			if !shouldMove {
				continue
			}

			moveTo := c
			for j := 0; j < len(propositions) && moveTo == c; j++ {
				p := propositions[(pidx+j)%len(propositions)]

				proposeMove := true
				for _, d := range p.check {
					if _, ok := elves[d.Apply(c)]; ok {
						proposeMove = false
					}
				}

				if proposeMove {
					moveTo = p.move.Apply(c)
				}
			}

			proposes[moveTo] = append(proposes[moveTo], e)
		}

		for c, es := range proposes {
			if len(es) > 1 {
				continue
			}

			e := es[0]
			delete(elves, elfAt[e])
			elfAt[e] = c
			elves[c] = e
		}

		pidx += 1
		pidx %= len(propositions)
	}

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := 0, 0
	for c := range elves {
		minX = maths.MinInt(minX, c.X)
		minY = maths.MinInt(minY, c.Y)

		maxX = maths.MaxInt(maxX, c.X)
		maxY = maths.MaxInt(maxY, c.Y)
	}

	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func second(elves map[maps.Coordinate]*terrain) int {
	propositions := []proposition{
		{move: maps.North, check: []maps.Direction{maps.North, maps.NorthEast, maps.NorthWest}},
		{move: maps.South, check: []maps.Direction{maps.South, maps.SouthEast, maps.SouthWest}},
		{move: maps.West, check: []maps.Direction{maps.West, maps.SouthWest, maps.NorthWest}},
		{move: maps.East, check: []maps.Direction{maps.East, maps.SouthEast, maps.NorthEast}},
	}

	elfAt := make(map[*terrain]maps.Coordinate, len(elves))
	for c, e := range elves {
		elfAt[e] = c
	}

	var pidx int
	round := 1
	for ; ; round++ {
		proposes := make(map[maps.Coordinate][]*terrain)

		for c, e := range elves {
			shouldMove := false
			for _, c := range c.Surrounding() {
				if _, ok := elves[c]; ok {
					shouldMove = true
					break
				}
			}

			if !shouldMove {
				continue
			}

			moveTo := c
			for j := 0; j < len(propositions) && moveTo == c; j++ {
				p := propositions[(pidx+j)%len(propositions)]

				proposeMove := true
				for _, d := range p.check {
					if _, ok := elves[d.Apply(c)]; ok {
						proposeMove = false
					}
				}

				if proposeMove {
					moveTo = p.move.Apply(c)
				}
			}

			proposes[moveTo] = append(proposes[moveTo], e)
		}

		var moved bool
		for c, es := range proposes {
			if len(es) > 1 {
				continue
			}

			moved = true
			e := es[0]
			delete(elves, elfAt[e])
			elfAt[e] = c
			elves[c] = e
		}

		if !moved {
			break
		}

		pidx += 1
		pidx %= len(propositions)
	}

	return round
}

func print(m map[maps.Coordinate]*terrain) {
	minX, minY := math.MaxInt, math.MaxInt
	for c := range m {
		minX = maths.MinInt(minX, c.X)
		minY = maths.MinInt(minY, c.Y)
	}

	coords := make(map[maps.Coordinate]*terrain)
	for c, e := range m {
		coords[maps.Coordinate{X: c.X - minX, Y: c.Y - minY}] = e
	}

	fmt.Println(maps.MapFromCoordinates(coords))
}

type terrain struct {
	isElf bool
}

func (e *terrain) String() string {
	if e == nil {
		return "."
	}

	if e.isElf {
		return "#"
	} else {
		return "."
	}
}

const testInput0 = `
.....
..##.
..#..
.....
..##.
.....`

const testInput1 = `
..............
..............
.......#......
.....###.#....
...#...#.#....
....#...##....
...#.###......
...##.#.##....
....#..#......
..............
..............
..............`

const Input = `
##...##......#.#.#..#####.##.#..#.#..#.###....##..#...#####...##.##..#
.#..##...##.....###.#.###..###.##.##.........###..#..#.#..##.#.#.#....
.###.#..#.#..##.##.#...##..#.#...###.#..##.#..#...#..####.##..##..#...
#..#..#.#....#.####......##.#.#.##..#.##.#.###..##.#...#..#..#.#..###.
.#.#....###.....###..##.###.#.###....##..#...#.####.#....#.#.###..##..
.##.#####...###..#...##...........#...##.#.###.#.....#...#...###.#....
.#.###.#..#.##.#...#....####...##..##.####..###.##......#..##..##...##
####..###.#..##...#..##.....###...######...##.....#..##..###.#.##.####
.#.##.###.#..#.#....######.#.####.##....#.#.#.#.###.###....#.######.#.
#####.#..####..######..##...#.####..####.....##..#.#.#.....##...#.###.
#..###.......##...#.###..#.##.####.....#...##.....#....#.#.##.#.##..#.
##..##.###.##.#....##.#.#..####..#.##.##..#.##....##.#.#..##.###...##.
.######..#..#..#####.###..####...#.#..#.##..####....#.###..#..#####...
#.#.###....##....##.#...#...###.##.##.##.##.#.#####.#####.......##.##.
....#.#.#.###.#...###.##.#..##.##..##.#.#....##.#.#.#####.#..##..###.#
.######...##..#..#...#..#.###.##.....#..##..#.##....#..####.##...##...
..#.#..#.###.#.#.##.#....##.###...#.#...##.#......#.######..##.#.#.#..
#...##..#..####.#.#...#.#.#...#....##..##..#.#..##.#..#.###....#.#...#
#..###...#.#.#.####.##..#.#..###.#..#...#..##..##.#....###.#...#.#..#.
.##....####.#.##.###.###..####..####.#.....##.#.#.####..##.##....####.
#.###.#...###....####......#...#..##...#.#.#..#.##.#.#.##..#....#...#.
##.#.#.####.##...###...#####......##..##.##.#..##..##...###..##.#.#.##
#.....#.##..#..##..#...###.###...#...#.#.##.....#..#.#.#.##.#..#.....#
#..#..###.#####..##.###.####...#.#####..##.....#.#.......#.#.#...###..
#.#.##.....###....#..#.###..###..#.#####..######..####.####...##..###.
#......#...#......#...##.#.##############.##......#..##.#.#.####.#.###
.####.#....#.###..##.#.##..######..#..#.###..#.......##.####.....#.#..
#..##..##.##..#..#..##...##.....##..#...##..#.###.#..###.##..#..#...##
###.##..#.#.####.####...#.#..#..##...###.....#...###..#..#...##..#.#..
#.##.#.###.....###.####.#.##.##.###.##..##.#..#.###...#....###.##..##.
##..###.#.##.##...##.##....#...#...#....###..#...#.#...#..##.#.#..#..#
##.#.###.#.#.######.#####.#..........##..#...##.#..#..#..##.##.#.#..##
.#..#####.##.##.#.#...#.##.....#..##.##..#.#..####.###..####.#..##..##
##.##.##.###.#..#.#..######.#.#.....#.#.#.#.#...#.##..#.#.###..#.#..#.
##...##.#...##.#..###..#..##......##...#............###.##.#.#.#.#.###
#..#..####..###.#.###.#.##..#####.##.####...##.##..#..##.#.###.#...###
.##.#.#.#...#..##..#...##...###.#...#.##.##..........##.##.###.#......
##..###.####.#....##.....#..##.#...#.##..#.#.#.##.#..#.#...#...##.##.#
#.#..#..#...##..#..###..#..#..#.#..#.#####.##..#...#..#..#..##.#.#.#..
.##..##..#..#.##.#..#.###..#.#...#######.#....#...####..##.#.##.#.#.##
.#####.....#....#.###...#########....#....#.##.#.##..####...#.##.#....
####..#..#.#....###.#...#.#.#####....#..#.###.#.#.#.##.########.##.#..
...#.######.#..##.##.#..##...###.#.#.#.#.#.#..#..###.####.##.##.#.##..
.#.#.##....###.#.#######.###...#.##.##.##.#..#######.#....##.....##..#
.#...##.#..#.##...#.####...#..#.##.#.##.#....##..##.###..##.##..#...#.
.###....#################..#.#.#..##.#.#..#.#.##..#.#.##.##.#....#..##
..#...#.#.###.#.#.#.####.#.##.#.#####..###..###....#####..##...##.#.##
.#...###..#.###..##.##.##.#......#.#..#.####.##.#.....#.....#.##.##.#.
###..##.#.##...##.#.....####.##.#.#.#...###..###.##.#..#####..#.#...##
.##.#.##...#.#...####.#.#..#...#.#.##..###.##..###.#.####..#.#..##..#.
####.##...###...##...########..#..##..#..#...#..#.#....##.#.#.#.##..#.
#..#...#.###..##.#..##..#.#.#.##.#..##.####..#.##.#.##..##...####..#..
...#.#........###.#####...###.#.###..#.....#####.####..########......#
.##..#....#.#...#..#####.#...###.#....#.##.##......##..#..#######....#
#..#####.##.##.##.##....#..#..#...#........##.#.#.##.##.####.###......
#.###..##.#....##....###..##.#.#.#.#####.##.##..##.##...####......###.
.##.....###.###..#####....##...##..#.#..#.#...##..#....#..#...#....##.
....##.##..##..#.#####.###.#..#.###.#..#..####..#.#.####...#.######.##
#.###...#.##.######.#..##.##.#..##...###.#.#..##...##....##..#...###.#
..#.##..#.#.###...####.#.#....##..##..##...#.##..#...#...#.....#..##.#
..##...#.#.#..#..###.###..#....##.#.#.#.#######.###.##....#..##.##.###
#......#######.#..#.######.##..#.#.###.###..###..##.####.....#....#.##
####..##..#.#....###.#..#...###..###.....######.##.#....####.#.#.####.
.##.###.###.###.###.#####.#.##.#.#..###..#.#..#.##...##.....##.#####.#
##..#.#.#.#.#.##.####.#.##...#.#.##......#....#.###.##....#...##..####
.######.#.###.#..........#######.#.#.##..#..##...#.####...#.##.....###
#..#..####.#.......#...#.#....####.#..#...#####..###.########..#...##.
#.....##.####.##.#.##.#..#.#.#.#..#.#...##.##.#.##.#..#..##..#.#.#.##.
..#....###.##.###.#..###.#.####.####.#.##...#######.##..####.##.......
.#.....#####...####..#.#..#..###.#...#..##...#..#..###.#.####...###.##
`
