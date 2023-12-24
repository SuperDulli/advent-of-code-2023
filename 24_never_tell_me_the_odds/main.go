package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"regexp"
)

type vector struct {
	x float32
	y float32
	z float32
}

type path struct {
	pos vector
	dir vector
}

func main() {
	lines := util.ReadLines(os.Args[1])

	paths := []path{}
	for _, line := range lines {
		splitLine := regexp.MustCompile(`\s@\s+`).Split(line, -1)
		pos := util.ConvertToNumbers(regexp.MustCompile(`,\s+`).Split(splitLine[0], -1))
		dir := util.ConvertToNumbers(regexp.MustCompile(`,\s+`).Split(splitLine[1], -1))
		paths = append(paths, path{
			pos: vector{float32(pos[0]), float32(pos[1]), float32(pos[2])},
			dir: vector{float32(dir[0]), float32(dir[1]), float32(dir[2])},
		})
	}

	sum := 0
	for i, a := range paths {
		for j := i + 1; j < len(paths); j++ {
			b := paths[j]
			intersection, intersectionInPast := intersection2D(a, b)
			if (intersection != vector{0, 0, 0}) && !intersectionInPast && isIntersection2DInArea(intersection, vector{200000000000000, 200000000000000, 0}, vector{400000000000000, 400000000000000, 0}) {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func intersection2D(a, b path) (vector, bool) {
	intersectionInPast := false
	det := b.dir.x*a.dir.y - b.dir.y*a.dir.x
	if det == 0 {
		return vector{}, intersectionInPast // no intersection
	}
	scalarA := (b.dir.x*(b.pos.y-a.pos.y) + b.dir.y*(a.pos.x-b.pos.x)) / det
	scalarB := (a.dir.x*(b.pos.y-a.pos.y) + a.dir.y*(a.pos.x-b.pos.x)) / det
	vecB := vector{
		x: b.pos.x + scalarB*b.dir.x,
		y: b.pos.y + scalarB*b.dir.y,
	}
	if scalarA < 0 || scalarB < 0 {
		intersectionInPast = true
	}
	return vecB, intersectionInPast
}

func isIntersection2DInArea(pos, areaStart, areaEnd vector) bool {
	return pos.x >= areaStart.x && pos.x <= areaEnd.x && pos.y >= areaStart.y && pos.y <= areaEnd.y
}
