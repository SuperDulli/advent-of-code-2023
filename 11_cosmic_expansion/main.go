package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"os"
)

func main() {
	matrix := util.GetCharMatrix(os.Args[1])

	// expand space
	var stretched [][]string
	for _, row := range matrix {
		stretched = append(stretched, row)
		if util.All(row, ".") {
			stretched = append(stretched, row)
		}
	}
	matrix = util.Transpose(stretched)
	stretched = nil
	for _, row := range matrix {
		stretched = append(stretched, row)
		if util.All(row, ".") {
			stretched = append(stretched, row)
		}
	}
	matrix = util.Transpose(stretched)

	// find all galaxies
	var galaxies []vector
	for y, row := range matrix {
		for x, tile := range row {
			if tile == "#" {
				galaxies = append(galaxies, vector{x, y})
			}
		}
	}

	// compute distances between galaxies
	var sum int
	for i, startGalaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			endGalaxy := galaxies[j]
			distance := manhattan(startGalaxy, endGalaxy)
			sum += distance
		}
	}
	fmt.Println(sum)
}

type vector struct {
	x int
	y int
}

func manhattan(start, end vector) int {
	return int(math.Abs(math.Abs(float64(start.x)-float64(end.x)) + math.Abs(float64(start.y)-float64(end.y))))
}
