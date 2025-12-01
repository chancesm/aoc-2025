package framework

import (
	"fmt"
)

type Runner interface {
	Part1(input string) string
	Part2(input string) string
}

// Implemented by days that embed their real input.
type InputProvider interface {
	Input() string
}

var registry = map[int]Runner{}

func Register(day int, r Runner) {
	if _, exists := registry[day]; exists {
		panic(fmt.Sprintf("day %d already registered", day))
	}
	registry[day] = r
}

func Get(day int) (Runner, bool) {
	r, ok := registry[day]
	return r, ok
}
