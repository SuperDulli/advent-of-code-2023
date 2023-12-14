package main

import (
	"aoc2023/util"
	"fmt"
	"os"
)

func main() {
	matrix := util.GetCharMatrix(os.Args[1])
	matrix = util.Transpose(matrix)

	sum := 0
	for _, row := range matrix {
		sum += countRocks(row)
	}
	fmt.Println(sum)
}

func countRocks(row []string) int {
	sum := 0
	lastDistance := len(row)
	for i, tile := range row {
		if tile == "#" {
			lastDistance = len(row) - i - 1
			continue
		}
		if tile == "O" {
			sum += lastDistance
			lastDistance--
		}
	}
	return sum
}
