package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	"github.com/Tch1b0/polaris/strings"
)

const (
	T_FIOAK = iota
	T_FOOAK
	T_FH
	T_THOAK
	T_TP
	T_OP
	T_HC
)

var cardTypes = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var cardTypes2 = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

type Hand struct {
	Cards []rune
	Bid   int
}

func (h Hand) GetType(part1 bool) int {
	occurences := map[rune]int{}
	for _, c := range h.Cards {
		v := occurences[c]
		occurences[c] = v + 1
	}

	ocCount := []int{}
	for key, v := range occurences {
		if (!part1 && key != 'J') || part1 {
			ocCount = append(ocCount, v)
		}
	}
	ocCount = array.Reverse(math.Sort(ocCount))

	v := 0
	if part1 {
		v = ocCount[0]
	} else {
		if len(ocCount) > 0 {
			v = ocCount[0]
		}
		v += occurences['J']
	}

	switch v {
	case 5:
		return T_FIOAK
	case 4:
		return T_FOOAK
	case 3:
		if len(ocCount) > 1 && ocCount[1] == 2 {
			return T_FH
		} else {
			return T_THOAK
		}
	case 2:
		if len(ocCount) > 1 && ocCount[1] == 2 {
			return T_TP
		} else {
			return T_OP
		}
	default:
		return T_HC
	}
}

func (h Hand) IsBetterThan(o Hand, part1 bool) bool {
	hType, oType := h.GetType(part1), o.GetType(part1)
	if hType != oType {
		return hType < oType
	}

	var ct []rune
	if part1 {
		ct = cardTypes
	} else {
		ct = cardTypes2
	}

	for i, c1 := range h.Cards {
		c2 := o.Cards[i]
		c1Value, c2Value := array.Index(ct, c1), array.Index(ct, c2)
		if c1Value != c2Value {
			return c1Value < c2Value
		}
	}
	return false
}

func getInput() []Hand {
	return input.Process("./days/07/input.txt", func(str string) []Hand {
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")
		hands := []Hand{}
		for _, line := range lines {
			h := Hand{}
			x := stdstrings.Split(line, " ")
			h.Cards = strings.ToRunes(x[0])

			v, err := strings.Atoi(x[1])
			if err != nil {
				panic(err)
			}
			h.Bid = v

			hands = append(hands, h)
		}

		return hands
	})
}

func part1(hands []Hand) int {
	for i, value := range hands {
		for j := 0; j < i; j++ {
			if hands[j].IsBetterThan(value, true) {
				hands = array.Move(hands, i, j)
				break
			}
		}
	}

	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.Bid
	}

	return sum
}

func part2(hands []Hand) int {
	for i, value := range hands {
		for j := 0; j < i; j++ {
			if hands[j].IsBetterThan(value, false) {
				hands = array.Move(hands, i, j)
				break
			}
		}
	}

	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.Bid
	}

	return sum
}

func main() {
	hands := getInput()

	fmt.Printf("PART 1: %d\n", part1(hands))

	fmt.Printf("PART 2: %d\n", part2(hands))
}
