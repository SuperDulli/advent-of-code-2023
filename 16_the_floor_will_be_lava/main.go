package main

import (
	"aoc2023/util"
	"fmt"
	"os"
)

type vector struct {
	x int
	y int
}

type Beam struct {
	pos vector
	dir vector
}

func main() {
	floor := util.GetCharMatrix(os.Args[1])

	// copies of floor to keep track of visited tiles
	energized := make([][]bool, len(floor))
	beams := make([][]vector, len(floor)) // remember the beams direction
	for i := range floor {
		energized[i] = make([]bool, len(floor[0]))
		beams[i] = make([]vector, len(floor[0]))
	}

	queue := []Beam{{
		pos: vector{0, 0},
		dir: vector{1, 0},
	}}

	for len(queue) != 0 {
		beam := queue[0]
		queue = queue[1:] // remove top element

		if beam.pos.x < 0 || beam.pos.x >= len(floor[0]) || beam.pos.y < 0 || beam.pos.y >= len(floor) {
			continue
		}

		if energized[beam.pos.y][beam.pos.x] && beams[beam.pos.y][beam.pos.x] == beam.dir {
			continue
		}
		energized[beam.pos.y][beam.pos.x] = true
		beams[beam.pos.y][beam.pos.x] = beam.dir

		switch floor[beam.pos.y][beam.pos.x] {
		case "|":
			if beam.dir.x != 0 {
				queue = append(queue, Beam{
					pos: vector{beam.pos.x, beam.pos.y - 1},
					dir: vector{0, -1},
				})
				queue = append(queue, Beam{
					pos: vector{beam.pos.x, beam.pos.y + 1},
					dir: vector{0, 1},
				})
				continue
			}
		case "-":
			if beam.dir.y != 0 {
				queue = append(queue, Beam{
					pos: vector{beam.pos.x - 1, beam.pos.y},
					dir: vector{-1, 0},
				})
				queue = append(queue, Beam{
					pos: vector{beam.pos.x + 1, beam.pos.y},
					dir: vector{1, 0},
				})
				continue
			}
		case "\\":
			newDir := vector{beam.dir.y, beam.dir.x} // swap x and y
			queue = append(queue, Beam{
				pos: vector{beam.pos.x + newDir.x, beam.pos.y + newDir.y},
				dir: newDir,
			})
			continue
		case "/":
			newDir := vector{-beam.dir.y, -beam.dir.x} // swap x and y
			queue = append(queue, Beam{
				pos: vector{beam.pos.x + newDir.x, beam.pos.y + newDir.y},
				dir: newDir,
			})
			continue
		}

		// no reflection or splitting
		beam.pos.x += beam.dir.x
		beam.pos.y += beam.dir.y
		queue = append(queue, beam)

	}

	sum := 0
	for _, row := range energized {
		for _, tile := range row {
			if tile {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
