package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
)

type Node struct {
	Code  string
	Left  string
	Right string
}

func (n Node) String() string {
	return fmt.Sprintf("%s = (%s, %s)", n.Code, n.Left, n.Right)
}

type Input struct {
	Instructions string
	Nodes        []Node
	memo         map[string]Node
}

func (i Input) GetStartNodes() []Node {
	n := []Node{}

	for _, node := range i.Nodes {
		if node.Code[2] == 'A' {
			n = append(n, node)
		}
	}

	return n
}

func (i Input) String(selected Node) string {
	s := ""
	for _, n := range i.Nodes {
		if selected.Code == n.Code {
			s += "\033[92m" + n.String() + "\033[0m\n"
		} else {
			s += n.String() + "\n"
		}
	}

	return s
}

func (i *Input) getNode(code string) *Node {
	return array.GetFirst(i.Nodes, func(n Node, i int) bool {
		return n.Code == code
	})
}

func getInput() Input {
	return input.Process("./days/08/input.txt", func(str string) Input {
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")
		in := Input{Instructions: lines[0], Nodes: []Node{}, memo: map[string]Node{}}

		for _, line := range lines[2:] {
			n := Node{
				Code:  line[0:3],
				Left:  line[7:10],
				Right: line[12:15],
			}

			in.Nodes = append(in.Nodes, n)
		}

		return in
	})
}

func part1(in Input) int {
	node := in.getNode("AAA")
	if node == nil {
		return -1
	}

	i := 0
	for node.Code != "ZZZ" {
		left := in.Instructions[i%len(in.Instructions)] == 'L'

		if left {
			node = in.getNode(node.Left)
		} else {
			node = in.getNode(node.Right)
		}

		i++

	}

	return i
}

func allNodesZ(nodes []Node) bool {
	return array.All(nodes, func(n Node, _ int) bool {
		return n.Code[2] == 'Z'
	})
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcd(k []int) int {
	l := func(a, b int) int { return (a * b) / gcd(a, b) }

	x := l(k[0], k[1])
	for i, v := range k {
		if i < 2 {
			continue
		}

		x = l(x, v)
	}

	return x
}

func part2(in Input) int {
	i := 0
	nodes := in.GetStartNodes()

	p := map[Node]int{}

	x := true

	for x {
		x = !allNodesZ(nodes)
		left := in.Instructions[i%len(in.Instructions)] == 'L'

		for j, node := range nodes {
			if node.Code[2] == 'Z' {
				if _, ok := p[node]; !ok {
					p[node] = i
				}
				continue
			}

			if left {
				nodes[j] = *in.getNode(node.Left)
			} else {
				nodes[j] = *in.getNode(node.Right)
			}
		}

		i++
	}

	k := []int{}
	for _, v := range p {
		k = append(k, v)
	}

	return lcd(k)
}

func main() {
	in := getInput()

	fmt.Printf("PART 1: %d\n", part1(in))

	fmt.Printf("PART 2: %d\n", part2(in))
}
