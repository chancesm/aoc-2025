package day01

import (
	"strings"
	"testing"
)

const sample = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestPart1_Sample(t *testing.T) {
	got := (Day01{}).Part1(strings.TrimSpace(sample))
	const want = "3"
	if got != want {
		t.Fatalf("Part1(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Sample(t *testing.T) {
	got := (Day01{}).Part2(strings.TrimSpace(sample))
	const want = "unimplemented"
	if got != want {
		t.Fatalf("Part2(sample) = %s, want %s", got, want)
	}
}
