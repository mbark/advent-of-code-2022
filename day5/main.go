package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2022/util"
	"strings"
)

const testInput = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func main() {
	in := util.ReadInput(Input, "\n\n")
	crates, procs := in[0], in[1]

	crateRows := strings.Split(crates, "\n")

	stacks1 := make([]stack, (len(crateRows[0])+1)/4)
	stacks2 := make([]stack, (len(crateRows[0])+1)/4)
	for _, row := range crateRows[:len(crateRows)-1] {
		for i := 0; i < len(row)+1; i += 4 {
			box := row[i+1]
			if box == ' ' {
				continue
			}

			stacks1[i/4] = stacks1[i/4].prepend(box)
			stacks2[i/4] = stacks2[i/4].prepend(box)
		}
	}

	var procedures []procedure
	for _, row := range strings.Split(procs, "\n") {
		s := strings.Split(row, " ")
		procedures = append(procedures, procedure{
			count: util.ParseInt[int](s[1]),
			from:  util.ParseInt[int](s[3]) - 1,
			to:    util.ParseInt[int](s[5]) - 1,
		})
	}

	fmt.Printf("first: %s\n", solve(stacks1, procedures, popOne))
	fmt.Printf("second: %s\n", solve(stacks2, procedures, popAll))
}

func solve(stacks []stack, procedures []procedure, fn func(p procedure, stacks []stack) []stack) string {
	for _, p := range procedures {
		stacks = fn(p, stacks)
	}

	var res []byte
	for _, s := range stacks {
		res = append(res, s[len(s)-1])
	}

	return string(res)
}

func popOne(p procedure, stacks []stack) []stack {
	for i := 0; i < p.count; i++ {
		lift, remaining := stacks[p.from].pop(1)
		stacks[p.to] = append(stacks[p.to], lift...)
		stacks[p.from] = remaining
	}
	return stacks
}

func popAll(p procedure, stacks []stack) []stack {
	lift, remaining := stacks[p.from].pop(p.count)
	stacks[p.to] = append(stacks[p.to], lift...)
	stacks[p.from] = remaining
	return stacks
}

type procedure struct{ count, from, to int }

func (p procedure) String() string {
	return fmt.Sprintf("move %d from %d to %d", p.count, p.from+1, p.to+1)
}

type stack []byte

func (s stack) push(st stack) stack {
	return append(s, st...)
}

func (s stack) pop(i int) (stack, stack) {
	from := len(s) - i
	return s[from:], s[:from]
}

func (s stack) prepend(b byte) stack {
	return append([]byte{b}, s...)
}

func (s stack) String() string {
	var boxes []string
	for _, b := range s {
		boxes = append(boxes, "["+string(b)+"]")
	}
	return strings.Join(boxes, " ")
}

const Input = `
[N] [G]                     [Q]    
[H] [B]         [B] [R]     [H]    
[S] [N]     [Q] [M] [T]     [Z]    
[J] [T]     [R] [V] [H]     [R] [S]
[F] [Q]     [W] [T] [V] [J] [V] [M]
[W] [P] [V] [S] [F] [B] [Q] [J] [H]
[T] [R] [Q] [B] [D] [D] [B] [N] [N]
[D] [H] [L] [N] [N] [M] [D] [D] [B]
 1   2   3   4   5   6   7   8   9 

move 3 from 1 to 2
move 1 from 7 to 1
move 1 from 6 to 5
move 5 from 5 to 9
move 2 from 5 to 2
move 1 from 6 to 8
move 1 from 5 to 7
move 5 from 4 to 6
move 1 from 7 to 6
move 1 from 2 to 4
move 5 from 2 to 6
move 2 from 1 to 5
move 2 from 1 to 9
move 16 from 6 to 4
move 6 from 8 to 3
move 7 from 2 to 4
move 5 from 9 to 3
move 1 from 1 to 4
move 1 from 1 to 3
move 3 from 7 to 4
move 2 from 5 to 4
move 31 from 4 to 8
move 22 from 8 to 4
move 9 from 3 to 6
move 7 from 9 to 5
move 4 from 5 to 6
move 6 from 3 to 2
move 2 from 6 to 7
move 5 from 2 to 7
move 1 from 2 to 4
move 1 from 7 to 5
move 4 from 5 to 4
move 2 from 6 to 9
move 2 from 4 to 6
move 7 from 6 to 4
move 2 from 6 to 1
move 1 from 6 to 8
move 8 from 8 to 1
move 1 from 7 to 6
move 4 from 1 to 5
move 9 from 4 to 8
move 4 from 1 to 7
move 3 from 5 to 3
move 2 from 1 to 9
move 1 from 3 to 2
move 1 from 9 to 8
move 1 from 2 to 1
move 1 from 1 to 8
move 1 from 5 to 1
move 2 from 3 to 1
move 2 from 6 to 9
move 19 from 4 to 1
move 4 from 4 to 2
move 6 from 1 to 4
move 1 from 2 to 4
move 4 from 4 to 3
move 7 from 7 to 3
move 7 from 8 to 2
move 2 from 7 to 4
move 3 from 2 to 1
move 8 from 8 to 2
move 3 from 9 to 1
move 2 from 9 to 1
move 10 from 2 to 7
move 4 from 3 to 1
move 1 from 8 to 3
move 1 from 4 to 5
move 1 from 3 to 6
move 1 from 2 to 1
move 10 from 1 to 3
move 1 from 4 to 7
move 1 from 6 to 4
move 7 from 3 to 2
move 5 from 2 to 8
move 11 from 7 to 2
move 3 from 4 to 3
move 1 from 4 to 3
move 5 from 8 to 9
move 17 from 2 to 4
move 11 from 1 to 5
move 4 from 1 to 3
move 5 from 9 to 2
move 4 from 2 to 1
move 3 from 5 to 7
move 6 from 5 to 3
move 1 from 5 to 8
move 6 from 1 to 8
move 3 from 8 to 5
move 1 from 1 to 4
move 1 from 7 to 2
move 15 from 3 to 4
move 1 from 1 to 3
move 10 from 3 to 9
move 2 from 7 to 4
move 1 from 2 to 8
move 21 from 4 to 9
move 1 from 2 to 3
move 1 from 8 to 1
move 9 from 4 to 2
move 1 from 1 to 5
move 5 from 2 to 7
move 2 from 8 to 5
move 1 from 8 to 1
move 2 from 2 to 8
move 2 from 4 to 9
move 24 from 9 to 5
move 3 from 4 to 1
move 2 from 2 to 5
move 12 from 5 to 1
move 10 from 1 to 5
move 23 from 5 to 6
move 8 from 9 to 1
move 3 from 8 to 1
move 1 from 1 to 2
move 1 from 3 to 7
move 11 from 6 to 1
move 1 from 2 to 4
move 6 from 6 to 8
move 4 from 6 to 7
move 1 from 7 to 3
move 1 from 3 to 4
move 23 from 1 to 8
move 1 from 4 to 2
move 1 from 2 to 1
move 1 from 6 to 7
move 6 from 5 to 3
move 1 from 7 to 8
move 1 from 1 to 8
move 1 from 9 to 3
move 6 from 7 to 2
move 3 from 5 to 9
move 5 from 2 to 3
move 28 from 8 to 3
move 4 from 1 to 9
move 5 from 9 to 5
move 2 from 8 to 5
move 1 from 9 to 4
move 2 from 7 to 5
move 1 from 4 to 2
move 1 from 4 to 8
move 2 from 8 to 3
move 6 from 5 to 2
move 1 from 7 to 2
move 39 from 3 to 2
move 2 from 3 to 8
move 1 from 9 to 6
move 2 from 2 to 9
move 2 from 9 to 6
move 1 from 8 to 1
move 1 from 1 to 6
move 5 from 6 to 9
move 2 from 5 to 8
move 20 from 2 to 4
move 2 from 4 to 8
move 2 from 8 to 3
move 3 from 3 to 1
move 22 from 2 to 5
move 2 from 9 to 1
move 3 from 1 to 7
move 1 from 2 to 6
move 1 from 2 to 9
move 1 from 1 to 8
move 2 from 7 to 9
move 1 from 6 to 8
move 1 from 2 to 7
move 1 from 1 to 3
move 1 from 9 to 8
move 1 from 8 to 5
move 3 from 8 to 7
move 3 from 7 to 8
move 15 from 4 to 1
move 1 from 4 to 3
move 10 from 1 to 6
move 3 from 8 to 1
move 5 from 9 to 4
move 7 from 5 to 1
move 4 from 6 to 3
move 15 from 5 to 2
move 4 from 6 to 4
move 7 from 2 to 1
move 6 from 4 to 6
move 1 from 5 to 9
move 1 from 5 to 7
move 1 from 3 to 5
move 11 from 1 to 8
move 3 from 4 to 6
move 4 from 1 to 5
move 1 from 2 to 5
move 2 from 8 to 3
move 11 from 6 to 1
move 1 from 3 to 7
move 1 from 9 to 8
move 6 from 5 to 8
move 3 from 8 to 4
move 1 from 4 to 5
move 3 from 3 to 1
move 9 from 8 to 2
move 2 from 1 to 5
move 11 from 2 to 5
move 1 from 3 to 6
move 2 from 8 to 5
move 3 from 4 to 6
move 1 from 8 to 3
move 2 from 1 to 9
move 1 from 3 to 8
move 16 from 5 to 7
move 3 from 1 to 6
move 1 from 3 to 5
move 1 from 6 to 7
move 1 from 9 to 4
move 1 from 5 to 4
move 1 from 3 to 2
move 1 from 1 to 2
move 3 from 4 to 9
move 1 from 2 to 7
move 2 from 8 to 3
move 6 from 2 to 8
move 11 from 1 to 3
move 6 from 3 to 1
move 4 from 3 to 2
move 2 from 3 to 1
move 1 from 1 to 3
move 4 from 8 to 4
move 4 from 8 to 2
move 11 from 7 to 2
move 9 from 7 to 5
move 1 from 7 to 3
move 4 from 5 to 7
move 14 from 2 to 3
move 17 from 3 to 7
move 2 from 5 to 2
move 1 from 5 to 7
move 1 from 5 to 6
move 4 from 6 to 7
move 8 from 1 to 2
move 2 from 6 to 4
move 1 from 6 to 8
move 6 from 4 to 1
move 1 from 8 to 5
move 6 from 7 to 8
move 5 from 8 to 3
move 12 from 2 to 1
move 1 from 8 to 4
move 4 from 3 to 1
move 4 from 2 to 4
move 3 from 9 to 3
move 3 from 3 to 2
move 1 from 3 to 2
move 3 from 4 to 1
move 2 from 5 to 7
move 22 from 1 to 8
move 17 from 8 to 6
move 21 from 7 to 6
move 3 from 2 to 8
move 3 from 1 to 5
move 3 from 5 to 2
move 2 from 4 to 6
move 7 from 6 to 5
move 1 from 9 to 4
move 14 from 6 to 4
move 5 from 8 to 3
move 1 from 6 to 3
move 3 from 3 to 9
move 2 from 9 to 1
move 2 from 7 to 1
move 16 from 6 to 8
move 2 from 6 to 7
move 1 from 2 to 7
move 1 from 3 to 8
move 7 from 4 to 1
move 2 from 7 to 2
move 4 from 4 to 7
move 5 from 2 to 4
move 1 from 7 to 3
move 3 from 5 to 8
move 1 from 7 to 5
move 12 from 1 to 6
move 3 from 7 to 2
move 7 from 4 to 2
move 3 from 3 to 2
move 1 from 4 to 2
move 1 from 9 to 8
move 8 from 6 to 8
move 12 from 2 to 4
move 5 from 5 to 2
move 11 from 4 to 9
move 3 from 6 to 3
move 2 from 4 to 2
move 4 from 2 to 6
move 5 from 2 to 8
move 12 from 8 to 4
move 20 from 8 to 5
move 13 from 5 to 3
move 1 from 8 to 5
move 5 from 5 to 9
move 16 from 9 to 1
move 9 from 4 to 5
move 12 from 3 to 9
move 5 from 6 to 5
move 9 from 9 to 7
move 14 from 1 to 4
move 14 from 4 to 1
move 15 from 5 to 7
move 4 from 8 to 2
move 3 from 4 to 3
move 3 from 1 to 8
move 1 from 5 to 9
move 1 from 5 to 3
move 3 from 9 to 8
move 4 from 3 to 4
move 1 from 4 to 6
move 20 from 7 to 2
move 2 from 3 to 8
move 3 from 7 to 2
move 4 from 2 to 1
move 1 from 6 to 7
move 3 from 4 to 2
move 2 from 2 to 3
move 4 from 3 to 4
move 1 from 8 to 1
move 3 from 8 to 1
move 2 from 7 to 8
move 1 from 4 to 5
move 14 from 2 to 5
move 6 from 1 to 5
move 1 from 4 to 3
move 15 from 1 to 4
move 1 from 8 to 2
move 1 from 9 to 5
move 4 from 8 to 7
move 13 from 5 to 6
move 1 from 8 to 1
move 2 from 7 to 9
move 12 from 6 to 4
move 1 from 3 to 6
move 1 from 1 to 6
move 4 from 5 to 2
move 5 from 5 to 6
move 2 from 6 to 2
move 1 from 7 to 5
move 2 from 6 to 9
move 1 from 5 to 9
move 16 from 2 to 5
move 17 from 4 to 1
move 3 from 1 to 3
move 1 from 2 to 6
move 2 from 6 to 1
move 3 from 3 to 1
move 14 from 1 to 8
move 3 from 5 to 2
move 4 from 8 to 2
move 3 from 4 to 5
move 15 from 5 to 3
move 1 from 7 to 6
move 3 from 1 to 8
move 2 from 3 to 7
move 1 from 1 to 2
move 1 from 7 to 6
move 4 from 2 to 8
move 2 from 6 to 2
move 1 from 7 to 6
move 3 from 8 to 2
move 12 from 8 to 6
move 1 from 5 to 6
move 3 from 2 to 5
move 2 from 2 to 5
move 4 from 6 to 5
move 4 from 3 to 5
move 1 from 8 to 4
move 11 from 6 to 4
move 6 from 3 to 1
move 2 from 9 to 8
move 20 from 4 to 5
move 1 from 4 to 9
move 2 from 3 to 8
move 1 from 3 to 8
move 17 from 5 to 8
move 5 from 5 to 9
move 9 from 5 to 1
move 2 from 6 to 7
move 23 from 8 to 2
move 2 from 7 to 5
move 3 from 9 to 4
move 16 from 2 to 4
move 11 from 1 to 8
move 4 from 5 to 8
move 11 from 2 to 6
move 2 from 6 to 1
move 5 from 9 to 5
move 5 from 5 to 6
move 5 from 8 to 6
move 1 from 6 to 7
move 7 from 8 to 1
move 12 from 1 to 2
move 1 from 9 to 5
move 1 from 1 to 3
move 1 from 1 to 4
move 1 from 5 to 3
move 1 from 3 to 6
move 1 from 8 to 2
move 18 from 6 to 2
move 1 from 6 to 2
move 2 from 8 to 3
move 3 from 3 to 8
move 18 from 4 to 9
move 11 from 9 to 2
move 2 from 9 to 6
move 2 from 4 to 1
move 1 from 1 to 5
move 1 from 5 to 4
move 1 from 4 to 8
move 42 from 2 to 1
move 3 from 9 to 3
move 1 from 8 to 1
move 1 from 3 to 4
move 3 from 8 to 7
move 1 from 4 to 1
move 2 from 3 to 2
move 17 from 1 to 6
move 15 from 6 to 3
move 2 from 9 to 7
move 1 from 3 to 6
move 2 from 7 to 6
move 2 from 2 to 4
move 1 from 2 to 3
move 1 from 4 to 9
move 1 from 4 to 1
move 1 from 6 to 3
move 20 from 1 to 9
move 6 from 1 to 9
move 7 from 9 to 3
move 20 from 9 to 1
move 1 from 6 to 7
move 2 from 6 to 7
move 1 from 6 to 5
move 1 from 6 to 8
move 4 from 7 to 3
move 3 from 7 to 2
move 1 from 6 to 4
move 1 from 2 to 1
move 1 from 4 to 9
move 21 from 3 to 2
move 5 from 3 to 8
move 1 from 5 to 1
move 2 from 8 to 7
move 4 from 8 to 3
move 4 from 2 to 5
move 19 from 2 to 3
move 1 from 9 to 2
move 23 from 3 to 2
move 2 from 7 to 4
move 3 from 5 to 9
move 16 from 2 to 1
move 1 from 5 to 4
move 1 from 9 to 3
move 2 from 3 to 8
move 3 from 4 to 6
move 1 from 6 to 2
move 1 from 8 to 6
move 5 from 2 to 6
move 7 from 6 to 5
move 4 from 2 to 6
move 6 from 5 to 9
move 1 from 8 to 4
move 18 from 1 to 9
move 1 from 5 to 2
move 9 from 9 to 4
move 5 from 6 to 3
move 9 from 4 to 1
move 4 from 9 to 2
move 1 from 4 to 8
move 1 from 8 to 3
move 7 from 1 to 8
move 6 from 3 to 2
move 10 from 2 to 9
move 21 from 1 to 8
move 1 from 2 to 8
move 19 from 8 to 4
move 1 from 8 to 3
move 16 from 4 to 8
move 1 from 4 to 2
move 2 from 1 to 5
move 1 from 2 to 3
move 1 from 4 to 5
move 1 from 4 to 8
move 2 from 1 to 3
move 3 from 3 to 2
move 5 from 9 to 1
move 1 from 3 to 4
move 4 from 9 to 4
move 2 from 1 to 9
move 2 from 2 to 5
move 1 from 2 to 7
move 3 from 1 to 7
move 10 from 8 to 6
move 4 from 8 to 5
move 3 from 4 to 3
move 3 from 3 to 4
move 1 from 9 to 8
move 2 from 7 to 2
move 1 from 2 to 1
move 4 from 9 to 3
`
