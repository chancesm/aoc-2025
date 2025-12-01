package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/chancesm/aoc-2025/days/day01"
	"github.com/chancesm/aoc-2025/internal/framework"
)

func main() {
	day := flag.Int("day", 1, "day number (1-25)")
	part := flag.Int("part", 1, "part number (1 or 2)")
	flag.Parse()
	fmt.Printf("Running Day: %d Part: %d\n", *day, *part)
	runner, ok := framework.Get(*day)
	if !ok {
		log.Fatalf("no runner registered for day %d", *day)
	}

	ip, ok := runner.(framework.InputProvider)
	if !ok {
		log.Fatalf("runner for day %d does not provide real input", *day)
	}

	input := ip.Input()

	var ans string
	switch *part {
	case 1:
		ans = runner.Part1(input)
	case 2:
		ans = runner.Part2(input)
	default:
		log.Fatalf("invalid part: %d", *part)
	}

	fmt.Println(ans)
}
