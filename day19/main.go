package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/util"
	"strings"
)

const testInput = `
Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian. 
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

type blueprint struct {
	index    int
	ore      oreRobot
	clay     clayRobot
	obsidian obsidianRobot
	geode    geodeRobot
}

func (b blueprint) String() string {
	return fmt.Sprintf(`Blueprint %d: 
	Each ore robot costs %d ore. 
	Each clay robot costs %d ore.
	Each obsidian robot costs %d ore and %d clay.
	Each geode robot costs %d ore and %d obsidian.
`,
		b.index,
		b.ore.ore,
		b.clay.ore,
		b.obsidian.ore, b.obsidian.clay,
		b.geode.ore, b.geode.obsidian)
}

type oreRobot struct{ ore int }
type clayRobot struct{ ore int }
type obsidianRobot struct{ ore, clay int }
type geodeRobot struct{ ore, obsidian int }

func main() {
	var blueprints []blueprint
	for _, l := range util.ReadInput(Input, "\n") {
		split := strings.Split(l, ":")
		index := strings.TrimPrefix(split[0], "Blueprint ")
		split = strings.Split(split[1], ".")

		ore := strings.Split(split[0], " ")[5]
		clay := strings.Split(split[1], " ")[5]
		obs := strings.Split(split[2], " ")
		obs1, obs2 := obs[5], obs[8]
		geode := strings.Split(split[3], " ")
		geode1, geode2 := geode[5], geode[8]

		blueprints = append(blueprints, blueprint{
			index:    util.ParseInt[int](index),
			ore:      oreRobot{ore: util.ParseInt[int](ore)},
			clay:     clayRobot{ore: util.ParseInt[int](clay)},
			obsidian: obsidianRobot{ore: util.ParseInt[int](obs1), clay: util.ParseInt[int](obs2)},
			geode:    geodeRobot{ore: util.ParseInt[int](geode1), obsidian: util.ParseInt[int](geode2)},
		})
	}

	fmt.Printf("first: %d\n", first(blueprints))
	fmt.Printf("second: %d\n", second(blueprints))
}

func first(blueprints []blueprint) int {
	var sum int
	for _, b := range blueprints {
		memo = make(map[int]int)
		hypothetical = make(map[int]int)
		initial := state{
			time:        0,
			oreBot:      1,
			clayBot:     0,
			obsidianBot: 0,
			geodeBot:    0,
			ore:         0,
			clay:        0,
			obsidian:    0,
			geode:       0,
		}
		s := solve(b, 24, initial)
		sum += b.index * s
	}

	return sum
}

func second(blueprints []blueprint) int {
	sum := 1
	for _, b := range blueprints[:util.MinInt(3, len(blueprints))] {
		memo = make(map[int]int)
		hypothetical = make(map[int]int)
		initial := state{
			time:        0,
			oreBot:      1,
			clayBot:     0,
			obsidianBot: 0,
			geodeBot:    0,
			ore:         0,
			clay:        0,
			obsidian:    0,
			geode:       0,
		}
		s := solve(b, 32, initial)
		sum *= s
	}

	return sum
}

type state struct {
	time                                   int
	oreBot, clayBot, obsidianBot, geodeBot int
	ore, clay, obsidian, geode             int
}

func (s state) hash() int {
	var val int
	mul := 1

	for _, n := range []int{s.time, s.oreBot, s.clayBot, s.obsidianBot, s.geodeBot, s.ore, s.clay, s.obsidian, s.geode} {
		val += n * mul
		mul *= 100
	}

	return val
}

func (s state) String() string {
	return fmt.Sprintf("time %d; bots: %d ore, %d clay, %d obsidian, %d geode; resources: %d ore, %d clay, %d obsidian, %d geode",
		s.time, s.oreBot, s.clayBot, s.obsidianBot, s.geodeBot, s.ore, s.clay, s.obsidian, s.geode)
}

func hypotheticalGeodes(b blueprint, s state, maxTime int) int {
	geodes := s.geode

	timeLeft := maxTime - s.time
	guaranteed := s.geodeBot * timeLeft
	added := timeLeft * (timeLeft - 1) / 2

	return geodes + guaranteed + added
}

func nextBot1(needs, has, bots int) int {
	needed := needs - util.MinInt(has, needs)
	takes := needed / bots
	if needed%bots > 0 {
		takes += 1
	}

	return takes
}

func nextBot2(needs1, has1, bots1, needs2, has2, bots2 int) int {
	takes1 := nextBot1(needs1, has1, bots1)
	takes2 := nextBot1(needs2, has2, bots2)

	return util.MaxInt(takes1, takes2)
}

var memo = make(map[int]int)
var hypothetical = make(map[int]int)
var loops int

func needsAtMost(b blueprint) (int, int, int) {
	var ore int
	ore = util.MaxInt(ore, b.ore.ore)
	ore = util.MaxInt(ore, b.clay.ore)
	ore = util.MaxInt(ore, b.obsidian.ore)
	ore = util.MaxInt(ore, b.geode.ore)

	return ore, b.obsidian.clay, b.geode.obsidian
}

func cap(has, bots, maxUse int) int {
	if has < maxUse {
		return has
	}

	if bots >= maxUse {
		return maxUse
	}

	return has
}

func solve(b blueprint, maxTime int, s state) (res int) {
	if s.time >= maxTime {
		return s.geode
	}

	h := hypotheticalGeodes(b, s, maxTime)
	if hypothetical[s.time] >= h || h == 0 {
		return 0
	}
	defer func() { hypothetical[s.time] = res }()

	mostOre, mostClay, mostObsidian := needsAtMost(b)
	hash := s.hash()
	if v, ok := memo[hash]; ok {
		return v
	}
	defer func() { memo[hash] = res }()

	var next []state
	if s.obsidianBot > 0 {
		next = append(next, state{
			time:     nextBot2(b.geode.ore, s.ore, s.oreBot, b.geode.obsidian, s.obsidian, s.obsidianBot),
			geodeBot: 1,
			ore:      -b.geode.ore,
			obsidian: -b.geode.obsidian,
		})
	}

	if s.clayBot > 0 && s.time <= maxTime-2 && s.obsidianBot < mostObsidian {
		next = append(next, state{
			time:        nextBot2(b.obsidian.ore, s.ore, s.oreBot, b.obsidian.clay, s.clay, s.clayBot),
			obsidianBot: 1,
			ore:         -b.obsidian.ore,
			clay:        -b.obsidian.clay,
		})
	}

	if s.time <= maxTime-2 && s.clayBot < mostClay {
		next = append(next, state{
			time:    nextBot1(b.clay.ore, s.ore, s.oreBot),
			clayBot: 1,
			ore:     -b.clay.ore,
		})
	}

	if s.time <= maxTime-2 && s.oreBot < mostOre {
		next = append(next, state{
			time:   nextBot1(b.ore.ore, s.ore, s.oreBot),
			oreBot: 1,
			ore:    -b.ore.ore,
		})
	}

	for i := range next {
		next[i].oreBot += s.oreBot
		next[i].clayBot += s.clayBot
		next[i].obsidianBot += s.obsidianBot
		next[i].geodeBot += s.geodeBot

		next[i].time += 1
		next[i].ore += s.ore + next[i].time*s.oreBot
		next[i].clay += s.clay + next[i].time*s.clayBot
		next[i].obsidian += s.obsidian + next[i].time*s.obsidianBot
		next[i].geode += s.geode + next[i].time*s.geodeBot

		next[i].time += s.time

		next[i].ore = cap(next[i].ore, next[i].oreBot, mostOre)
		next[i].clay = cap(next[i].clay, next[i].clayBot, mostClay)
		next[i].obsidian = cap(next[i].obsidian, next[i].obsidianBot, mostObsidian)
	}

	max := s.geode + (maxTime-s.time)*s.geodeBot
	for _, n := range next {
		if n.time >= maxTime {
			continue
		}

		m := s.geode + (maxTime-n.time)*s.geodeBot
		if h := hypothetical[n.time]; h < m {
			hypothetical[n.time] = m
		}

		if geodes := solve(b, maxTime, n); geodes > max {
			max = geodes
		}
	}

	return max
}

const Input = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 11 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 20 clay. Each geode robot costs 2 ore and 8 obsidian.
Blueprint 3: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 15 clay. Each geode robot costs 4 ore and 16 obsidian.
Blueprint 4: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 15 clay. Each geode robot costs 3 ore and 16 obsidian.
Blueprint 5: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 4 ore and 8 clay. Each geode robot costs 3 ore and 7 obsidian.
Blueprint 6: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 18 clay. Each geode robot costs 3 ore and 8 obsidian.
Blueprint 7: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 20 clay. Each geode robot costs 3 ore and 15 obsidian.
Blueprint 8: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 15 clay. Each geode robot costs 3 ore and 20 obsidian.
Blueprint 9: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 4 ore and 9 obsidian.
Blueprint 10: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 9 clay. Each geode robot costs 3 ore and 15 obsidian.
Blueprint 11: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 9 clay. Each geode robot costs 3 ore and 9 obsidian.
Blueprint 12: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 16 clay. Each geode robot costs 3 ore and 14 obsidian.
Blueprint 13: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 6 clay. Each geode robot costs 2 ore and 10 obsidian.
Blueprint 14: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 7 clay. Each geode robot costs 3 ore and 8 obsidian.
Blueprint 15: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 17 clay. Each geode robot costs 3 ore and 16 obsidian.
Blueprint 16: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 15 clay. Each geode robot costs 4 ore and 17 obsidian.
Blueprint 17: Each ore robot costs 3 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 9 clay. Each geode robot costs 3 ore and 7 obsidian.
Blueprint 18: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 8 clay. Each geode robot costs 3 ore and 19 obsidian.
Blueprint 19: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 19 clay. Each geode robot costs 4 ore and 15 obsidian.
Blueprint 20: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 3 ore and 19 obsidian.
Blueprint 21: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 15 clay. Each geode robot costs 3 ore and 7 obsidian.
Blueprint 22: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 7 clay. Each geode robot costs 2 ore and 16 obsidian.
Blueprint 23: Each ore robot costs 2 ore. Each clay robot costs 2 ore. Each obsidian robot costs 2 ore and 7 clay. Each geode robot costs 2 ore and 14 obsidian.
Blueprint 24: Each ore robot costs 3 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 19 clay. Each geode robot costs 3 ore and 19 obsidian.
Blueprint 25: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 11 clay. Each geode robot costs 3 ore and 14 obsidian.
Blueprint 26: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 7 clay. Each geode robot costs 3 ore and 9 obsidian.
Blueprint 27: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 16 clay. Each geode robot costs 2 ore and 9 obsidian.
Blueprint 28: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 19 clay. Each geode robot costs 4 ore and 11 obsidian.
Blueprint 29: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 5 clay. Each geode robot costs 2 ore and 10 obsidian.
Blueprint 30: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 20 clay. Each geode robot costs 2 ore and 17 obsidian.
`
