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

func TestAdjust(t *testing.T) {
	got := adjust(-1)
	want := int64(99)
	if got != want {
		t.Fatalf("adjust(-1) = %d, want %d", got, want)
	}

}

func TestApplyMod(t *testing.T) {
	got1 := applyMod(1, -120)
	want1 := int64(81)
	if got1 != want1 {
		t.Fatalf("applyMod(1,-120) = %d, want %d", got1, want1)
	}
	
}

func TestPart1_Sample(t *testing.T) {
	got := (Day01{}).Part1(strings.TrimSpace(sample))
	const want = "3"
	if got != want {
		t.Fatalf("Part1(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Sample(t *testing.T) {
	got := (Day01{}).Part2(strings.TrimSpace(sample))
	const want = "6"
	if got != want {
		t.Fatalf("Part2(sample) = %s, want %s", got, want)
	}
}
