package day01

import (
	"strings"
	"testing"
)

const sample = `
11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124
`

func TestPart1_Sample(t *testing.T) {
	got := (Day02{}).Part1(strings.TrimSpace(sample))
	const want = "1227775554"
	if got != want {
		t.Fatalf("Part1(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Sample(t *testing.T) {
	got := (Day02{}).Part2(strings.TrimSpace(sample))
	const want = "4174379265"
	if got != want {
		t.Fatalf("Part2(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Full(t *testing.T) {
	day2 := Day02{}
	got := day2.Part2(day2.Input())
	want := "65794984339"
	if got != want {
		t.Fatalf("Part2(full) = %s, want %s", got, want)
	}
}
