package main

import (
	"aoc2023/util"
	"fmt"
	"os"
)

type node struct {
	left    string
	right   string
	visited bool
}

func main() {
	lines := util.ReadLines(os.Args[1])
	directions := lines[0]
	nodes := make(map[string]node)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		nodes[line[0:3]] = node{
			line[7:10],
			line[12:15],
			false,
		}
	}

	pos := "AAA"
	steps := 0
	for pos != "ZZZ" {
		move := directions[steps%len(directions)]
		if move == 'L' {
			pos = nodes[pos].left
		} else {

			pos = nodes[pos].right
		}
		steps++
	}
	fmt.Println(steps)

	// part 2

	var positions []string
	for start := range nodes {
		if start[2] == 'A' {
			positions = append(positions, start)
		}
	}

	var loops []int
	for _, pos := range positions {
		loops = append(loops, getLoopLen(pos, nodes, directions))
	}

	fmt.Println(LCM(loops[0], loops[1], loops[2:]...)) // the end of each loop is a node with Z at the end
}

func getLoopLen(pos string, nodes map[string]node, directions string) int {
	steps := 0
	for pos[2] != 'Z' {
		node := nodes[pos]
		node.visited = true
		nodes[pos] = node
		move := directions[steps%len(directions)]
		if move == 'L' {
			pos = nodes[pos].left
		} else {

			pos = nodes[pos].right
		}
		steps++
	}
	return steps
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
