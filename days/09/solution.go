package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/strings"
)

type NumCol struct {
	Nums         []int
	Derivs       *NumCol
	PredictedVal *int
}

func (n *NumCol) fillDerivs() {
	derivs := NumCol{}
	derivs.Nums = make([]int, len(n.Nums)-1)

	for i := 0; i < len(n.Nums)-1; i++ {
		derivs.Nums[i] = n.Nums[i+1] - n.Nums[i]
	}

	n.Derivs = &derivs

	if !derivs.allZeros() {
		n.Derivs.fillDerivs()
	}
}

func (n NumCol) allZeros() bool {
	return array.All(n.Nums, func(num, _ int) bool { return num == 0 })
}

func (n *NumCol) PredictedNextValue() int {
	if n.Derivs == nil {
		n.fillDerivs()
	}

	if n.allZeros() {
		return 0
	}

	return n.Nums[len(n.Nums)-1] + n.Derivs.PredictedNextValue()
}

func (n *NumCol) PredictedPreviousValue() int {
	if n.Derivs == nil {
		n.fillDerivs()
	}

	if n.allZeros() {
		return 0
	}

	return n.Nums[0] - n.Derivs.PredictedPreviousValue()
}

func spacedNumsToArr(str string) []int {
	nums := []int{}
	for _, v := range stdstrings.Split(str, " ") {
		i, err := strings.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	return nums
}

func getInput() []NumCol {
	return input.Process("./days/09/input.txt", func(str string) []NumCol {
		cols := []NumCol{}
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")

		for _, line := range lines {
			arr := spacedNumsToArr(line)
			cols = append(cols, NumCol{Nums: arr})
		}

		return cols
	})
}

func part1(numCols []NumCol) int {
	sum := 0

	for _, col := range numCols {
		sum += col.PredictedNextValue()
	}

	return sum
}

func part2(numCols []NumCol) int {
	sum := 0

	for _, col := range numCols {
		sum += col.PredictedPreviousValue()
	}

	return sum
}

func main() {
	in := getInput()

	fmt.Printf("PART 1: %d\n", part1(in))

	fmt.Printf("PART 2: %d\n", part2(in))
}
