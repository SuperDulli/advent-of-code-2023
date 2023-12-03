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

	// part 2
	// get all numbers that are touching a * and remember the stars position
	gearParts := getGearParts(matrix)
	var sum2 int
	for i, gearPart := range gearParts {
		if gearPart.checked {
			continue
		}
		gearParts[i].checked = true
		gearNumbers := []int{gearPart.number}
		for j, gearPart2 := range gearParts {
			if i == j || gearPart.pos != gearPart2.pos || gearPart2.checked {
				continue
			}
			gearParts[j].checked = true
			gearNumbers = append(gearNumbers, gearPart2.number)
		}
		// if exactly two numbers are touching the same gear add their ratio to the sum
		if len(gearNumbers) == 2 {
			sum2 += gearNumbers[0] * gearNumbers[1]
		}
	}
	println(sum2)
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

// part 2

type cord struct {
	x int
	y int
}

type gearPart struct {
	number  int
	pos     cord
	checked bool
}

func getGearParts(matrix [][]string) []gearPart {
	var parts []gearPart
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for row, line := range matrix {
		var part []string
		isTouching := false
		gearPos := cord{}
		for col, char := range line {
			if _, err := strconv.Atoi(char); err != nil {
				if len(part) > 0 && isTouching {
					number, err := strconv.Atoi(strings.Join(part, ""))
					if err != nil {
						log.Fatal(err)
					}
					parts = append(parts, gearPart{number: number, pos: gearPos})
				}
				part = []string{}
				isTouching = false
				gearPos = cord{}
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
				if "*" == matrix[x][y] {
					isTouching = true
					gearPos.x = x
					gearPos.y = y
				}
			}
		}
		if len(part) > 0 && isTouching {
			number, err := strconv.Atoi(strings.Join(part, ""))
			if err != nil {
				log.Fatal(err)
			}
			parts = append(parts, gearPart{number: number, pos: gearPos})
		}
	}
	return parts
}
