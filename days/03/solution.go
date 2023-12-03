package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	"github.com/Tch1b0/polaris/runes"
	"github.com/Tch1b0/polaris/strings"
)

// shortcut for creating an int vector
func vec(x, y int) math.Vector2[int] {
	return math.Vector2[int]{X: x, Y: y}
}

type Part struct {
	Positions []math.Vector2[int]
	value     string
	IsSymbol  bool
	IsNumber  bool
}

func (p Part) Eq(o Part) bool {
	return array.All(p.Positions, func(pos math.Vector2[int], i int) bool {
		return pos == o.Positions[i]
	})
}

func (p Part) String() string {
	return fmt.Sprintf("{%s at %s}", p.value, p.Positions[0].String())
}

type Engine struct {
	Parts []Part
}

func (e Engine) atPos(pos math.Vector2[int]) *Part {
	for _, part := range e.Parts {
		for _, partPos := range part.Positions {
			if partPos == pos {
				return &part
			}
		}
	}

	return nil
}

func (e Engine) neighboursOf(p Part) []Part {
	neighbourPositions := []math.Vector2[int]{}
	for i, pos := range p.Positions {
		neighbourPositions = append(neighbourPositions, pos.Add(vec(0, 1)))
		neighbourPositions = append(neighbourPositions, pos.Sub(vec(0, 1)))

		if i == 0 {
			neighbourPositions = append(neighbourPositions, pos.Sub(vec(1, 0)))
			neighbourPositions = append(neighbourPositions, pos.Sub(vec(1, 1)))
			neighbourPositions = append(neighbourPositions, pos.Add(vec(-1, 1)))
		}

		if i == len(p.Positions)-1 {
			neighbourPositions = append(neighbourPositions, pos.Add(vec(1, 0)))
			neighbourPositions = append(neighbourPositions, pos.Add(vec(1, 1)))
			neighbourPositions = append(neighbourPositions, pos.Add(vec(1, -1)))
		}
	}

	neighbours := []Part{}
	for _, np := range neighbourPositions {
		n := e.atPos(np)
		// check if neighbour exists, and isn't already being tracked
		if n != nil && array.None(neighbours, func(v Part, _ int) bool { return v.Eq(*n) }) {
			neighbours = append(neighbours, *n)
		}
	}

	return neighbours
}

func isSymbol(r rune) bool {
	return !runes.IsDigit(r) && r != '.'
}

func getInput() Engine {
	return input.Process("./days/03/input.txt", func(str string) Engine {
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")
		e := Engine{Parts: []Part{}}

		for y, line := range lines {
			buf := ""
			bufPositions := []math.Vector2[int]{}

			parseBuf := func() {
				if buf != "" {
					p := Part{
						Positions: bufPositions,
						value:     buf,
						IsSymbol:  false,
						IsNumber:  true,
					}
					e.Parts = append(e.Parts, p)

					buf = ""
					bufPositions = []math.Vector2[int]{}
				}
			}

			for x, c := range line {
				if runes.IsDigit(c) {
					buf += string(c)
					bufPositions = append(bufPositions, vec(x, y))
				} else {
					parseBuf()

					p := Part{
						Positions: []math.Vector2[int]{{X: x, Y: y}},
						value:     string(c),
						IsSymbol:  isSymbol(c),
						IsNumber:  false,
					}
					e.Parts = append(e.Parts, p)
				}
			}
			parseBuf()
		}

		return e
	})
}

func hasSymbolNeighbours(e Engine, p Part) bool {
	for _, n := range e.neighboursOf(p) {
		if n.IsSymbol {
			return true
		}
	}

	return false
}

func part1(engine Engine) int {
	sum := 0

	for _, p := range engine.Parts {
		if p.IsNumber && hasSymbolNeighbours(engine, p) {
			v, err := strings.Atoi(p.value)
			if err != nil {
				panic(err)
			}
			sum += v
		}
	}

	return sum
}

func getNumberNeighbours(e Engine, p Part) []Part {
	nums := []Part{}

	for _, n := range e.neighboursOf(p) {
		if n.IsNumber {
			nums = append(nums, n)
		}
	}

	return nums
}

func part2(engine Engine) int {
	sum := 0

	for _, p := range engine.Parts {
		if p.value != "*" {
			continue
		}

		numNeighbours := getNumberNeighbours(engine, p)
		if len(numNeighbours) != 2 {
			continue
		}

		v1, err := strings.Atoi(numNeighbours[0].value)
		if err != nil {
			panic(err)
		}
		v2, err := strings.Atoi(numNeighbours[1].value)
		if err != nil {
			panic(err)
		}

		sum += v1 * v2
	}

	return sum
}

func main() {
	in := getInput()

	fmt.Printf("PART 1: %d\n", part1(in))

	fmt.Printf("PART 2: %d\n", part2(in))
}
