package main

import (
	"aoc2023/util"
	"fmt"
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
	garden := util.GetCharMatrix(os.Args[1])

	start := Vector{}

	distances := make([][]int, len(garden))
	for i := range distances {
		distances[i] = make([]int, len(garden[i]))
		for j := range distances[i] {
			distances[i][j] = -1
			if garden[i][j] == "S" {
				start = Vector{j, i}
				continue
			}
		}
	}


	// flood fill the garden to get the distance to each tile from the start
	steps := 64

	queue := []QueuePos{{0, start}}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		if pos.distance > steps{
			break
		}
		if distances[pos.pos.y][pos.pos.x] != -1 {
			continue
		}
		distances[pos.pos.y][pos.pos.x] = pos.distance
		neighbors := neighbors(pos.pos, garden, distances)

		for _, neighbor := range neighbors {
			queue = append(queue, QueuePos{pos.distance + 1, neighbor})
		}

	}
	fmt.Println(reachable(distances))

}

func neighbors(pos Vector, garden [][]string, distances [][]int) []Vector {
	var neighbors []Vector
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if (x == 0 && y == 0) || (x != 0 && y != 0){
				continue
			}
			nx := pos.x + x
			ny := pos.y + y
			if nx < 0 || nx > len(garden[0])-1 || ny < 0 || ny > len(garden)-1 {
				continue
			}
			if garden[ny][nx] != "." || distances[ny][nx] != -1 {
				continue
			}
			neighbors = append(neighbors, Vector{nx, ny})
		}
	}
	return neighbors
}

// if the distance is a multiple of 2 it can be reached when the number of steps is also a multiple of 2
func reachable(distances [][]int) int  {
	count := 0
	for _, row := range distances {
		for _, tile := range row {
			if tile % 2 == 0 {
				count++
			}
		}
	}
	return count
}
