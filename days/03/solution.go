package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/runes"
	"github.com/Tch1b0/polaris/strings"
)

var (
	colorReset = "\033[0m"

	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type Engine [][]rune

func getInput() Engine {
	return input.Process("./days/03/input.txt", func(str string) Engine {
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")
		e := Engine{}

		for _, line := range lines {
			e = append(e, strings.ToRunes(line))
		}

		return e
	})
}

func isSymbol(r rune) bool {
	return !runes.IsDigit(r) && r != '.'
}

func hasSymbolNeighbours(e Engine, x, y int) bool {
	xIsBottom := x == 0
	xIsTop := x == len(e[y])-1

	yIsBottom := y == 0
	yIsTop := y == len(e)-1

	return (!xIsBottom && isSymbol(e[y][x-1])) || // check left
		(!xIsTop && isSymbol(e[y][x+1])) || // check right
		(!yIsBottom && isSymbol(e[y-1][x])) || // check top
		(!yIsTop && isSymbol(e[y+1][x])) || // check bottom
		(!yIsBottom && !xIsBottom && isSymbol(e[y-1][x-1])) || // check top left
		(!yIsBottom && !xIsTop && isSymbol(e[y-1][x+1])) || // check top right
		(!yIsTop && !xIsTop && isSymbol(e[y+1][x+1])) || // check bottom right
		(!yIsTop && !xIsBottom && isSymbol(e[y+1][x-1])) // check bottom left
}

func part1(engine Engine, visual bool) int {
	sum := 0

	for y, line := range engine {
		buf := ""
		hasSymbol := false

		for x, c := range line {
			if runes.IsDigit(c) {

				if hasSymbolNeighbours(engine, x, y) {
					hasSymbol = true
				}

				buf += string(c)

			} else {

				if len(buf) != 0 && hasSymbol {
					v, err := strings.Atoi(buf)
					if err != nil {
						panic(err)
					}

					if visual {
						fmt.Print(colorGreen + buf + colorReset)
					}

					sum += v
				} else if visual && len(buf) > 0 && !hasSymbol {
					fmt.Print(colorRed + buf + colorReset)
				}

				if visual {
					fmt.Print(string(c))
				}

				buf = ""
				hasSymbol = false
			}

		}
		if len(buf) != 0 && hasSymbol {
			v, err := strings.Atoi(buf)
			if err != nil {
				panic(err)
			}
			sum += v
		}

		if visual {
			fmt.Println()
		}
	}

	return sum
}

func part2() int {
	return -1
}

func main() {
	in := getInput()

	fmt.Printf("PART 1: %d\n", part1(in, false))

	fmt.Printf("PART 2: %d\n", part2(in))
}
