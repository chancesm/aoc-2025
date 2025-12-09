package day04

import (
	"strings"
	"testing"
)

const sample = `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestPart1_Sample(t *testing.T) {
	got := (Day04{}).Part1(strings.TrimSpace(sample))
	const want = "13"
	if got != want {
		t.Fatalf("Part1(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Sample(t *testing.T) {
	got := (Day04{}).Part2(strings.TrimSpace(sample))
	const want = "43"
	if got != want {
		t.Fatalf("Part2(sample) = %s, want %s", got, want)
	}
}
