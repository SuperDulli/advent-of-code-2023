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
	var sum, sum2 int
	for _, line := range lines {
		data := getData(line)
		sum += gameIsPossible(data)
		sum2 += powerSet(minSet(data.sets))
	}
	fmt.Println(sum)
	fmt.Println(sum2)
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

// part 2

func minSet(sets []cubeSet) cubeSet {
	var min cubeSet
	for _, set := range sets {
		if set.red > min.red {
			min.red = set.red
		}
		if set.green > min.green {
			min.green = set.green
		}
		if set.blue > min.blue {
			min.blue = set.blue
		}
	}
	return min
}

func powerSet(set cubeSet) int {
	return set.red * set.green * set.blue
}
