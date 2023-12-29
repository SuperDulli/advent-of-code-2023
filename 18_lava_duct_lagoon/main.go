package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"strings"
)

type vector struct {
	x int
	y int
}

func main() {
	lines := util.ReadLines(os.Args[1])

	pos := vector{0, 0}
	corners := []vector{}
	boundary := 0
	for _, line := range lines {
		splitLine := strings.Fields(line)
		direction := vector{}
		switch splitLine[0] {
		case "R":
			direction = vector{1, 0}
		case "L":
			direction = vector{-1, 0}
		case "U":
			direction = vector{0, -1}
		case "D":
			direction = vector{0, 1}
		}

		// move
		steps := util.ConvertToNumber(splitLine[1])
		pos.x += steps * direction.x
		pos.y += steps * direction.y
		boundary += steps
		corners = append(corners, pos)
	}
	if boundary%2 != 0 {
		panic("boundary is not divisible by 2")
	}

	// shoelace formula
	area := 0
	for i := 0; i < len(corners); i++ {
		x1index := i - 1
		if x1index < 0 {
			x1index = len(corners) - 1
		}
		x2index := i + 1
		if x2index >= len(corners) {
			x2index = 0
		}
		area += corners[i].y * (corners[x1index].x - corners[x2index].x)
	}
	if area%2 != 0 {
		panic("area is not an integer")
	}
	area = area / 2

	// Pick's theorem for points
	inside := area - boundary/2 + 1

	fmt.Println(inside + boundary)
}
