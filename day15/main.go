package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/maps"
	"github.com/mbark/advent-of-code-2022/util"
	"strings"
)

const testInput = `
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func main() {
	var sensors []sensor
	input := map[string]struct {
		y     int
		max   int
		input string
	}{
		"test": {y: 10, max: 20, input: testInput},
		"real": {y: 2000000, max: 4000000, input: Input},
	}

	in := input["real"]
	for _, l := range util.ReadInput(in.input, "\n") {
		l = strings.ReplaceAll(l, ",", "")
		l = strings.ReplaceAll(l, ":", "")
		split := strings.Split(l, " ")
		getInt := func(i int) int { return util.ParseInt[int](strings.Split(split[i], "=")[1]) }

		at := maps.Coordinate{X: getInt(2), Y: getInt(3)}
		beacon := maps.Coordinate{X: getInt(8), Y: getInt(9)}
		sensors = append(sensors, sensor{
			at:       at,
			beacon:   beacon,
			distance: at.ManhattanDistance(beacon),
		})
	}

	fmt.Printf("first: %d\n", first(sensors, in.y))
	fmt.Printf("second: %d\n", second(sensors, in.max))
}

type sensor struct {
	at       maps.Coordinate
	beacon   maps.Coordinate
	distance int
}

func (s sensor) String() string {
	return fmt.Sprintf("Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
		s.at.X, s.at.Y, s.beacon.X, s.beacon.Y)
}

func first(sensors []sensor, y int) int {
	covered := make(map[maps.Coordinate]bool)

	for _, s := range sensors {
		distanceToY := maps.Coordinate{X: s.at.X, Y: y}.ManhattanDistance(s.at)
		distanceToBeacon := s.at.ManhattanDistance(s.beacon)

		for i := 0; i <= distanceToBeacon-distanceToY; i++ {
			covered[maps.Coordinate{X: s.at.X + i, Y: y}] = true
			covered[maps.Coordinate{X: s.at.X - i, Y: y}] = true
		}
	}

	for _, s := range sensors {
		delete(covered, s.beacon)
	}

	return len(covered)
}

func second(sensors []sensor, max int) int {
	isIt := func(at maps.Coordinate) bool {
		if at.X > max || at.X < 0 || at.Y > max || at.Y < 0 {
			return false
		}

		for _, s := range sensors {
			if at.ManhattanDistance(s.at) <= s.distance {
				return false
			}
		}

		return true
	}
	getFreq := func(at maps.Coordinate) int {
		return 4000000*at.X + at.Y
	}

	for _, s := range sensors {
		distance := s.distance
		at := maps.Coordinate{X: s.at.X, Y: s.at.Y + distance + 1}
		if isIt(at) {
			return getFreq(at)
		}

		// up right
		for at.Y > s.at.Y {
			at = maps.Coordinate{X: at.X + 1, Y: at.Y - 1}
			if isIt(at) {
				return getFreq(at)
			}
		}
		// up left
		for at.Y > s.at.Y-distance-1 {
			at = maps.Coordinate{X: at.X - 1, Y: at.Y - 1}
			if isIt(at) {
				return getFreq(at)
			}
		}
		// down left
		for at.Y < s.at.Y {
			at = maps.Coordinate{X: at.X - 1, Y: at.Y + 1}
			if isIt(at) {
				return getFreq(at)
			}
		}
		// down right
		for at.Y < s.at.Y+distance+1 {
			at = maps.Coordinate{X: at.X + 1, Y: at.Y + 1}
			if isIt(at) {
				return getFreq(at)
			}
		}
	}

	return 0
}

const Input = `Sensor at x=13820, y=3995710: closest beacon is at x=1532002, y=3577287
Sensor at x=3286002, y=2959504: closest beacon is at x=3931431, y=2926694
Sensor at x=3654160, y=2649422: closest beacon is at x=3702627, y=2598480
Sensor at x=3702414, y=2602790: closest beacon is at x=3702627, y=2598480
Sensor at x=375280, y=2377181: closest beacon is at x=2120140, y=2591883
Sensor at x=3875726, y=2708666: closest beacon is at x=3931431, y=2926694
Sensor at x=3786107, y=2547075: closest beacon is at x=3702627, y=2598480
Sensor at x=2334266, y=3754737: closest beacon is at x=2707879, y=3424224
Sensor at x=1613400, y=1057722: closest beacon is at x=1686376, y=-104303
Sensor at x=3305964, y=2380628: closest beacon is at x=3702627, y=2598480
Sensor at x=1744420, y=3927424: closest beacon is at x=1532002, y=3577287
Sensor at x=3696849, y=2604845: closest beacon is at x=3702627, y=2598480
Sensor at x=2357787, y=401688: closest beacon is at x=1686376, y=-104303
Sensor at x=2127900, y=1984887: closest beacon is at x=2332340, y=2000000
Sensor at x=3705551, y=2604421: closest beacon is at x=3702627, y=2598480
Sensor at x=1783014, y=2978242: closest beacon is at x=2120140, y=2591883
Sensor at x=2536648, y=2910642: closest beacon is at x=2707879, y=3424224
Sensor at x=3999189, y=2989409: closest beacon is at x=3931431, y=2926694
Sensor at x=3939169, y=2382534: closest beacon is at x=3702627, y=2598480
Sensor at x=2792378, y=2002602: closest beacon is at x=2332340, y=2000000
Sensor at x=3520934, y=3617637: closest beacon is at x=2707879, y=3424224
Sensor at x=2614525, y=1628105: closest beacon is at x=2332340, y=2000000
Sensor at x=2828931, y=3996545: closest beacon is at x=2707879, y=3424224
Sensor at x=2184699, y=2161391: closest beacon is at x=2332340, y=2000000
Sensor at x=2272873, y=1816621: closest beacon is at x=2332340, y=2000000
Sensor at x=1630899, y=3675405: closest beacon is at x=1532002, y=3577287
Sensor at x=3683190, y=2619409: closest beacon is at x=3702627, y=2598480
Sensor at x=180960, y=185390: closest beacon is at x=187063, y=-1440697
Sensor at x=1528472, y=3321640: closest beacon is at x=1532002, y=3577287
Sensor at x=3993470, y=2905566: closest beacon is at x=3931431, y=2926694
Sensor at x=1684313, y=20931: closest beacon is at x=1686376, y=-104303
Sensor at x=2547761, y=2464195: closest beacon is at x=2120140, y=2591883
Sensor at x=3711518, y=845968: closest beacon is at x=3702627, y=2598480
Sensor at x=3925049, y=2897039: closest beacon is at x=3931431, y=2926694
Sensor at x=1590740, y=3586256: closest beacon is at x=1532002, y=3577287
Sensor at x=1033496, y=3762565: closest beacon is at x=1532002, y=3577287
`
