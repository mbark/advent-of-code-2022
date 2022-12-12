package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/maps"
	"math"
)

const testInput = `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

func main() {
	var start, end maps.Coordinate
	var lowest []maps.Coordinate
	m := maps.New[elevation](Input, func(x, y int, b byte) elevation {
		switch b {
		case 'S':
			start = maps.Coordinate{X: x, Y: y}
			return elevation{height: 0, start: true}
		case 'E':
			end = maps.Coordinate{X: x, Y: y}
			return elevation{height: int('z') - 97, end: true}
		default:
			if b == 'a' {
				lowest = append(lowest, maps.Coordinate{X: x, Y: y})
			}
			return elevation{height: int(b) - 97}
		}
	})

	fmt.Printf("first: %d\n", bfs(m, start, end))
	fmt.Printf("second: %d\n", second(m, end, lowest))
}

type elevation struct {
	height int
	start  bool
	end    bool
}

func (e elevation) String() string {
	switch {
	case e.start:
		return "S"
	case e.end:
		return "E"
	default:
		return string(byte(97 + e.height))
	}
}

func second(m maps.Map[elevation], end maps.Coordinate, starts []maps.Coordinate) int {
	shortest := math.MaxInt
	for _, s := range starts {
		steps := bfs(m, s, end)
		if steps > 0 && steps < shortest {
			shortest = steps
		}
	}

	return shortest
}

func bfs(m maps.Map[elevation], start, end maps.Coordinate) int {
	visited := make(map[maps.Coordinate]bool)
	from := make(map[maps.Coordinate]maps.Coordinate)
	queue := []maps.Coordinate{start}
	visited[start] = true

	var curr maps.Coordinate
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if curr == end {
			break
		}

		atHeight := m.At(curr).height
		for _, a := range m.Adjacent(curr) {
			height := m.At(a).height

			if _, ok := visited[a]; !ok && atHeight+1 >= height {
				visited[a] = true
				queue = append(queue, a)
				from[a] = curr
			}
		}
	}

	if curr != end {
		return -1
	}

	var steps int
	for at := end; at != start; at = from[at] {
		steps += 1
	}

	return steps
}

const Input = `
abccccccccccccccccccaaaaaaaaacccccccccccccccccccccccccccccccccccccaaaa
abcccccccccccccccaaaaaaaaaaacccccccccccccccccccccccccccccccccccccaaaaa
abcaaccaacccccccccaaaaaaaaaacccccccccccccccccccccaaacccccccccccccaaaaa
abcaaaaaaccccccccaaaaaaaaaaaaacccccccccccccccccccaacccccccccccccaaaaaa
abcaaaaaacccaaacccccaaaaaaaaaaaccccccccccccccccccaaaccccccccccccccccaa
abaaaaaaacccaaaaccccaaaaaacaaaacccccccccccaaaacjjjacccccccccccccccccca
abaaaaaaaaccaaaaccccaaaaaaccccccaccccccccccaajjjjjkkcccccccccccccccccc
abaaaaaaaaccaaacccccccaaaccccccaaccccccccccajjjjjjkkkaaacccaaaccaccccc
abccaaacccccccccccccccaaccccaaaaaaaacccccccjjjjoookkkkaacccaaaaaaccccc
abcccaacccccccccccccccccccccaaaaaaaaccccccjjjjoooookkkkcccccaaaaaccccc
abcccccccaacccccccccccccccccccaaaacccccccijjjoooooookkkkccaaaaaaaccccc
abccaaccaaaccccccccccccccccccaaaaacccccciijjooouuuoppkkkkkaaaaaaaacccc
abccaaaaaaaccccccccccaaaaacccaacaaaccciiiiiooouuuuupppkkklllaaaaaacccc
abccaaaaaacccccccccccaaaaacccacccaaciiiiiiqooouuuuuupppkllllllacaccccc
abcccaaaaaaaacccccccaaaaaaccccaacaiiiiiqqqqoouuuxuuupppppplllllccccccc
abccaaaaaaaaaccaaaccaaaaaaccccaaaaiiiiqqqqqqttuxxxuuuppppppplllccccccc
abccaaaaaaaacccaaaaaaaaaaacccaaaahiiiqqqttttttuxxxxuuuvvpppplllccccccc
abcaaaaaaacccaaaaaaaaaaacccccaaaahhhqqqqtttttttxxxxuuvvvvvqqlllccccccc
abcccccaaaccaaaaaaaaaccccccccacaahhhqqqttttxxxxxxxyyyyyvvvqqlllccccccc
abcccccaaaccaaaaaaaacccccccccccaahhhqqqtttxxxxxxxyyyyyyvvqqqlllccccccc
SbcccccccccccaaaaaaaaaccccccccccchhhqqqtttxxxxEzzzyyyyvvvqqqmmlccccccc
abcccccccccccaaaaaaaacccaacccccccchhhppptttxxxxyyyyyvvvvqqqmmmcccccccc
abccccccccccaaaaaaaaaaccaacccccccchhhpppptttsxxyyyyyvvvqqqmmmccccccccc
abcaacccccccaaaaaaacaaaaaaccccccccchhhppppsswwyyyyyyyvvqqmmmmccccccccc
abaaaacccccccaccaaaccaaaaaaacccccccchhhpppsswwyywwyyyvvqqmmmddcccccccc
abaaaaccccccccccaaaccaaaaaaacccccccchhhpppsswwwwwwwwwvvqqqmmdddccccccc
abaaaacccccccccaaaccaaaaaaccccccccccgggpppsswwwwrrwwwwvrqqmmdddccccccc
abccccccaaaaaccaaaacaaaaaaccccccaacccggpppssswwsrrrwwwvrrqmmdddacccccc
abccccccaaaaaccaaaacccccaaccccaaaaaacggpppssssssrrrrrrrrrnmmdddaaccccc
abcccccaaaaaaccaaaccccccccccccaaaaaacggppossssssoorrrrrrrnnmdddacccccc
abcccccaaaaaaccccccccaaaaccccccaaaaacgggoooossoooonnnrrnnnnmddaaaacccc
abccccccaaaaaccccccccaaaacccccaaaaaccgggoooooooooonnnnnnnnndddaaaacccc
abccccccaaaccccccccccaaaacccccaaaaacccgggoooooooffennnnnnnedddaaaacccc
abcccccccccccccccccccaaacccccccaacccccggggffffffffeeeeeeeeeedaaacccccc
abccccccccccccccccccaaacccccaccaaccccccggfffffffffeeeeeeeeeecaaacccccc
abccccccccccccccccccaaaacccaaaaaaaaaccccfffffffaaaaaeeeeeecccccccccccc
abccccccccaacaaccccaaaaaacaaaaaaaaaaccccccccccaaaccaaaaccccccccccccccc
abccccccccaaaaacccaaaaaaaaaaacaaaaccccccccccccaaaccccaaccccccccccaaaca
abcccccccaaaaaccccaaaaaaaaaaacaaaaacccccccccccaaaccccccccccccccccaaaaa
abcccccccaaaaaacccaaaaaaaaaacaaaaaacccccccccccaaccccccccccccccccccaaaa
abcccccccccaaaaccaaaaaaaaaaaaaaccaaccccccccccccccccccccccccccccccaaaaa
`
