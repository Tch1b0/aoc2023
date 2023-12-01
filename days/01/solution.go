package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/strings"
)

var digit = map[string]int{
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
		if c >= '0' && c <= '9' {
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
	for key, value := range digit {
		if stdstrings.Contains(str, key) {
			return &value
		}
	}

	return nil
}

func getResultNum(str string) int {
	var buf string
	digits := ""

	// first number
	buf = ""
	for _, c := range str {
		buf += string(c)
		if c >= '0' && c <= '9' {
			digits += string(c)
			break
		} else if d := containsWordDigit(buf); d != nil {
			digits += strings.Itoa(*d)
			break
		}
	}

	// last number
	buf = ""
	for i := len(str) - 1; i >= 0; i-- {
		c := str[i]
		buf = string(c) + buf
		if c >= '0' && c <= '9' {
			digits += string(c)
			break
		} else if d := containsWordDigit(buf); d != nil {
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
