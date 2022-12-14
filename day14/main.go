package main

import "C"
import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/maps"
	"github.com/mbark/advent-of-code-2022/util"
	"strings"
)

const testInput = `
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func main() {
	var lines []line
	for _, l := range util.ReadInput(Input, "\n") {
		split := strings.Split(l, " -> ")
		for i := 0; i < len(split)-1; i++ {
			start, end := maps.CoordinateFromString(split[i]), maps.CoordinateFromString(split[i+1])
			lines = append(lines, line{start: start, end: end})
		}
	}

	fmt.Printf("first: %d\n", first(lines))
	fmt.Printf("second: %d\n", second(lines))
}

type line struct {
	start, end maps.Coordinate
}

func (l line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.start.X, l.start.Y, l.end.X, l.end.Y)
}

func (l line) coordinates() []maps.Coordinate {
	var coords []maps.Coordinate
	if l.start.X != l.end.X {
		xMin := util.MinInt(l.start.X, l.end.X)
		xMax := util.MaxInt(l.start.X, l.end.X)

		for x := xMin; x <= xMax; x++ {
			coords = append(coords, maps.Coordinate{X: x, Y: l.start.Y})
		}
	} else {
		yMin := util.MinInt(l.start.Y, l.end.Y)
		yMax := util.MaxInt(l.start.Y, l.end.Y)

		for y := yMin; y <= yMax; y++ {
			coords = append(coords, maps.Coordinate{Y: y, X: l.start.X})
		}
	}

	return coords
}

func first(lines []line) int {
	filled := make(map[maps.Coordinate]bool)
	for _, l := range lines {
		for _, c := range l.coordinates() {
			filled[c] = true
		}
	}
	rock := len(filled)

	var maxY int
	for c := range filled {
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	falling := maps.Coordinate{X: 500, Y: 0}
	for {
		if falling.Y > maxY {
			break
		}

		down := falling.Down()
		switch {
		case !filled[down]:
			falling = down

		case !filled[down.Left()]:
			falling = down.Left()

		case !filled[down.Right()]:
			falling = down.Right()

		default:
			filled[falling] = true
			falling = maps.Coordinate{X: 500, Y: 0}
		}
	}

	return len(filled) - rock
}

func second(lines []line) int {
	filled := make(map[maps.Coordinate]bool)
	for _, l := range lines {
		for _, c := range l.coordinates() {
			filled[c] = true
		}
	}
	rock := len(filled)

	var floor int
	for c := range filled {
		if c.Y > floor {
			floor = c.Y
		}
	}
	floor += 2

	start := maps.Coordinate{X: 500, Y: 0}
	falling := start
	for {
		if filled[start] {
			break
		}

		down := falling.Down()
		switch {
		case down.Y == floor:
			filled[falling] = true
			falling = start

		case !filled[down]:
			falling = down

		case !filled[down.Left()]:
			falling = down.Left()

		case !filled[down.Right()]:
			falling = down.Right()

		default:
			filled[falling] = true
			falling = start
		}
	}

	return len(filled) - rock
}

const Input = `
494,132 -> 494,134 -> 487,134 -> 487,139 -> 502,139 -> 502,134 -> 499,134 -> 499,132
509,82 -> 514,82
511,113 -> 511,116 -> 509,116 -> 509,120 -> 522,120 -> 522,116 -> 516,116 -> 516,113
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
510,65 -> 510,57 -> 510,65 -> 512,65 -> 512,63 -> 512,65 -> 514,65 -> 514,64 -> 514,65
490,39 -> 490,42 -> 483,42 -> 483,49 -> 503,49 -> 503,42 -> 495,42 -> 495,39
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
526,150 -> 530,150
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
512,127 -> 516,127
494,132 -> 494,134 -> 487,134 -> 487,139 -> 502,139 -> 502,134 -> 499,134 -> 499,132
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
522,153 -> 522,156 -> 515,156 -> 515,164 -> 533,164 -> 533,156 -> 528,156 -> 528,153
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
517,77 -> 524,77 -> 524,76
522,153 -> 522,156 -> 515,156 -> 515,164 -> 533,164 -> 533,156 -> 528,156 -> 528,153
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
529,148 -> 533,148
521,110 -> 525,110
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
511,113 -> 511,116 -> 509,116 -> 509,120 -> 522,120 -> 522,116 -> 516,116 -> 516,113
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
506,123 -> 510,123
490,39 -> 490,42 -> 483,42 -> 483,49 -> 503,49 -> 503,42 -> 495,42 -> 495,39
494,132 -> 494,134 -> 487,134 -> 487,139 -> 502,139 -> 502,134 -> 499,134 -> 499,132
522,153 -> 522,156 -> 515,156 -> 515,164 -> 533,164 -> 533,156 -> 528,156 -> 528,153
522,153 -> 522,156 -> 515,156 -> 515,164 -> 533,164 -> 533,156 -> 528,156 -> 528,153
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
510,65 -> 510,57 -> 510,65 -> 512,65 -> 512,63 -> 512,65 -> 514,65 -> 514,64 -> 514,65
517,77 -> 524,77 -> 524,76
520,84 -> 525,84
506,127 -> 510,127
522,153 -> 522,156 -> 515,156 -> 515,164 -> 533,164 -> 533,156 -> 528,156 -> 528,153
494,132 -> 494,134 -> 487,134 -> 487,139 -> 502,139 -> 502,134 -> 499,134 -> 499,132
494,132 -> 494,134 -> 487,134 -> 487,139 -> 502,139 -> 502,134 -> 499,134 -> 499,132
497,23 -> 497,14 -> 497,23 -> 499,23 -> 499,16 -> 499,23 -> 501,23 -> 501,17 -> 501,23
528,88 -> 533,88
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
497,23 -> 497,14 -> 497,23 -> 499,23 -> 499,16 -> 499,23 -> 501,23 -> 501,17 -> 501,23
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
490,39 -> 490,42 -> 483,42 -> 483,49 -> 503,49 -> 503,42 -> 495,42 -> 495,39
550,150 -> 554,150
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
514,88 -> 519,88
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
506,84 -> 511,84
506,52 -> 512,52 -> 512,51
516,74 -> 520,74
524,104 -> 528,104
506,52 -> 512,52 -> 512,51
511,113 -> 511,116 -> 509,116 -> 509,120 -> 522,120 -> 522,116 -> 516,116 -> 516,113
527,106 -> 531,106
497,23 -> 497,14 -> 497,23 -> 499,23 -> 499,16 -> 499,23 -> 501,23 -> 501,17 -> 501,23
507,67 -> 507,68 -> 518,68 -> 518,67
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
538,146 -> 542,146
497,23 -> 497,14 -> 497,23 -> 499,23 -> 499,16 -> 499,23 -> 501,23 -> 501,17 -> 501,23
515,110 -> 519,110
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
530,108 -> 534,108
497,23 -> 497,14 -> 497,23 -> 499,23 -> 499,16 -> 499,23 -> 501,23 -> 501,17 -> 501,23
503,129 -> 507,129
535,144 -> 539,144
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
507,67 -> 507,68 -> 518,68 -> 518,67
500,88 -> 505,88
507,67 -> 507,68 -> 518,68 -> 518,67
500,127 -> 504,127
511,113 -> 511,116 -> 509,116 -> 509,120 -> 522,120 -> 522,116 -> 516,116 -> 516,113
517,86 -> 522,86
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
497,23 -> 497,14 -> 497,23 -> 499,23 -> 499,16 -> 499,23 -> 501,23 -> 501,17 -> 501,23
527,110 -> 531,110
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
513,84 -> 518,84
510,65 -> 510,57 -> 510,65 -> 512,65 -> 512,63 -> 512,65 -> 514,65 -> 514,64 -> 514,65
533,110 -> 537,110
503,125 -> 507,125
535,148 -> 539,148
507,88 -> 512,88
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
510,65 -> 510,57 -> 510,65 -> 512,65 -> 512,63 -> 512,65 -> 514,65 -> 514,64 -> 514,65
524,108 -> 528,108
490,39 -> 490,42 -> 483,42 -> 483,49 -> 503,49 -> 503,42 -> 495,42 -> 495,39
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
544,150 -> 548,150
524,86 -> 529,86
512,80 -> 517,80
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
490,39 -> 490,42 -> 483,42 -> 483,49 -> 503,49 -> 503,42 -> 495,42 -> 495,39
521,106 -> 525,106
522,153 -> 522,156 -> 515,156 -> 515,164 -> 533,164 -> 533,156 -> 528,156 -> 528,153
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
503,86 -> 508,86
497,23 -> 497,14 -> 497,23 -> 499,23 -> 499,16 -> 499,23 -> 501,23 -> 501,17 -> 501,23
509,125 -> 513,125
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
532,146 -> 536,146
516,82 -> 521,82
510,65 -> 510,57 -> 510,65 -> 512,65 -> 512,63 -> 512,65 -> 514,65 -> 514,64 -> 514,65
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
509,129 -> 513,129
541,148 -> 545,148
490,39 -> 490,42 -> 483,42 -> 483,49 -> 503,49 -> 503,42 -> 495,42 -> 495,39
494,132 -> 494,134 -> 487,134 -> 487,139 -> 502,139 -> 502,134 -> 499,134 -> 499,132
497,23 -> 497,14 -> 497,23 -> 499,23 -> 499,16 -> 499,23 -> 501,23 -> 501,17 -> 501,23
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
511,113 -> 511,116 -> 509,116 -> 509,120 -> 522,120 -> 522,116 -> 516,116 -> 516,113
490,39 -> 490,42 -> 483,42 -> 483,49 -> 503,49 -> 503,42 -> 495,42 -> 495,39
541,144 -> 545,144
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
511,113 -> 511,116 -> 509,116 -> 509,120 -> 522,120 -> 522,116 -> 516,116 -> 516,113
538,150 -> 542,150
510,65 -> 510,57 -> 510,65 -> 512,65 -> 512,63 -> 512,65 -> 514,65 -> 514,64 -> 514,65
511,113 -> 511,116 -> 509,116 -> 509,120 -> 522,120 -> 522,116 -> 516,116 -> 516,113
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
515,129 -> 519,129
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
510,65 -> 510,57 -> 510,65 -> 512,65 -> 512,63 -> 512,65 -> 514,65 -> 514,64 -> 514,65
522,153 -> 522,156 -> 515,156 -> 515,164 -> 533,164 -> 533,156 -> 528,156 -> 528,153
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
510,86 -> 515,86
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
494,132 -> 494,134 -> 487,134 -> 487,139 -> 502,139 -> 502,134 -> 499,134 -> 499,132
521,88 -> 526,88
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
510,65 -> 510,57 -> 510,65 -> 512,65 -> 512,63 -> 512,65 -> 514,65 -> 514,64 -> 514,65
497,129 -> 501,129
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
532,150 -> 536,150
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
547,148 -> 551,148
493,36 -> 493,33 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,33 -> 497,36 -> 499,36 -> 499,35 -> 499,36 -> 501,36 -> 501,34 -> 501,36 -> 503,36 -> 503,34 -> 503,36 -> 505,36 -> 505,26 -> 505,36 -> 507,36 -> 507,31 -> 507,36 -> 509,36 -> 509,33 -> 509,36
518,108 -> 522,108
544,146 -> 548,146
527,101 -> 527,99 -> 527,101 -> 529,101 -> 529,98 -> 529,101 -> 531,101 -> 531,92 -> 531,101 -> 533,101 -> 533,91 -> 533,101 -> 535,101 -> 535,99 -> 535,101 -> 537,101 -> 537,100 -> 537,101 -> 539,101 -> 539,93 -> 539,101
538,142 -> 542,142
`
