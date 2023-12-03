package main

import (
	"fmt"

	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	"github.com/Tch1b0/polaris/strings"
)

type Set struct {
	RedCubes   int
	GreenCubes int
	BlueCubes  int
}

func (s *Set) setColorCount(color string, count int) {
	switch color {
	case "red":
		s.RedCubes = count
	case "green":
		s.GreenCubes = count
	case "blue":
		s.BlueCubes = count
	default:
		panic(fmt.Sprintf("Unknown color \"%s\"", color))
	}
}

func (s Set) String() string {
	return fmt.Sprintf("{R: %d, G: %d, B: %d}", s.RedCubes, s.GreenCubes, s.BlueCubes)
}

type Game struct {
	Id   int
	Sets []Set
}

func getInput() []Game {
	return input.Process("./days/02/input.txt", func(str string) []Game {
		// remove the evil "\r", which was trying to make me go insane, and split lines
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")
		games := []Game{}

		for _, line := range lines {
			g := Game{Sets: []Set{}}

			// get the id by splitting the string at the ":" subtracting the "Game " prefix of it
			strId := stdstrings.Replace(stdstrings.Split(line, ":")[0], "Game ", "", 1)
			id, err := strings.Atoi(strId)
			if err != nil {
				panic(err)
			}
			g.Id = id

			// iterate over all sets in the game by taking the right side of the colon and seperating the sets by the semicolon
			for _, strSet := range stdstrings.Split(stdstrings.Split(line, ":")[1], ";") {
				set := Set{}

				for _, part := range stdstrings.Split(strSet, ",") {
					// omit the trailing whitespace " X COLOR" to "X COLOR"
					part = part[1:]

					// split number and color at space, so from "X COLOR" to ["X", "COLOR"]
					splittedPart := stdstrings.Split(part, " ")
					strCount, color := splittedPart[0], splittedPart[1]

					count, err := strings.Atoi(strCount)
					if err != nil {
						panic(err)
					}

					set.setColorCount(color, count)
				}

				g.Sets = append(g.Sets, set)
			}

			games = append(games, g)
		}

		return games
	})
}

func part1(games []Game) int {
	sum := 0

	for _, game := range games {
		possible := true

		for _, set := range game.Sets {

			// check whether cube numbers are in bound
			if !(set.RedCubes <= 12 && set.GreenCubes <= 13 && set.BlueCubes <= 14) {
				possible = false
				break
			}

		}

		if possible {
			sum += game.Id
		}
	}

	return sum
}

func part2(games []Game) int {
	power := 0

	for _, game := range games {
		// collections of all cube numbers in sets
		r, g, b := []int{}, []int{}, []int{}

		for _, set := range game.Sets {
			r = append(r, set.RedCubes)
			g = append(g, set.GreenCubes)
			b = append(b, set.BlueCubes)
		}

		power += math.Max(r) * math.Max(g) * math.Max(b)
	}

	return power
}

func main() {
	in := getInput()

	fmt.Printf("PART 1: %d\n", part1(in))

	fmt.Printf("PART 2: %d\n", part2(in))
}
