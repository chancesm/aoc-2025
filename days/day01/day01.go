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

func (Day01) Part2(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	dial := int64(50)
	count0 := 0
	for _, line := range lines {
		start := dial
		mod := parseLine(line)
		dial2, crossed0 := applyMod2(dial, mod)
		dial = dial2
		if dial == 0 || (crossed0 > 0 && start != 0) {
			count0 += 1
		}
		fmt.Printf("%s: %d, %d\n", line, dial, count0)
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

func adjust2(val int64) (int64, int) {
	callCount := 0
	if val > 99 {
		out, cc := adjust2(val - 100)
		return out, cc + callCount
	}
	if val < 0 {
		out, cc := adjust2(val + 100)
		return out, cc + callCount
	}
	return val, callCount
}

func applyMod(dial, mod int64) int64 {
	temp := dial + mod

	return adjust(temp)
}
func applyMod2(dial, mod int64) (int64, int) {
	temp := dial + mod
	adjusted, callCount := adjust2(temp)
	if adjusted == 0 || (temp != adjusted){
		callCount += 1
	}
	return adjusted, callCount

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
