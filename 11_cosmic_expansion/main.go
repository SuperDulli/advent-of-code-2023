package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"os"
)

func main() {
	matrix := util.GetCharMatrix(os.Args[1])

	// find empty rows
	var emptyRows, emptyCols []int
	for n, row := range matrix {
		if util.All(row, ".") {
			emptyRows = append(emptyRows, n)
		}
	}
	matrix = util.Transpose(matrix)
	for n, row := range matrix {
		if util.All(row, ".") {
			emptyCols = append(emptyCols, n)
		}
	}
	matrix = util.Transpose(matrix)

	// find all galaxies
	var galaxies []vector
	for y, row := range matrix {
		for x, tile := range row {
			if tile == "#" {
				galaxies = append(galaxies, vector{x, y})
			}
		}
	}

	sumDistances(galaxies, emptyRows, emptyCols, 1)       // part 1
	sumDistances(galaxies, emptyRows, emptyCols, 1000000) // part 2
}

type vector struct {
	x int
	y int
}

func manhattan(start, end vector, emptyRows, emptyCols []int, factor int) int {
	distance := int(math.Abs(math.Abs(float64(start.x)-float64(end.x)) + math.Abs(float64(start.y)-float64(end.y))))
	xMin := int(math.Min(float64(start.x), float64(end.x)))
	yMin := int(math.Min(float64(start.y), float64(end.y)))
	xMax := int(math.Max(float64(start.x), float64(end.x)))
	yMax := int(math.Max(float64(start.y), float64(end.y)))
	for _, row := range emptyRows {
		if row > yMin && row < yMax {
			distance += factor - 1
		}
	}
	for _, col := range emptyCols {
		if col > xMin && col < xMax {
			distance += factor - 1
		}
	}
	return distance
}

// compute distances between galaxies
func sumDistances(galaxies []vector, emptyRows, emptyCols []int, factor int) int {
	var sum int
	for i, startGalaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			endGalaxy := galaxies[j]
			distance := manhattan(startGalaxy, endGalaxy, emptyRows, emptyCols, factor)
			sum += distance
		}
	}
	fmt.Println(sum)
	return sum
}
