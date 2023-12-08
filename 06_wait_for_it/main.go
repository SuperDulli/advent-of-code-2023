package main

import (
	"aoc2023/util"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadLines(os.Args[1])
	times := util.ConvertToNumbers(strings.Split(strings.Split(lines[0], ":")[1], " "))
	distances := util.ConvertToNumbers(strings.Split(strings.Split(lines[1], ":")[1], " "))

	product := 1
	for i, time := range times {
		distance := distances[i]
		buttonHoldTimes := solveQuadratic(time, distance+1)
		product *= int(math.Floor(buttonHoldTimes[0]) - math.Ceil(buttonHoldTimes[1]) + 1)
	}

	fmt.Println(product)

	// part 2

	time, err := strconv.Atoi(strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), ""))
	if err != nil {
		log.Fatal(err)
	}
	distance, err := strconv.Atoi(strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), ""))
	if err != nil {
		log.Fatal(err)
	}
	buttonHoldTimes := solveQuadratic(time, distance+1)
	fmt.Println(int(math.Floor(buttonHoldTimes[0]) - math.Ceil(buttonHoldTimes[1]) + 1))
}

func solveQuadratic(p, q int) []float64 {
	x1 := float64(p)/2 + math.Sqrt(math.Pow(float64(p)/2, float64(2))-float64(q))
	x2 := float64(p)/2 - math.Sqrt(math.Pow(float64(p)/2, float64(2))-float64(q))
	return []float64{x1, x2}
}
