package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/runes"
	"github.com/Tch1b0/polaris/strings"
)

var digitWord = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getInput() []string {
	return input.Process("./days/01/input.txt", func(str string) []string {
		return stdstrings.Split(str, "\n")
	})
}

func getNumbers(str string) []int {
	nums := []int{}
	for _, c := range str {
		if runes.IsDigit(c) {
			nums = append(nums, int(c-'0'))
		}
	}

	return nums
}

func part1(in []string) int {
	sum := 0
	for _, line := range in {
		nums := getNumbers(line)
		v, _ := strings.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
		sum += v
	}

	return sum
}

func containsWordDigit(str string) *int {
	for key, value := range digitWord {
		if stdstrings.Contains(str, key) {
			return &value
		}
	}

	return nil
}

func getResultNum(str string) int {
	var buf string
	digits := ""

	// get first number

	// reset buffer
	buf = ""

	// iterate over string forwards
	for _, c := range str {
		// append new character to buffer
		buf += string(c)

		// check if character is a literal digit
		if runes.IsDigit(c) {
			digits += string(c)
			break
		} else /* check for word digit */ if d := containsWordDigit(buf); d != nil {
			digits += strings.Itoa(*d)
			break
		}
	}

	// get last number

	// reset buffer
	buf = ""

	// iterate over string backwards
	for i := len(str) - 1; i >= 0; i-- {
		// get current character
		c := rune(str[i])

		// append character at front of buffer, so it comes out as forward text
		buf = string(c) + buf

		// check if character is a literal digit
		if runes.IsDigit(c) {
			digits += string(c)
			break
		} else /* check for word digit */ if d := containsWordDigit(buf); d != nil {
			digits += strings.Itoa(*d)
			break
		}
	}

	res, _ := strings.Atoi(digits)
	return res
}

func part2(in []string) int {
	sum := 0
	for _, line := range in {
		v := getResultNum(line)
		sum += v
	}
	return sum
}

func main() {
	in := getInput()
	fmt.Printf("PART 1: %d\n", part1(in))

	fmt.Printf("PART 2: %d\n", part2(in))
}
