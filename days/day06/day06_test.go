package Day06

import (
	"strings"
	"testing"
)

const sample = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func TestPart1_Sample(t *testing.T) {
	got := (Day06{}).Part1(strings.TrimSpace(sample))
	const want = "4277556"
	if got != want {
		t.Fatalf("Part1(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Sample(t *testing.T) {
	got := (Day06{}).Part2(strings.TrimSpace(sample))
	const want = "3263827"
	if got != want {
		t.Fatalf("Part2(sample) = %s, want %s", got, want)
	}
}

