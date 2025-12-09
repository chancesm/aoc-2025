package day05

import (
	"strings"
	"testing"
)

const sample = `
3-5
10-14
16-20
12-18
13-17

1
5
8
11
17
32
`

func TestPart1_Sample(t *testing.T) {
	got := (Day05{}).Part1(strings.TrimSpace(sample))
	const want = "3"
	if got != want {
		t.Fatalf("Part1(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Sample(t *testing.T) {
	got := (Day05{}).Part2(strings.TrimSpace(sample))
	const want = "14"
	if got != want {
		t.Fatalf("Part2(sample) = %s, want %s", got, want)
	}
}
// 315984915553122 is too low
// 343143696885068 is too high
// 343143696885064 is too high