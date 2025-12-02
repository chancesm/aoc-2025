package day01

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/chancesm/aoc-2025/internal/framework"
	"github.com/chancesm/aoc-2025/pkg/util"
)

//go:embed input.txt
var input string

type Day02 struct{}

func init() {
	framework.Register(2, Day02{})
}

func (Day02) Input() string {
	return input
}

func (Day02) Part1(inputStr string) string {
	invalidIDs := make([]int, 0)
	ranges := util.CommaSeparated(inputStr)
	for _, rng := range ranges {
		start, end := splitRange(rng)
		for i := start; i <= end; i++ {
			str_val := strconv.Itoa(i)
			str_len := len(str_val)
			if str_len%2 != 0 {
				continue
			}
			middle := str_len / 2
			if str_val[:middle] == str_val[middle:] {
				invalidIDs = append(invalidIDs, i)
			}

		}
	}
	count := 0
	for _, v := range invalidIDs {
		count += v
	}
	return strconv.Itoa(count)
}

func (Day02) Part2(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	return strconv.Itoa(len(lines))
}

func splitRange(s string) (int, int) {
	vals := strings.Split(s, "-")
	start, _ := strconv.ParseInt(vals[0], 10, 0)
	end, _ := strconv.ParseInt(vals[1], 10, 0)
	return int(start), int(end)
}
