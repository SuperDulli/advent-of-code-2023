package main

import (
	"aoc2023/util"
	"fmt"
	"log"
	"os"
)

type Vector struct {
	x int
	y int
}
type QueuePos struct {
	distance int
	pos      Vector
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: file steps")
	}
	garden := util.GetCharMatrix(os.Args[1])

	start := Vector{}

	distances := make(map[Vector]int)
	for i := range garden {
		for j := range garden[i] {
			if garden[i][j] == "S" {
				garden[i][j] = "."
				start = Vector{j, i}
				break
			}
		}
	}

	// flood fill the garden to get the distance to each tile from the start
	steps := util.ConvertToNumber(os.Args[2])

	queue := []QueuePos{{0, start}}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		if pos.distance > steps {
			break
		}
		_, ok := distances[pos.pos]
		if ok {
			continue
		}
		distances[pos.pos] = pos.distance
		neighbors := neighbors(pos.pos, garden)

		for _, neighbor := range neighbors {
			queue = append(queue, QueuePos{pos.distance + 1, neighbor})
		}

	}
	fmt.Println(reachable(distances, steps))

}

func neighbors(pos Vector, garden [][]string) []Vector {
	var neighbors []Vector
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if (x == 0 && y == 0) || (x != 0 && y != 0) {
				continue
			}
			actualX := pos.x + x
			actualY := pos.y + y
			nx := actualX % len(garden[0])
			ny := actualY % len(garden)
			if nx < 0 {
				nx = len(garden[0]) + nx
			}
			if ny < 0 {
				ny = len(garden) + ny
			}

			if garden[ny][nx] != "." {
				continue
			}
			neighbors = append(neighbors, Vector{actualX, actualY})
		}
	}
	return neighbors
}

// if the distance is a multiple of 2 it can be reached when the number of steps is also a multiple of 2
func reachable(distances map[Vector]int, steps int) int {
	count := 0
	for _, distance := range distances {
		if distance%2 == steps % 2 {
			count++
		}

	}
	return count
}
