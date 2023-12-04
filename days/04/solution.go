package main

import (
	"fmt"
	"math"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/strings"
)

type Card struct {
	Id   int
	Got  []int
	Want []int
}

func (c Card) CountCorrect() int {
	count := 0
	for _, n := range c.Got {
		if array.GetFirst(c.Want, func(v int, _ int) bool { return v == n }) != nil {
			count++
		}
	}
	return count
}

func spacedNumsToArr(spacedNums string) []int {
	nums := []int{}

	for _, n := range stdstrings.Split(spacedNums, " ") {
		v, err := strings.Atoi(n)
		if err != nil {
			continue
		}
		nums = append(nums, v)
	}

	return nums
}

func getInput() []Card {
	return input.Process("./days/04/input.txt", func(str string) []Card {
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")
		cards := []Card{}

		for i, line := range lines {
			x := stdstrings.Split(line, ":")[1]
			spacedNums := stdstrings.Split(x, "|")
			left := spacedNumsToArr(spacedNums[0])
			right := spacedNumsToArr(spacedNums[1])

			cards = append(cards, Card{Id: i + 1, Got: right, Want: left})
		}

		return cards
	})
}

func part1(cards []Card) int {
	sum := 0

	for _, card := range cards {
		exp := card.CountCorrect() - 1

		if exp > -1 {
			sum += int(math.Pow(2, float64(exp)))
		}
	}

	return sum
}

func part2(cards []Card) int {
	sum := 0

	for i, card := range cards {
		count := card.CountCorrect()

		// iterate backwards, because forwards loop would be more complex
		for j := count; j > 0; j-- {
			// add the count of the copies and its copies
			sum += part2(cards[i+j : i+j+1])
		}

		// count this card
		sum++
	}

	return sum
}

func main() {
	in := getInput()

	fmt.Printf("PART 1: %d\n", part1(in))

	fmt.Printf("PART 2: %d\n", part2(in))
}
