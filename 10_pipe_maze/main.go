package main

import (
	"aoc2023/util"
	"fmt"
)

type vector struct {
	x int
	y int
}

func main() {
	maze := util.GetCharMatrix("input.txt")
	// maze := util.GetCharMatrix(os.Args[1])

	// find start
	var start vector
	for y, row := range maze {
		for x, c := range row {
			if c == "S" {
				start.x = x
				start.y = y
			}
		}
	}
	fmt.Println(start)

	startDirection := startDirection(start, maze)
	fmt.Println(startDirection)

	pos := start
	dir := startDirection
	steps := 0
	for ok := true; ok; ok = pos != start {
		fmt.Println(steps, pos, maze[pos.y][pos.x], dir)
		pos = vector{pos.x + dir.x, pos.y + dir.y}
		steps++
		dir = moveDirection(pos, dir, maze)
	}
	fmt.Println(steps)
	distance := steps / 2
	if distance%2 == 0 {
		fmt.Println(distance)
	} else {
		fmt.Println(distance + 1)
	}

}

func entryDirection(direction vector) vector {
	if direction.x == 0 {
		return vector{direction.x, -direction.y}
	}
	return vector{-direction.x, direction.y}
}

func getConnections(pipe string) []vector {
	switch pipe {
	case "|":
		return []vector{{0, -1}, {0, 1}}
	case "-":
		return []vector{{-1, 0}, {1, 0}}
	case "L":
		return []vector{{0, -1}, {1, 0}}
	case "J":
		return []vector{{0, -1}, {-1, 0}}
	case "7":
		return []vector{{0, 1}, {-1, 0}}
	case "F":
		return []vector{{0, 1}, {1, 0}}
	case "S":
		return []vector{{0, 0}, {0, 0}} // don't move when finish is reached
	}
	fmt.Printf("%s has no connections\n", pipe)
	return []vector{}
}

func startDirection(pos vector, maze [][]string) vector {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y != 0 || x != 0 && y == 0 {
				dir := entryDirection(vector{x, y})
				connections := getConnections(maze[pos.y+y][pos.x+x])
				if len(connections) < 2 {
					continue
				}
				if dir == connections[0] || dir == connections[1] {
					return vector{x, y}
				}

			}
		}
	}
	return vector{}
}

func moveDirection(pos, dir vector, maze [][]string) vector {
	entry := entryDirection(dir)
	connections := getConnections(maze[pos.y][pos.x])
	if entry == connections[0] {
		return connections[1]
	} else if entry == connections[1] {
		return connections[0]
	}
	return vector{}
}
