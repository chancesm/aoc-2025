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

func TestPart2_Small1(t *testing.T) {
	got := (Day03{}).Part2("987654321111111")
	const want = "987654321111"
	if got != want {
		t.Fatalf("Part2(small1) = %s, want %s", got, want)
	}
}
func TestPart2_Small2(t *testing.T) {
	got := (Day03{}).Part2("811111111111119")
	const want = "811111111119"
	if got != want {
		t.Fatalf("Part2(small2) = %s, want %s", got, want)
	}
}
func TestPart2_Small3(t *testing.T) {
	got := (Day03{}).Part2("234234234234278")
	const want = "434234234278"
	if got != want {
		t.Fatalf("Part2(small3) = %s, want %s", got, want)
	}
}
func TestPart2_Small4(t *testing.T) {
	got := (Day03{}).Part2("818181911112111")
	const want = "888911112111"
	if got != want {
		t.Fatalf("Part2(small4) = %s, want %s", got, want)
	}
}
