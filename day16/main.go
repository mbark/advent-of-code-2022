package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/util"
	"sort"
	"strconv"
	"strings"
)

const testInput = `
Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func main() {
	valves := make(map[string]valve)
	for _, l := range util.ReadInput(Input, "\n") {
		split := strings.Split(l, ";")
		vl := strings.Split(split[0], " ")
		tl := strings.TrimPrefix(split[1], " tunnels lead")
		tl = strings.TrimPrefix(tl, " tunnel leads")
		tl = strings.TrimPrefix(tl, " to valves ")
		tl = strings.TrimPrefix(tl, " to valve ")

		flow := util.ParseInt[int](strings.Split(vl[4], "=")[1])
		v := valve{
			name: vl[1],
			flow: flow,
			to:   strings.Split(tl, ", "),
		}
		valves[v.name] = v
	}

	fmt.Printf("first: %d\n", first(valves))
	fmt.Printf("second: %d\n", second(valves))
}

type valve struct {
	name string
	flow int
	to   []string
}

func (v valve) String() string {
	return fmt.Sprintf("%s rate=%d; to: [%s]", v.name, v.flow, strings.Join(v.to, ", "))
}

func first(valves map[string]valve) int {
	memo = make(map[string]int)
	distances := make(map[string]map[string]int)
	for _, v := range valves {
		if v.flow == 0 && v.name != "AA" {
			continue
		}

		s := buildDistanceGraph(v, valves)
		distances[v.name] = s
	}

	var best int
	for v := range distances {
		release := solve(valves[v], distances, make(map[string]bool), valves, 30, distances["AA"][v])
		if release > best {
			best = release
		}
	}

	return best
}

func second(valves map[string]valve) int {
	distances := make(map[string]map[string]int)
	for _, v := range valves {
		if v.flow == 0 && v.name != "AA" {
			continue
		}

		s := buildDistanceGraph(v, valves)
		distances[v.name] = s
	}
	fromAA := distances["AA"]
	delete(distances, "AA")

	var results []result
	for v := range distances {
		ress := solve2(valves[v], distances, make(map[string]int), valves, 26, fromAA[v], 0)
		results = append(results, ress...)
	}

	sort.Slice(results, func(i, j int) bool { return results[i].flow > results[j].flow })

	var best int
	for _, r := range results[:3000] {
		for v := range distances {
			if _, ok := r.open[v]; ok {
				continue
			}

			res := solve2(valves[v], distances, copyMap(r.open), valves, 26, fromAA[v], r.flow)
			if len(res) == 0 {
				continue
			}
			sort.Slice(res, func(i, j int) bool { return res[i].flow > res[j].flow })

			if res[0].flow > best {
				best = res[0].flow
			}
		}
	}

	return best
}

var memo = map[string]int{}

type position struct {
	at   valve
	time int
}

func solve(at valve, distances map[string]map[string]int, open map[string]bool, valves map[string]valve, totalTime, currentTime int) int {
	if currentTime > 30 {
		return 0
	}

	var os []string
	for o := range open {
		os = append(os, o)
	}
	sort.Strings(os)

	memoKey := at.name + "-" + strings.Join(os, ",") + "-" + strconv.Itoa(currentTime)
	if val, ok := memo[memoKey]; ok {
		return val
	}

	currentTime += 1
	flow := (totalTime - currentTime) * at.flow
	open[at.name] = true

	var best int
	for v, dist := range distances[at.name] {
		if open[v] {
			continue
		}

		release := solve(valves[v], distances, copyMap(open), valves, totalTime, currentTime+dist)
		if release > best {
			best = release
		}
	}

	memo[memoKey] = best + flow
	return best + flow
}

type result struct {
	flow int
	open map[string]int
}

var solve2Memo = map[string][]result{}

func solve2(at valve, distances map[string]map[string]int, open map[string]int, valves map[string]valve, totalTime, currentTime, totalFlow int) (ress []result) {
	if currentTime > totalTime {
		return []result{{flow: totalFlow, open: open}}
	}

	var os []string
	for o, when := range open {
		os = append(os, o+":"+strconv.Itoa(when))
	}
	sort.Strings(os)

	memoKey := at.name + "-" + strings.Join(os, ",") + "-" + strconv.Itoa(currentTime)
	if val, ok := solve2Memo[memoKey]; ok {
		return val
	}
	defer func() { solve2Memo[memoKey] = ress }()

	currentTime += 1
	flow := totalFlow + (totalTime-currentTime)*at.flow
	open[at.name] = currentTime

	var results []result
	for v, dist := range distances[at.name] {
		if _, ok := open[v]; ok {
			continue
		}

		res := solve2(valves[v], distances, copyMap(open), valves, totalTime, currentTime+dist, flow)
		results = append(results, res...)
	}

	return results
}

func copyMap[K, V comparable](m map[K]V) map[K]V {
	newm := make(map[K]V, len(m))
	for k, v := range m {
		newm[k] = v
	}
	return newm
}

func buildDistanceGraph(start valve, valves map[string]valve) map[string]int {
	visisted := map[string]bool{start.name: true}
	queue := []valve{start}
	from := make(map[string]valve)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		for _, to := range next.to {
			if visisted[to] {
				continue
			}

			visisted[to] = true
			from[to] = next
			queue = append(queue, valves[to])
		}
	}

	steps := map[string]int{start.name: 0}
	for v, vlv := range valves {
		if vlv.flow == 0 {
			continue
		}

		if v == start.name {
			continue
		}

		var s int
		for at := v; at != start.name; at = from[at].name {
			s += 1
		}

		steps[v] = s
	}

	return steps
}

const Input = `
Valve NQ has flow rate=0; tunnels lead to valves SU, XD
Valve AB has flow rate=0; tunnels lead to valves XD, TE
Valve IA has flow rate=0; tunnels lead to valves CS, WF
Valve WD has flow rate=0; tunnels lead to valves DW, II
Valve XD has flow rate=10; tunnels lead to valves AB, NQ, VT, SC, MU
Valve SL has flow rate=0; tunnels lead to valves RP, DS
Valve FQ has flow rate=15; tunnels lead to valves EI, YC
Valve KF has flow rate=0; tunnels lead to valves FL, QP
Valve QP has flow rate=0; tunnels lead to valves KF, RP
Valve DS has flow rate=0; tunnels lead to valves SL, AA
Valve IK has flow rate=0; tunnels lead to valves XC, AA
Valve HQ has flow rate=0; tunnels lead to valves VM, WV
Valve WR has flow rate=0; tunnels lead to valves WV, HF
Valve HH has flow rate=20; tunnels lead to valves PI, CF, CN, NF, AR
Valve DW has flow rate=19; tunnels lead to valves KD, WD, HS
Valve RP has flow rate=14; tunnels lead to valves SL, QP, BH, LI, WP
Valve EC has flow rate=0; tunnels lead to valves NF, XC
Valve AA has flow rate=0; tunnels lead to valves NH, ES, UC, IK, DS
Valve VM has flow rate=18; tunnel leads to valve HQ
Valve NF has flow rate=0; tunnels lead to valves HH, EC
Valve PS has flow rate=0; tunnels lead to valves AR, SU
Valve IL has flow rate=0; tunnels lead to valves XC, KZ
Valve WP has flow rate=0; tunnels lead to valves CS, RP
Valve WF has flow rate=0; tunnels lead to valves FL, IA
Valve XW has flow rate=0; tunnels lead to valves OL, NL
Valve EH has flow rate=0; tunnels lead to valves UK, YR
Valve UC has flow rate=0; tunnels lead to valves AA, FL
Valve CS has flow rate=3; tunnels lead to valves IA, CN, LD, RJ, WP
Valve AR has flow rate=0; tunnels lead to valves PS, HH
Valve CF has flow rate=0; tunnels lead to valves HH, FL
Valve NH has flow rate=0; tunnels lead to valves AA, LD
Valve RJ has flow rate=0; tunnels lead to valves DJ, CS
Valve XC has flow rate=17; tunnels lead to valves IL, EC, YR, IK, DJ
Valve TE has flow rate=24; tunnels lead to valves AB, YA
Valve CN has flow rate=0; tunnels lead to valves HH, CS
Valve KD has flow rate=0; tunnels lead to valves DW, UK
Valve SC has flow rate=0; tunnels lead to valves EI, XD
Valve MU has flow rate=0; tunnels lead to valves XD, YP
Valve SU has flow rate=22; tunnels lead to valves PS, LI, II, NQ
Valve FL has flow rate=8; tunnels lead to valves KF, WF, CF, UC, HS
Valve OL has flow rate=4; tunnels lead to valves KZ, HF, XW
Valve EI has flow rate=0; tunnels lead to valves FQ, SC
Valve NL has flow rate=0; tunnels lead to valves XW, UK
Valve YP has flow rate=21; tunnels lead to valves YA, MU, YC
Valve BH has flow rate=0; tunnels lead to valves VT, RP
Valve II has flow rate=0; tunnels lead to valves SU, WD
Valve YA has flow rate=0; tunnels lead to valves TE, YP
Valve HS has flow rate=0; tunnels lead to valves FL, DW
Valve DJ has flow rate=0; tunnels lead to valves RJ, XC
Valve KZ has flow rate=0; tunnels lead to valves OL, IL
Valve YR has flow rate=0; tunnels lead to valves EH, XC
Valve UK has flow rate=7; tunnels lead to valves KD, NL, EH
Valve YC has flow rate=0; tunnels lead to valves FQ, YP
Valve ES has flow rate=0; tunnels lead to valves PI, AA
Valve LI has flow rate=0; tunnels lead to valves SU, RP
Valve LD has flow rate=0; tunnels lead to valves NH, CS
Valve VT has flow rate=0; tunnels lead to valves BH, XD
Valve PI has flow rate=0; tunnels lead to valves ES, HH
Valve WV has flow rate=11; tunnels lead to valves WR, HQ
Valve HF has flow rate=0; tunnels lead to valves OL, WR
`
