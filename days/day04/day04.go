package day04

import (
	_ "embed"
	"strconv"

	"github.com/chancesm/aoc-2025/internal/framework"
	"github.com/chancesm/aoc-2025/pkg/util"
)

//go:embed input.txt
var input string

type Day04 struct{}

func init() {
	framework.Register(4, Day04{})
}

func (Day04) Input() string {
	return input
}

func (Day04) Part1(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	final := 0

	grid := make([][]int, 0)

	for l_i, line := range lines {
		grid = append(grid, make([]int, len(line)))
		for i, r := range line {
			if r == '@' {
				grid[l_i][i] = 1
			}
		}
	}

	//
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 && isAccessible(x, y, grid) {
				grid[x][y] = 2
				final += 1
			}
		}
	}

	return strconv.Itoa(final)
}

func (Day04) Part2(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	final := 0

	grid := make([][]int, 0)
	// make grid from input
	for l_i, line := range lines {
		grid = append(grid, make([]int, len(line)))
		for i, r := range line {
			if r == '@' {
				grid[l_i][i] = 1
			}
		}
	}
	keepGoing := true
	for keepGoing {
		// parseGrid
		count, kg := parseGrid(grid)
		keepGoing = kg
		if kg {
			final += count
			updateGrid(grid)
		}
	}
	return strconv.Itoa(final)
}

func isAccessible(x, y int, grid [][]int) bool {
	dirs := []int{-1, 0, 1}
	count := 0
	for _, x_mod := range dirs {
		for _, y_mod := range dirs {
			r := x + x_mod
			c := y + y_mod
			if (x_mod == 0 && y_mod == 0) || r == -1 || c == -1 || r == len(grid) || c == len(grid[x]) {
				continue
			}
			if grid[r][c] == 1 || grid[r][c] == 2 {
				count += 1
			}
		}
	}
	return count < 4
}

func parseGrid(grid [][]int) (int, bool) {
	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 && isAccessible(x, y, grid) {
				grid[x][y] = 2
				count += 1
			}
		}
	}
	if count == 0 {
		return count, false
	}
	return count, true
}

func updateGrid(grid [][]int) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 2 {
				grid[x][y] = 0
			}
		}
	}
}
