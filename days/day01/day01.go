package day01

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/chancesm/aoc-2025/internal/framework"
	"github.com/chancesm/aoc-2025/pkg/util"
)

//go:embed input.txt
var input string

type Day01 struct{}

func init() {
	framework.Register(1, Day01{})
}

func (Day01) Input() string {
	return input
}

func (Day01) Part1(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	dial := int64(50)
	count0 := 0
	for _, line := range lines {
		mod := parseLine(line)
		dial = applyMod(dial, mod)
		if dial == 0 {
			count0 += 1
		}
	}
	return strconv.Itoa(count0)
}

// I gave up and went with the "brute force" method
func (Day01) Part2(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	dial := 50
	count0 := 0
	for _, line := range lines {
		mod := int(parseLine(line))
		if mod > 0 {
			for i := 0; i < mod; i++ {
				dial = dial + 1
				if dial == 100 {
					dial = 0
					count0 += 1
				}
			}
		}
		if mod < 0 {
			for i := 0; i > mod; i-- {
				dial = dial - 1
				if dial == 0 {
					count0 += 1
				}
				if dial == -1 {
					dial = 99
				}
			}

		}
	}

	return strconv.Itoa(count0)
}

func adjust(val int64) int64 {
	if val > 99 {
		return adjust(val - 100)
	}
	if val < 0 {
		return adjust(val + 100)
	}
	return val
}

func applyMod(dial, mod int64) int64 {

	temp := dial + mod
	adjusted := adjust(temp)

	return adjusted
}

func parseLine(str string) int64 {
	dir := str[:1]
	var strMod string
	if dir == "L" {
		strMod = "-"
	} else {
		strMod = "+"
	}
	val := str[1:]
	mod, _ := strconv.ParseInt(fmt.Sprintf("%s%s", strMod, val), 10, 0)
	return mod
}
