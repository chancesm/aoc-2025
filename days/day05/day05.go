package day05

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/chancesm/aoc-2025/internal/framework"
	"github.com/chancesm/aoc-2025/pkg/util"
)

//go:embed input.txt
var input string

type Day05 struct{}

func init() {
	framework.Register(5, Day05{})
}

func (Day05) Input() string {
	return input
}

type Range struct {
	first int
	last  int
}

func (r *Range) InRange(x int) bool {
	return x >= r.first && x <= r.last
}
func sortRanges(a Range, b Range) int {
	if a.first < b.first {
		return -1
	} else if a.first > b.first {
		return 1
	}
	return 0
}

func (Day05) Part1(inputStr string) string {
	str_ranges, str_ids := splitInput(inputStr)
	ranges_split := util.NonEmptyLines(str_ranges)
	ids := util.NonEmptyLines(str_ids)
	ranges := make([]Range, len(ranges_split))
	for i, r := range ranges_split {
		vals := strings.Split(r, "-")
		f, _ := strconv.Atoi(vals[0])
		l, _ := strconv.Atoi(vals[1])
		ranges[i] = Range{first: f, last: l}
	}
	fresh := 0
idLoop:
	for _, id := range ids {
		for _, r := range ranges {
			vid, _ := strconv.Atoi(id)
			if r.InRange(vid) {
				fresh += 1
				continue idLoop
			}
		}
	}
	_ = ranges
	return strconv.Itoa(fresh)
}

func (Day05) Part2(inputStr string) string {
	str_ranges, _ := splitInput(inputStr)
	ranges_split := util.NonEmptyLines(str_ranges)

	// make ranges
	ranges := make([]Range, len(ranges_split))
	for i, r := range ranges_split {
		vals := strings.Split(r, "-")
		f, _ := strconv.Atoi(vals[0])
		l, _ := strconv.Atoi(vals[1])
		ranges[i] = Range{first: f, last: l}
	}
	slices.SortFunc(ranges, sortRanges)
	lastObserved := 0
	count := 0
	for _, r := range ranges {
		if lastObserved >= r.last {
			continue
		}
		if r.first <= lastObserved {
			r.first = lastObserved + 1
		}
		count += (r.last - r.first + 1)
		lastObserved = r.last
	}
	return strconv.Itoa(count)
}

func splitInput(input string) (ranges, ids string) {
	temp := strings.Split(strings.TrimSpace(input), "\n\n")
	return temp[0], temp[1]
}
