package main

import (
	"aoc2023/util"
	"fmt"
	"os"
)

type node struct {
	left  string
	right string
}

func main() {
	// lines := util.ReadLines("example3.txt")

	lines := util.ReadLines(os.Args[1])
	directions := lines[0]
	nodes := make(map[string]node)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		nodes[line[0:3]] = node{
			line[7:10],
			line[12:15],
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

	steps = 0
	for !allFinished(positions) {
		fmt.Println(positions)
		for i, pos := range positions {
			move := directions[steps%len(directions)]
			if move == 'L' {
				pos = nodes[pos].left
			} else {

				pos = nodes[pos].right
			}
			positions[i] = pos
		}
		steps++
	}
	fmt.Println(steps)
}

func allFinished(positions []string) bool {
	for _, pos := range positions {
		if pos[2] != 'Z' {
			return false
		}
	}
	return true
}
