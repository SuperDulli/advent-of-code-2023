package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := util.ReadLines(os.Args[1])

	sum := 0
	for _, line := range lines {
		var differences [][]int
		history := util.ConvertToNumbers(strings.Fields(line))
		differences = append(differences, history)
		diff := calcDifference(history)
		differences = append(differences, diff)
		for !all(diff, 0) {
			diff = calcDifference(diff)
			differences = append(differences, diff)
		}
		fmt.Println(differences)

		increment := 0
		for i := len(differences) - 2; i >= 0; i-- {
			increment += differences[i][len(differences[i])-1]
		}
		fmt.Println(increment)

		sum += increment
	}

	fmt.Println(sum)
}

func calcDifference(history []int) []int {
	var diff []int
	for i := 0; i < len(history)-1; i++ {
		diff = append(diff, history[i+1]-history[i])
	}
	return diff
}

func all(arr []int, value int) bool {
	for _, v := range arr {
		if v != value {
			return false
		}
	}
	return true
}
