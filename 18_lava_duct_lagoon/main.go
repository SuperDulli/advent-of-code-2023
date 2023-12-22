package main

import (
	"aoc2023/util"
	"fmt"
	"strings"
)

type vector struct {
	x int
	y int
}

func main() {
	lines := util.ReadLines("example.txt")

	pos := vector{0,0}
	terrain := make(map[vector]int, 0)
	for _, line := range lines {
		terrain[pos] = 1
		splitLine := strings.Fields(line)
		direction := vector{}
		switch splitLine[0] {
		case "R":
			direction = vector{1,0}
		case "L":
			direction = vector{-1,0}
		case "U":
			direction = vector{0,-1}
		case "D":
			direction = vector{0,1}
		}

		// move
		for i := 0; i < util.ConvertToNumber(splitLine[1]); i++ {
			fmt.Println(pos)
			pos.x += direction.x
			pos.y += direction.y
			terrain[pos] = 1
		}
	}
	fmt.Println(terrain) // boundary
}
