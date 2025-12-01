package day01

import (
	_ "embed"
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
	return strconv.Itoa(len(lines))
}

func (Day01) Part2(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	return strconv.Itoa(len(lines))
}
