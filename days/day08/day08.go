package Day08

import (
	_ "embed"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/chancesm/aoc-2025/internal/framework"
	"github.com/chancesm/aoc-2025/pkg/util"
)

//go:embed input.txt
var input string

const RegExpPattern string = `[\*\+][ ]+`

type Day08 struct{}

func init() {
	framework.Register(8, Day08{})
}

func (Day08) Input() string {
	return input
}

type Problem struct {
	values  []int
	mod     string
	maxChar int
}

type Box struct {
	x int
	y int
	z int
}

type Connection struct {
	a        Box
	b        Box
	distance float64
}

func (Day08) Part1(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)
	boxes := make([]Box, len(lines))
	for i_l, line := range lines {
		x, y, z := parseCoord(line)
		boxes[i_l] = Box{
			x: x,
			y: y,
			z: z,
		}
	}
	connections := make([]Connection, 0)
	for i_b1 := 0; i_b1 < len(boxes)-1; i_b1++ {
		for i_b2 := i_b1 + 1; i_b2 < len(boxes); i_b2++ {
			connections = append(connections, Connection{
				a:        boxes[i_b1],
				b:        boxes[i_b2],
				distance: calcDistance(boxes[i_b1], boxes[i_b2]),
			})
		}
	}
	slices.SortFunc(connections, func(a, b Connection) int {
		if a.distance > b.distance {
			return 1
		}
		if a.distance < b.distance {
			return -1
		}
		return 0
	})
	for i_conn := 0; i_conn < 10; i_conn++ {
		// build circuits...yay
	}
	return strconv.Itoa(len(lines))
}

func (Day08) Part2(inputStr string) string {
	lines := util.NonEmptyLines(inputStr)

	return strconv.Itoa(len(lines))
}

func parseCoord(line string) (int, int, int) {
	parts := strings.Split(line, ",")
	coords := make([]int, 3)
	for i_p, part := range parts {
		v, _ := strconv.Atoi(part)
		coords[i_p] = v
	}
	x := coords[0]
	y := coords[1]
	z := coords[2]
	return x, y, z
}

func calcDistance(a, b Box) float64 {
	return math.Sqrt(sqr(a.x-b.x) + sqr(a.y-b.y) + sqr(a.z-b.z))
}

func sqr(v int) float64 {
	return math.Pow(float64(v), 2)
}
