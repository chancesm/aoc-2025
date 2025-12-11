package Day06

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"github.com/chancesm/aoc-2025/internal/framework"
	"github.com/chancesm/aoc-2025/pkg/util"
)

//go:embed input.txt
var input string

const RegExpPattern string = `[\*\+][ ]+`

type Day06 struct{}

func init() {
	framework.Register(6, Day06{})
}

func (Day06) Input() string {
	return input
}

type Problem struct {
	values  []int
	mod     string
	maxChar int
}

func (Day06) Part1(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	probValLen := len(lines) - 1
	_ = probValLen

	problems := make([]Problem, 0)
	problemsCreated := false
	for i, line := range lines {
		vals := strings.Fields(line)
		if i == len(lines)-1 {
			// set mods
			for j, mod := range vals {
				problems[j].mod = mod
			}
		} else {
			for j, val := range vals {
				if !problemsCreated {
					for k := 0; k < len(vals); k++ {
						problems = append(problems, Problem{
							values: make([]int, probValLen),
							mod:    "",
						})
					}
					problemsCreated = true
				}
				num_val, _ := strconv.Atoi(val)
				problems[j].values[i] = num_val
			}
		}
	}
	final := 0
	for _, p := range problems {
		var prob_ans int
		if p.mod == "*" {
			prob_ans = 1
			for _, n := range p.values {
				prob_ans *= n
			}
		} else {
			prob_ans = 0
			for _, n := range p.values {
				prob_ans += n
			}
		}
		final += prob_ans
	}
	return strconv.Itoa(final)
}

func (Day06) Part2(inputStr string) string {
	re := regexp.MustCompile(RegExpPattern)
	lines := util.NonEmptyLinesNoTrim(inputStr)
	probValLen := len(lines) - 1
	_ = probValLen
	lastLine := lines[len(lines)-1]
	patterns := re.FindAll([]byte(lastLine), -1)
	problems := make([]Problem, 0)

	for i_p, p := range patterns {
		value_len := len(p) - 1
		if i_p == len(patterns)-1 {
			value_len = len(p)
		}
		problems = append(problems, Problem{
			values:  make([]int, value_len),
			mod:     string(p[0]),
			maxChar: len(p) - 1,
		})
	}
	var currentProb Problem
	currentProbIdx := -1
	valueIdx := 0
	for c := 0; c < len(lines[0]); c++ {
		pow := 1
		for r := len(lines) - 1; r >= 0; r-- {
			val := string(lines[r][c])
			if val == "*" || val == "+" {
				// operate on next problem
				valueIdx = 0
				currentProbIdx += 1
				currentProb = problems[currentProbIdx]
				continue
			}
			if val != " " {
				// parse, mult by pow, and add to problem's values at value index
				v, _ := strconv.Atoi(val)
				currentProb.values[valueIdx] += v * pow
				// mult pow by 10
				pow *= 10
			}
		}
		valueIdx++
	}
	final := 0
	for _, p := range problems {
		var prob_ans int
		if p.mod == "*" {
			prob_ans = 1
			for _, n := range p.values {
				prob_ans *= n
			}
		} else {
			prob_ans = 0
			for _, n := range p.values {
				prob_ans += n
			}
		}
		final += prob_ans
	}
	return strconv.Itoa(final)
}

//Note: we are looping char int values, not the string pieces. Need to fix this
