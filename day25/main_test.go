package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestToSnafu(t *testing.T) {
	tests := map[int]string{
		1:         "1",
		2:         "2",
		3:         "1=",
		4:         "1-",
		5:         "10",
		6:         "11",
		7:         "12",
		8:         "2=",
		9:         "2-",
		10:        "20",
		15:        "1=0",
		20:        "1-0",
		2022:      "1=11-2",
		12345:     "1-0---0",
		4890:      "2=-1=0",
		314159265: "1121-1110-1=0",
	}

	var keys []int
	for i := range tests {
		keys = append(keys, i)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, d := range keys {
		t.Run(fmt.Sprintf("%d", d), func(t *testing.T) {
			actual := ToSnafu(d)
			if actual != tests[d] {
				t.Fatalf("expected %s but got %s", tests[d], actual)
			}
		})
	}
}
