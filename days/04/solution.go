package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/array"
)

type Card struct {
	Got []int
	Want []int
}

func spacedNumsToArr(spacedNums string) []int {
	nums := []int{}

	for _, n := range stdstrings.Split(spacedNums, " ") {
		v, err := strings.Atoi(n)
		if err != nil { panic(err) }
		nums = append(nums, v)
	}

	return nums
}

func getInput() []Card {
	return input.Process("./days/04/input.txt", func(str string) []Card {
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")
		cards := []Card{}

		for _, line := lines {
			x := stdstrings.Split(line, ":")[1]
			spacedNums := stdstrings.Split(x, "|")
			left := spacedNumsToArr(spacedNums[0])
			right := spacedNumsToArr(spacedNums[1])
		
			cards = append(cards, Card{Got: right, Want: left})
		}

		return cards
	}
}

func part1(cards []Card) int {
	sum := 0

	for _, card := range cards {
		exp := -1
		for _, n := range card.Got {
			if array.Contains(card.Want, n) {
				exp++
			}
		}

		if exp > -1 {
			sum += math.Pow(2, exp)
		}
	}

	return sum
}

func part2() int {
	return -1
}

func main() {
	in := getInput()

	fmt.Printf("PART 1: %d\n", part1(in))

	fmt.Printf("PART 2: %d\n", part2())
}
