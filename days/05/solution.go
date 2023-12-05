package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	"github.com/Tch1b0/polaris/strings"
)

type MapRange struct {
	Src  math.Span[int]
	Dest math.Span[int]
}

func (mr MapRange) MapValue(v int) int {
	if !mr.Src.Contains(v) {
		panic(fmt.Sprintf("Value %d not included in %v", v, mr))
	}

	return v - mr.Src.From + mr.Dest.From
}

type AtoBMap struct {
	A        string
	B        string
	AtoBVals []MapRange
	memo     map[int]int
}

func (ab *AtoBMap) MapValue(v int) int {
	// if val, ok := ab.memo[v]; ok {
	// 	return val
	// }

	x := array.GetFirst(ab.AtoBVals, func(m MapRange, _ int) bool {
		return m.Src.Contains(v)
	})

	if x != nil {
		val := x.MapValue(v)
		// ab.memo[v] = val
		return val
	} else {
		// ab.memo[v] = v
		return v
	}
}

type Production struct {
	Seeds    []int
	AtoBMaps []AtoBMap
}

func (p Production) ProcessSeed(v int) int {
	for _, m := range p.AtoBMaps {
		v = m.MapValue(v)
	}
	return v
}

func spaceSeperatedToInt(str string) []int {
	return array.Map(
		stdstrings.Split(str, " "),
		func(v string, _ int) int {
			n, err := strings.Atoi(v)
			if err != nil {
				panic(err)
			}
			return n
		},
	)
}

func getInput() Production {
	return input.Process("./days/05/input.txt", func(str string) Production {
		p := Production{AtoBMaps: []AtoBMap{}}

		blocks := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n\n")
		p.Seeds = spaceSeperatedToInt(stdstrings.ReplaceAll(blocks[0], "seeds: ", ""))

		blocks = blocks[1:]

		for _, block := range blocks {
			m := AtoBMap{memo: make(map[int]int)}

			lines := stdstrings.Split(block, "\n")
			l := stdstrings.Split(stdstrings.ReplaceAll(lines[0], " map:", ""), "-to-")
			m.A = l[0]
			m.B = l[1]

			lines = lines[1:]
			for _, line := range lines {
				nums := spaceSeperatedToInt(line)
				a, b, c := nums[0], nums[1], nums[2]

				vals := MapRange{
					Src: math.Span[int]{
						From: b,
						To:   b + c,
					},
					Dest: math.Span[int]{
						From: a,
						To:   a + c,
					},
				}

				m.AtoBVals = append(m.AtoBVals, vals)
			}

			p.AtoBMaps = append(p.AtoBMaps, m)

		}

		return p
	})
}

func part1(p Production) int {
	soils := []int{}
	for _, seed := range p.Seeds {
		soils = append(soils, p.ProcessSeed(seed))
	}

	return math.Min(soils)
}

func part2(p Production) int {
	fmt.Println("Part 2 will take about 30 minutes")

	k := -1
	for i := 0; i < len(p.Seeds)-1; i += 2 {
		percentageDone := (float64(i) / float64(len(p.Seeds))) * 100
		fmt.Printf("\r%d%% ", int(percentageDone))

		start := p.Seeds[i]
		length := p.Seeds[i+1]

		spn := math.Span[int]{From: start, To: start + length}

		for seed := range spn.Iter() {
			soil := p.ProcessSeed(seed)
			if soil < k || k == -1 {
				k = soil
			}
		}
	}

	fmt.Println()

	return k
}

func main() {
	in := getInput()

	fmt.Printf("PART 1: %d\n", part1(in))

	fmt.Printf("PART 2: %d\n", part2(in))
}
