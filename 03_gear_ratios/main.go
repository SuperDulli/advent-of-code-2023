package main

import (
	"aoc2023/util"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	matrix := util.GetCharMatrix(os.Args[1])
	parts := getParts(matrix)
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Println(sum)
}

func getParts(matrix [][]string) []string {
	var parts []string
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for row, line := range matrix {
		var part []string
		isTouching := false
		for col, char := range line {
			if _, err := strconv.Atoi(char); err != nil {
				if len(part) > 0 && isTouching {
					parts = append(parts, strings.Join(part, ""))
				}
				part = []string{}
				isTouching = false
				continue
			}
			part = append(part, char)
			if isTouching {
				continue
			}
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				x := row + dx
				y := col + dy
				if !isInbound(matrix, x, y) {
					continue
				}
				if isPart(matrix, x, y) {
					isTouching = true
				}
			}
		}
		if len(part) > 0 && isTouching {
			parts = append(parts, strings.Join(part, ""))
		}
	}
	return parts
}

func isPart(matrix [][]string, x int, y int) bool {
	return strings.Contains("*&$%-+/=@#", matrix[x][y])
}

func isInbound(matrix [][]string, x int, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	height := len(matrix)
	if x >= height {
		return false
	}
	width := len(matrix[x])
	return y < width
}
