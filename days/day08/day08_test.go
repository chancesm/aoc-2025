package Day08

import (
	"strings"
	"testing"
)

const sample = `
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

func TestPart1_Sample(t *testing.T) {
	got := (Day08{}).Part1(strings.TrimSpace(sample))
	const want = "?????"
	if got != want {
		t.Fatalf("Part1(sample) = %s, want %s", got, want)
	}
}

func TestPart2_Sample(t *testing.T) {
	got := (Day08{}).Part2(sample)
	const want = "?????"
	if got != want {
		t.Fatalf("Part2(sample) = %s, want %s", got, want)
	}
}
