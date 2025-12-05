package day03

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/chancesm/aoc-2025/internal/framework"
	"github.com/chancesm/aoc-2025/pkg/util"
)

//go:embed input.txt
var input string

type Day03 struct{}

func init() {
	framework.Register(3, Day03{})
}

func (Day03) Input() string {
	return input
}

func (Day03) Part1(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	final := 0
	for _, line := range lines {
		var first rune
		var second rune
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			current := runes[i]
			if current > first && i <= len(runes)-2 {
				first = current
				second = runes[i+1]
				continue
			}
			if current > second {
				second = current
			}
		}
		voltage, _ := strconv.Atoi(fmt.Sprintf("%c%c", first, second))
		final += voltage
	}

	return strconv.Itoa(final)
}

func (Day03) Part2(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	final := 0
	for _, line := range lines {
		iters := make([]rune, 12)
		runes := []rune(line)

		for i := 0; i < len(runes); i++ {
			
		voltage, _ := strconv.Atoi(string(iters))
		final += voltage
	}

	return strconv.Itoa(final)
}
