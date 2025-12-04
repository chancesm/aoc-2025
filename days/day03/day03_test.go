package day03

import (
	"strings"
	"testing"
)

const sample = `
987654321111111
811111111111119
234234234234278
818181911112111
`

func TestPart1_Sample(t *testing.T) {
	got := (Day03{}).Part1(strings.TrimSpace(sample))
	const want = "357"
	if got != want {
		t.Fatalf("Part1(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Sample(t *testing.T) {
	got := (Day03{}).Part2(strings.TrimSpace(sample))
	const want = "3121910778619"
	if got != want {
		t.Fatalf("Part2(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Small(t *testing.T) {
	got := (Day03{}).Part2("818181911112111")
	const want = "888911112111"
	if got != want {
		t.Fatalf("Part2(small) = %s, want %s", got, want)
	}
}
