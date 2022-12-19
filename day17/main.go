package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/maps"
	"strings"
)

const testInput = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func main() {
	var pushes []push
	for _, b := range Input {
		pushes = append(pushes, newpush(b))
	}

	fmt.Printf("first: %d\n", first(pushes))
	fmt.Printf("second: %d\n", second(pushes, 1000000000000))
}

var (
	xMin = 0
	xMax = 8
)

func first(pushes []push) int {
	state := mapState{
		top:       0,
		rockCount: 0,
		pieceIdx:  0,
		pushIdx:   0,
		rocks:     make(map[maps.Coordinate]bool),
	}

	for state.rockCount < 2022 {
		piece := spawnPiece(pieces[state.pieceIdx], maps.Coordinate{X: 2, Y: state.top + 3})
		for {
			p := pushes[state.pushIdx]
			state.pushIdx = (state.pushIdx + 1) % len(pushes)
			pushed := piece.push(p)

			canMove := !(isOutside(pushed) || hitWall(pushed, state))
			if canMove {
				piece = pushed
			}

			fallen := piece.fall()
			canFallDown := !(isAtBottom(fallen) || hitWall(fallen, state))
			if canFallDown {
				piece = fallen
				continue
			}

			for c := range piece.points {
				if c.Y+piece.diff.Y > state.top {
					state.top = c.Y + piece.diff.Y
				}

				state.rocks[c.Add(piece.diff)] = true
			}
			state.rockCount += 1
			break
		}

		state.pieceIdx = (state.pieceIdx + 1) % len(pieces)
	}

	return state.top
}

type hashState struct {
	count int
	top   int
}

func second(pushes []push, goal int) int {
	state := mapState{
		top:       0,
		rockCount: 0,
		pieceIdx:  0,
		pushIdx:   0,
		rocks:     make(map[maps.Coordinate]bool),
	}

	hashes := make(map[int]hashState)
	hashIndex := make(map[int]int)
	var hash int
	for state.rockCount < goal {
		piece := spawnPiece(pieces[state.pieceIdx], maps.Coordinate{X: 2, Y: state.top + 3})
		for {
			p := pushes[state.pushIdx]
			state.pushIdx = (state.pushIdx + 1) % len(pushes)
			pushed := piece.push(p)

			canMove := !(isOutside(pushed) || hitWall(pushed, state))
			if canMove {
				piece = pushed
			}

			fallen := piece.fall()
			canFallDown := !(isAtBottom(fallen) || hitWall(fallen, state))
			if canFallDown {
				piece = fallen
				continue
			}

			for c := range piece.points {
				if c.Y+piece.diff.Y > state.top {
					state.top = c.Y + piece.diff.Y
				}

				state.rocks[c.Add(piece.diff)] = true
			}
			state.rockCount += 1
			break
		}

		state.pieceIdx = (state.pieceIdx + 1) % len(pieces)
		hash = state.hash()
		if s, ok := hashes[hash]; ok {
			cyclesEvery := state.rockCount - s.count
			heightPerCycle := state.top - s.top

			divisible := goal / cyclesEvery
			leftover := goal % cyclesEvery

			return divisible*heightPerCycle + hashes[hashIndex[leftover]].top
		}

		hashIndex[state.rockCount] = hash
		hashes[hash] = hashState{count: state.rockCount, top: state.top}
	}

	fmt.Println(len(hashes))

	return state.top
}

type mapState struct {
	top       int
	rockCount int
	pieceIdx  int
	pushIdx   int
	rocks     map[maps.Coordinate]bool
}

func (s mapState) hash() int {
	hash := 1000*s.pieceIdx + 1000*1000*s.pushIdx

	tops := make(map[int]int, xMax-2)
	for c := range s.rocks {
		if c.Y > tops[c.X] {
			tops[c.X] = c.Y
		}
	}

	for x, y := range tops {
		hash += x + 10*(s.top-y)
	}

	return hash
}

type part string

func (s mapState) Render(piece *piece) string {
	var cells [][]part
	yMax := s.top + 7
	for yRev := 0; yRev <= yMax; yRev++ {
		y := yMax - yRev
		var c []part
		for x := 0; x <= xMax; x++ {
			switch {
			case y == 0 && (x == xMin || x == xMax):
				c = append(c, "+")
			case y == 0:
				c = append(c, "-")
			case x == xMin || x == xMax:
				c = append(c, "|")
			case s.rocks[maps.Coordinate{X: x, Y: y}]:
				c = append(c, "#")
			case piece != nil && piece.points[maps.Coordinate{X: x - piece.diff.X, Y: y - piece.diff.Y}]:
				c = append(c, "@")
			default:
				c = append(c, ".")
			}

		}

		cells = append(cells, c)
	}

	m := maps.Map[part]{Columns: xMax, Rows: yMax, Cells: cells}
	return m.String()
}

func isOutside(piece piece) bool {
	for c := range piece.points {
		if c.X+piece.diff.X <= xMin || c.X+piece.diff.X >= xMax {
			return true
		}
	}

	return false
}

func isAtBottom(piece piece) bool {
	for c := range piece.points {
		if c.Y+piece.diff.Y == 0 {
			return true
		}
	}

	return false
}

func hitWall(piece piece, state mapState) bool {
	for c := range piece.points {
		if _, ok := state.rocks[c.Add(piece.diff)]; ok {
			return true
		}
	}

	return false
}

var pieces = []piece{
	{
		points: map[maps.Coordinate]bool{
			{X: 1, Y: 1}: true,
			{X: 2, Y: 1}: true,
			{X: 3, Y: 1}: true,
			{X: 4, Y: 1}: true,
		},
	},
	{
		points: map[maps.Coordinate]bool{
			{X: 2, Y: 1}: true,
			{X: 1, Y: 2}: true,
			{X: 2, Y: 2}: true,
			{X: 3, Y: 2}: true,
			{X: 2, Y: 3}: true,
		},
	},
	{
		points: map[maps.Coordinate]bool{
			{X: 3, Y: 3}: true,
			{X: 3, Y: 2}: true,
			{X: 1, Y: 1}: true,
			{X: 2, Y: 1}: true,
			{X: 3, Y: 1}: true,
		},
	},
	{
		points: map[maps.Coordinate]bool{
			{X: 1, Y: 1}: true,
			{X: 1, Y: 2}: true,
			{X: 1, Y: 3}: true,
			{X: 1, Y: 4}: true,
		},
	},
	{
		points: map[maps.Coordinate]bool{
			{X: 1, Y: 1}: true,
			{X: 2, Y: 1}: true,
			{X: 1, Y: 2}: true,
			{X: 2, Y: 2}: true,
		},
	},
}

type piece struct {
	points map[maps.Coordinate]bool
	diff   maps.Coordinate
}

func spawnPiece(p piece, start maps.Coordinate) piece {
	m := make(map[maps.Coordinate]bool, len(p.points))
	for c := range p.points {
		m[maps.Coordinate{X: c.X + start.X, Y: c.Y + start.Y}] = true
	}

	return piece{points: m}
}

func (p piece) String() string {
	var cs []string
	for c := range p.points {
		cs = append(cs, c.String())
	}

	return "[" + strings.Join(cs, ",") + "]"
}

func (p piece) push(push push) piece {
	var diff maps.Coordinate
	switch push {
	case pushLeft:
		diff = maps.Coordinate{X: p.diff.X - 1, Y: p.diff.Y}

	case pushRight:
		diff = maps.Coordinate{X: p.diff.X + 1, Y: p.diff.Y}
	}

	return piece{points: p.points, diff: diff}
}

func (p piece) fall() piece {

	m := make(map[maps.Coordinate]bool, len(p.points))
	for c := range p.points {
		m[c.Up()] = true
	}

	return piece{points: p.points, diff: maps.Coordinate{
		X: p.diff.X,
		Y: p.diff.Y - 1,
	}}
}

type push bool

const (
	pushLeft  = push(true)
	pushRight = push(false)
)

func newpush(b int32) push {
	if b == '<' {
		return pushLeft
	} else {
		return pushRight
	}
}

func (p push) String() string {
	switch p {
	case pushLeft:
		return "<"
	case pushRight:
		return ">"
	}
	return ""
}

const Input = `>>><<<>>><<<>><<>><<>>>><>>>><><<><<<>>>><<<>>>><<<>><><><><<><<<>><>>><>>>><<>>>><<<<>>><><<>>><<<<><<<>><>><<<<>><<<<><<<<><><<<><<<>>><><<>>><<>>><<<<>><>>><<<<>>>><<<<><>>>><<<><<>><>><<<><<<>>>><<<<>><>>><>>>><<><<<><<>><<>>>><<<>><<<>>><<<<><<>><<><>><<>>><<<>>>><<<>>><<<<>><>>>><<<<>><<><<<>>>><<<>><<><<<<>><<>>>><<<<><<<<>>>><<<<>>><>><<<<><<<>><<<<>>>><<<>>><<<<>>>><<<<>>><<<>><>><<<><<>>>><<<>>><<<>><<<>>><<<<>><><<<<>>><<>>>><>><<>><<<><>><>>><<><<<<>>>><<<<>>><<>><>>><<<<><<>><<>><>>>><<<<><<<>>>><<>>><<<>><<><<><<>>><<<>>>><<<<>>><<>>>><<<<>>>><>>><><>>><<<>><<<<><>>>><<>><<>><<<>>>><<<<><<>>>><<>>>><<>><>>>><<>>>><<<><<>><>>><<>>><<><>>>><<<<>>>><<>><<<<>>><<>><>><><<<>>>><<<><<<>><><>>>><<<<>>><>><>><><<>>>><<>>><<><<<<>>>><<<>>><>>><<<<>>><<>>><<<>>><<<>>><<<>>><<<<>>><<<>><><<<<>>>><<<><<<<>><<<>><><<<<>>>><<><<<<>>>><<<<>><<<><>><<<<>><<><<<<>>>><>>><><><><<>>><<<>>>><<<<><<<><<<<><<<>>><<<><>><<<><<<<>>><<<>>><<>>><<<<><><<<<><<<<>>><<>>><>>>><<<>><<>>><<>>><><><<<><><<<<>>><<<<>><<<><<<<>>><>>>><><<<<><<<>>>><<<<>>>><<<<>><<<<>><<<<>><>><>>>><<>><>>><<<><>>><<<>>>><>><>>>><<>><><<<<>>>><<>><<<><<<>>>><>>>><<<<>><>>>><<<<><<<<><>>>><<<<>><<<>>>><<><<<<>>>><<<>><<<<>>>><<<<>><<<<>>>><>><<>>>><<<<><<<>><<<<><<<<>>><<<>><<<<>>><<<<>><<<<>>>><>>><<>>>><<>><<<<><<>>><><<<<>>>><<<>><<<<><<>>>><>><<<>>>><>><<>>><<<>>>><<<<>><<<>><<<><<<<>>>><<<<>>>><<<<>><<<><>>>><<<<>>>><<<>>>><<>>><<>>>><<>><<<<><<<>>>><<<>>>><<<<>><<<<>>>><<><<>>>><<<<>>>><>><<<<>><<<><<<><<<<>><<<<>>><<>>>><><<<<>>>><<<<><>>><<<<>><>><<>>><<<<>>><<>><<<>>><<<>>><>><<>><>><>>>><<<>>>><<<><>>>><>>>><><>>>><<>>>><<<<>><<>>>><<<<>><<<<>><<<>>><<<<>><>>><<>>><<<>><<<<>>><>>>><<<<><<<><>>>><<<<>>><>>>><>><<><<<><<<>>><<<<>>>><>><<<<>><><<><<>>>><<<>>>><<>>>><>>><<<<>>><<<<>>>><<>><<>>>><><>><<<<>>>><<<>><<>>><<<><<<>>>><<>><<<>><<<<>>><<<><<<><><<<>><>>><>><<<<>>>><>>>><<<<><>>><<>>><>>>><<>>><<<>>>><>>><<<<>>>><<>>>><<<><<<><<>>><<<<>><>>>><<<>>><<<<>>>><<<<>><<<>>>><<<>><<><<>><<<<>>><>><>>>><<<>><<<>>>><<<><<>>><>>>><<<<><>>>><<>>><<<>>>><<<><<><<>>><<>><<<>>><>>><<<><>><<>>><<<>>><<<<>>>><><<>>>><<<>>><<<<>>><>>>><<<<>><<>>><<>>>><<>>>><>><>>>><<><><<<<>>>><<<<>><<<>><<<><<<>>>><<<>>><>>>><<><<>>><<>>>><><<<><<<<>><<<<>>><<<><>>><<>>>><<<>>><<>><<>>><>><<<<>>>><>><<>>>><<<<>><>><<<>>>><<<>><<>>>><>>>><<><<<>>><>>>><<<>>><<><<<<>>><<>>>><<<>><<<>>><>>><>><<<>>>><>><<><<>><><>>><>>><<>><>><<<<>>>><<>>>><<<<>>>><><>>><<>>>><><>><>>><<<><<<>>>><><<<<>>><<<>>><>>>><<<<>>>><<<<>>>><<<>>><<<<><<>><<>><<<>>><>>><<>><>>>><<<>><<<>>>><<<<>><<<<>><<<><<<>>>><<>>>><<<<><<<>><>><<<<>>><<<>>>><<<<>>>><<<>>><<<><>><>><><<<>><<>><>>><<<><<<><<<<>>>><><<<>><>>>><<<>>><<>><><>>>><><<<><<>><<<>>>><<>><<<<>>><<<><<<>>><<>><<>><<>><<<<><<>><>>>><<<<>><>><<<><>>>><<<<><<>><<<<>><<<>><<><<<<>>>><>><<<><<<<>>>><<<<>>>><>>>><<>>><<>>>><<>>>><>>>><<>>><<<><><<>><><<<<><>>>><>><<<>>><<<>>>><<>><<>>>><<<<>>><<>><><<<><<<<><<<><<>>><>>><<<<>>><<<>><>><>><<<>>>><><<<<>><>>><<<>><<<<>>>><<>><<<>>><<>>><><<<>>>><<<>><<<><>><<<>><<<>><<<>>><<<<><>>><><<<<>><>>><<<<>>><<<<>>><>>>><<<<>><>>>><<<>>><<<><<<<>>>><<<>><<<<>>><<<<><<<><>><<<>>>><<>>><<<>>><>>>><<<<>>>><<<<>>>><>>><>>><><<<<>><<>>>><<<>>><<>>>><<<>>><>>>><>>>><<<>>><>><<>><>>>><<>>>><<>><<<<>><>>><<<>>><>>><<>><<<>>><<>>><<<>><<>>>><<>>>><<>><<<><<<>><>>>><<<>>><<>>><<>><>><><<<<>>>><>>><>><>>>><<<<>>>><<>>>><><<>>><<>>><<><<<<><<>><<<<><<>>>><>>>><><>>><>><<<<>>><>>><<>><<<>><>>><<><<<<>>><><><<<<>><<<>><>><<<>><<<<>>><<<>><<<<>>>><<<>>>><>>><<<<><>>>><<<<>>>><><<>>>><>>><<>>><<>>>><<>>>><<>>>><>>>><>>>><>><<><>>><<<>><>><<><<<<>>><<<>>><>><<<>>><<<<>>><<>>><>>>><<<<><<>><>>><<<<><<<>>>><<<><<>><<>><><<>>><<<>><<<>>><>><<<>>><>><<<>>><><<>>>><<<<>>><<<><<>>>><><<>>>><<<>>><>>><<<>>>><<<>><><<<<>>><<>>>><<<>><<>><<<>>>><<<>>><>>><>>><<><<>>>><<<>>>><<<<>>>><<<<><><<<<><<>><<>><>>>><<>><<>>><>>><<>><<<<>>>><<>>><><<><>>><<>>>><<><<>>><<>><<>><<<<>><<<<>><>>><<>><<<>>>><<<>><>>><<<<>>><>>>><<>><<>>>><<<<>>><<<<>>><>>><<<>>>><<<<>><<<><<<>>>><<<>>><>>>><<<<>><>>>><>>><<>><<<<><<>>>><<>>><<<>>><<<<><>><<<<>><>>>><<><<<<><<<>>><><>>><<<<><>>>><<<><<<><<<<><<>><>><<<<>>>><<<>>>><>><>>><><>>><<>><>><<>><>>><>>><>>><><<>><<<>>><<<>>><<<<>>>><<>>>><<><<<>>><<>>>><>>>><<<<>><><<>>>><<>><<<>>>><<<>><>><<>><<<<><<>>>><<<<>>><<<<>><<>><<<>>>><<<<>>><<<<>>><>>><<<<>>>><<<<>>>><<>>>><>><<<><<<<>><<>>><<>><<<<>>>><<<<>>>><<>>>><<>>><<<>><<<><<><><>>>><>><<<<><>>>><><<>><<<<>><<>>><>>>><<<>>><<<<>>><<>><<<<>><><<<<>><<<<>>><<><>>><<>><<<<>>>><<<<>>><<<<><<<<>><>>>><<>>><<<>><<<<>><<<<><<<>>>><<<>><<>>>><<<>>>><<><>>>><>>><<>><<<<><<<<>><<<<>>><><<>><<<<>>><<<>>>><<<><<<<><<>>>><<><<<<>><<<<>>>><<<>><>><<<<>><<<<>>>><<<<>><<<>>>><<<<>>>><<<>>><<<>>><<<<>>>><<<<>>>><<>>><<<<><<<<>><<<>><<<<>>><<<<>><<<<><<<<>><>><<>><<<>>><<<><<<<>><<>><<<>>><<>>>><<>><>>><<<<>><><>>><<<<>><><><<<>><<<<>>>><<<>>>><<<<>>>><<<>><<<<>>><<<>>><>><<><<>><<><<><><>>>><>>>><<<><<>>>><<<><<>><<<>><<<<>><<>>>><<<>>>><<<><<>>>><<>><<><<>>>><<<>>>><<<<>>><<<<>>><<>><>><<<>>>><<<<><<<>>>><<<<><>>><<<>>>><<<>><>>>><><<>>><<>>><<<>>>><<><<<>>>><<<><<<>>><<>>>><>><<<<><>>>><<<>>><<<<><<>><><<<><><>><<<<>>><<<>><<<<><<<<>><<>>>><><<<<>>>><<<>>><<>>>><<>>>><>><<<<>><<<><<>><<>>><<<<>>>><<<>>><>><>><<<><>>><<<<>><>><>>>><<><>>><<<>><<<<>>>><>>><<<<>><<<<><<<>><<<<><<<<>>><>>><<<<>>><>>>><<><<<>>>><<<<><<<<>>>><<<>>>><<<<>>><<<<>><<<<><<<>>>><><<<<><<<>><>>>><<<<>>>><<>>><<>><<<<>>><<>><<<<>>>><><<<<><<>>><<<<>><<<<><<>><<<>>><><><<>>><>>>><<>>><>>>><<<>><<<<>>>><>>><<<>>><<>>><<<<>><><><<<><<><<<>>><<>>>><><<<<>><<>>>><<<>>>><<<<>><<>><<<<>>><><<<<>>>><<<<><<<<><>>><>>><<>><<<>>>><<<<>>>><<><>><<<>>>><>>><<<>>>><<>><<<<>>>><<<>>>><<<>>>><<><>>><<>>><<<<><<>>>><>>><<<>>>><<<<>>><<><><>><><<<<><<<<><<<>>>><<<<>>><<<<>><<>><<<<>>><>>>><<<>>><<>>><>><<>><<>>><<>>><<<><<<>>>><><<><><<<<>>>><<>>><<<<>>><<<><<<<>><>>><<<<><<<<><<>>>><>><<<>>>><<>><><<<>><<>><<<<>>>><>><<<<>>>><<<<><>><>>>><<<><><<>>>><<<<>>>><<<<>>><>>>><<<<>>>><<<>><<<<>>>><<<<>>><>><><<<<>><<>>><<>>>><<<><>>><>>>><<<>>>><<<>>>><<<<><<<>>>><>>>><<<<>>><>>>><<><<<>>><<>><>>><>><>>><>>>><<<<>>><>>><<<<>>><<<>>>><>><<<<>>>><<<<><<>>>><<><>>>><<<<>>>><<<<>>>><<>>><>>>><<<<>>>><<<>><><>><>>>><<<<>>><<<>>>><><<>>>><<<<><<<>>><<>>>><<>><<<<>>>><><<<<>>><<>><>><>><<<<>>><<<><<<><<<<>>><<>>><<<<><<<<>><><>><>>>><>>>><<<>>>><<<><<><><><<>>>><<<<>>><<>>><>>><<<<>><><><<<>>>><<>>><<<><<<<>>>><<<<>>>><<<<><<<<>>>><<<>>><<<<>>><<>><<<<><<><>>>><><<<><<<><<>>><<<<>>>><<>><><>><<>><<<>><>><>><<<>><><<<>><><>>>><<>>>><<<><<><>>><<>>><<<<>>>><>><<>>>><<<><<<<>>><>>>><<>>><<<<>>>><>>><><>>><<<<>>><<<<>>>><<<>>>><<><><<>>>><><<><<<<>><<<<>>><<>>><<<>><>>>><<<>><<<>>>><>><>><><>>><>>><<<>>>><>><<<<>>>><<<<>>>><<<>>>><>><<<<><<<<><<<>><<>>>><<<<>><<<><<<<><<>>>><><<<><<<<>><<<><<><<<<>>>><<<<>><<>>>><<<<><><<<><<>><<<<>>><<<>><<<>>>><>><<<<>>><<<>>>><>>><>>><<<<>>>><<<>><<>>><<<<><>>>><<<<>>>><><<<>>><>>>><<<<><><<<<>>>><<<>>>><<<>><>><<<<>>>><<<<>>>><>>><<<<>>><<>>>><>><<<<>>><<<<>><<<>>><>>><<<>>><><<<<>>><<<<>>>><<<>>><<<>>>><<<<><<<>><<>><<<<>>><>><<<>>><>><<<>>>><<<<>>>><<<<>><<>><<>><>><><<>>><<<<><<<>><<<<>>><<><<<>>>><<<<>><>>><><<<>>><<>>><<<<><<>><<<<><>>>><><<<<>>><><<<><<<><<<>><<>>><<><>>><>>>><<<>>><<<>>><<>><<><>><<>><>>><<<<><<>>>><<>>>><<<>><<>>><><><<>>><>>><><<<>>><<<>>><<>><>>><>><<>><<>>>><<<>><<<<>>>><<<<><<<<>>>><<<>>><>><>>>><<<>>>><<<><<<><<<<><<<>><<>><<>><<<><<><<<<>>>><<<>>><<<>>>><<<<>>>><<<><>>>><><<>>><<<><><<<<><<<>>>><<><<<<><>>><<<><<<<>>>><<<<>>><><<<<>>><<<<>><<<<>>><<>><>><><<<<><<>>>><<<>><<<><<<>>><>><<><>>>><<<<>>>><<<>><<<><<><<<><<>><<<<><<<>>><><<<><<<><>><<>>><<<>><>>><<<>><>><<<>>><<>>><>><>>><>>><<<><<>>>><<<>><<><<<>>>><><<>><<<<><<>>>><>>>><<<<>><>>>><<<<>><<<><<><>>><<<>>><<<<>>><<>>><<>><<<>><>>><<>>>><<>><<<>>>><<<>><>>><>>>><>>>><<>><<<<><<<>>>><<<>>>><<>><<<<>>><<><<><<<><<<><>><<<<>>>><<<>>>><<<><<<><<<<>>>><<<<>>>><<<>>><<<<>><<>>>><<>><>>><>><><><<<<>>><<<<><>>>><>>><<<<>>><>>>><<<>><<<<>>><<>>>><>>>><><<<<>>><<<<><<<>>><<<>><>><<>><>>><<<>>>><>>>><<<>>><<><<<>>><>>>><><<>>><<<<><<>>><<<<>><><<><<>>>><>><<>><<><<>>><<>><<<>>><<><>><>>><<<>>>><<>>>><<<>>>><<>>><<<><<<<>>><<<><<><<<>>>><<>><<<<>><<><<<><<>>><<>>><<<<>>>><<>>><<<<><<<<>>><><>>><<<>>><<><<<<>>>><>>>><<<<><<<>>>><>>>><<<><<<>><<<<>>>><>>>><<<<><<>>>><<>><<>><><<>>><<<>>><<<>>>><<<>><<<<>><>>><<>><<<>>>><<>>><<><<>>><<<<><<><<>>>><<<>>>><<<<>><><>>>><<<<>>>><<>>>><<>>><><<>><<<<>>>><<<<><<<<>>>><<<><<<>><>>>><<<<>>>><<>>><<><<>>>><<<>>><<>>><><<>>><<<<><>>>><<<>>>><<<>>><>>><<<<>>><>><<><>>><<>>><<><>>>><<<><<<>><<<<><<<<>>>><<<<>>><>>>><<>>>><<>><<>><<<<>>><<<>>>><<<>>>><<<<>><<>>>><>>>><>>><<<>>><<>><<<><<<>><>>>><<>>><>><<>>><<<<>><>>>><<<<><<<>>><<>><<<>><<<><<<>>><<<>>>><<<><<<<>>>><<<<><>>>><>><<<<>><<>>>><<><<<<><<<<>>>><><<<<><>><<<<>>>><<<<>>><>>><<>>><<<><<<><<<<><<<>><<>><<<>>>><>><<>>><<<><>>><>>>><<<>>><<<<>>>><>>>><>>><><>>><<<>>>><><><<>>><<<>><>>><<<<>>>><<<<>>>><<<>><>>>><><<<<><><<<><<<<>>>><<><<<<><<><<>>>><>>><<<<>>>><<<><>><<<<>><<<>>><>><>>>><<<>>><><<<<>>><<<>>><><<<<>>><<<<>><<<>><<>>><><>>>><<<>><<<><<><<<<>>>><<<<>>>><<>>>><<>><>>>><<>>>><<<<><<>><<<<><<<>>>><<><<><<<>>>><<<>>>><>><<<>><<>><<<<>><<>><<<>><<<<>>>><<>><<<<>><>><>><<<<><<<<>><<<<><<<>>>><<>><>>>><<<><<>><<>><<<>>>><<<<>>><><>>>><<><<>>><<<<>>>><<>>><<<>>>><<>>><<<>>><<>>><<<<><>>><<<<>>><<<<>><<<><<>><<<<>><<<>>>><<><<<>>><>><>><<<<>>><>>>><><>>><<>>>><>><<>>><<<>>>><>>>><<><><<>>><<<><><<>>><>>>><<>>><<><<<>>>><<<<>><<<<>>>><<<<>>>><<<>>>><<><<<>>><<<><<<<><<>>>><<<<><<<<>>>><<>><<>>><>><<<<>>>><<<>><>>><<<>><>>>><>>><<<>><<<>>><<<<>>><<>><<<><><<>>><<<>>>><<<><<>>>><<<><<>>>><>>>><<<<>>><<<<>>>><<<>>>><<<>>><<<<>>><<>><<<>><<<>>><<<><<<>>><<<<>><<>>><>>>><<<<><>>>><<>><<<>><<>><<<<>>>><><>><<<><<><<<>><<<>>><>>>><<>>>><<>>>><<><<<<>>>><>><<<><<<>><>>><<<>><><<<>>><<<<>>>><>><<><<>>>><<<<><<<<>>><<<>>><<>><>><<<><<>><<<<><<<>><<>><<<<>>><<<<>>>><<<>><<<><<<<>>>><><<<<>><<<>><<<<>><<<>>><<<<>><<><<><>><<<>>><<<<>>><<<<>>><<<>>>><<<>>><<<<>>>><<<<>>><<<>>>><<<>>><>>>><>>>><<<>><<<<>><>>><<>>><<<>>><>>`
