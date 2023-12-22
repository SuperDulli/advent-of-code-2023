package main

import (
	"aoc2023/util"
	"cmp"
	"fmt"
	"slices"
)

type vector struct {
	x int
	y int
}

func main() {
	matrix := util.GetNumberMatrix("example.txt")

	util.Print2D(matrix)

	distance := make([][]int, len(matrix))
	prev := make(map[vector]vector)
	queue := []vector{}
	for i, _ := range distance {
		distance[i] = make([]int, len(matrix[i]))
		for j, _ := range matrix[i] {
			distance[i][j] = -1 // infinity
			queue = append(queue, vector{j, i})
		}
	}
	distance[0][0] = 0
	util.Print2D(distance)
	fmt.Println(prev)

	for len(queue) > 0 {
		// sort queue shortest distance at front
		slices.SortFunc(queue, func(a, b vector) int {
			if distance[a.y][a.x] == -1 {
				return 1
			}
			if distance[b.y][b.x] == -1 {
				return -1
			}
			return cmp.Compare(distance[a.y][a.x], distance[b.y][b.x])
		})
		// fmt.Println("q",queue)

		visit := queue[0]
		queue = queue[1:]
		fmt.Println("visit", visit)

		neighbors := neighbors(visit, prev, distance)
		for _, neighbor := range neighbors {
			if !slices.Contains(queue, neighbor) {
				continue
			}
			pathLength := distance[visit.y][visit.x] + matrix[neighbor.y][neighbor.x]
			fmt.Println(neighbor, pathLength)
			if distance[neighbor.y][neighbor.x] == -1 || pathLength < distance[neighbor.y][neighbor.x] {
				fmt.Println("shorter path found")
				oldLength := distance[neighbor.y][neighbor.x]
				distance[neighbor.y][neighbor.x] = pathLength
				prev[neighbor] = visit

				// check if not moved in the same direction for the last 3 moves
				path := reconstructPath(neighbor, prev)
				if len(path) >= 5 {
					xDifference := path[0].x - path[4].x
					yDifference := path[0].y - path[4].y
					if (xDifference == 4 && yDifference == 0) || (xDifference == 0 && yDifference == 4) {
						fmt.Println("skip", path)
						distance[neighbor.y][neighbor.x] = oldLength
						delete(prev, neighbor)
					}
				}
				util.Print2D(distance)
			}
		}
	}
	util.Print2D(distance)

	// reconstruct path
	pos := vector{len(matrix) - 1, len(matrix[0]) - 1}
	path := reconstructPath(pos, prev)
	// reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	fmt.Println(path)
}

func neighbors(pos vector, prev map[vector]vector, distance [][]int) []vector {
	result := []vector{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if (x == 0 && y == 0) || (x != 0 && y != 0) {
				continue
			}
			actualX := pos.x + x
			actualY := pos.y + y
			if actualX < 0 || actualX >= len(distance[0]) || actualY < 0 || actualY >= len(distance) {
				continue
			}
			neighbor := vector{actualX, actualY}

			result = append(result, neighbor)
		}

	}
	return result
}

func reconstructPath(target vector, prev map[vector]vector) []vector {
	path := []vector{}
	ok := true
	for ok {
		path = append(path, target)
		target, ok = prev[target]
	}
	return path
}
