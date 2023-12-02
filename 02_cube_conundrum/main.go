package main

import (
	util "aoc2023/util"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cubeSet struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []cubeSet
}

func main() {
	lines := util.ReadLines(os.Args[1])
	var sum int
	for _, line := range lines {
		data := getData(line)
		sum += gameIsPossible(data)
	}
	fmt.Println(sum)
}

func getData(line string) game {
	split1 := strings.Split(line, ":")
	splitSets := strings.Split(split1[1], ";")
	var sets []cubeSet
	for _, set := range splitSets {
		cubes := strings.Split(set, ",")
		var cubeSet cubeSet
		for _, cube := range cubes {
			splitCube := strings.Split(strings.TrimSpace(cube), " ")
			amount, err := strconv.Atoi(splitCube[0])
			if err != nil {
				log.Fatal(err)
			}
			switch splitCube[1] {
			case "red":
				cubeSet.red = amount
			case "green":
				cubeSet.green = amount
			case "blue":
				cubeSet.blue = amount
			}
		}
		sets = append(sets, cubeSet)
	}

	id, err := strconv.Atoi(strings.Split(split1[0], " ")[1])
	if err != nil {
		log.Fatal(err)
	}

	return game{id: id, sets: sets}
}

func gameIsPossible(game game) int {
	maxCubes := cubeSet{red: 12, green: 13, blue: 14}
	for _, set := range game.sets {
		if set.red > maxCubes.red || set.green > maxCubes.green || set.blue > maxCubes.blue {
			return 0
		}
	}
	return game.id
}
