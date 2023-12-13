package main

import (
	"aoc2023/util"
	"fmt"
	"os"
)

func main() {
	// lines := util.ReadLines("example.txt")
	lines := util.ReadLines(os.Args[1])

	var patterns [][][]string
	var patternLines []string
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, util.ConvertToMatrix(patternLines))
			patternLines = nil
			continue
		}
		patternLines = append(patternLines, line)
	}
	patterns = append(patterns, util.ConvertToMatrix(patternLines))

	sum := 0
	for _, pattern := range patterns {
		reflectNumber := reflectNumber(pattern)
		sum += reflectNumber
	}
	fmt.Println(sum)
}

func reflectNumber(pattern [][]string) int {
	verticalNumber := findReflection(pattern)
	pattern = util.Transpose(pattern)
	horizontalNumber := findReflection(pattern)
	return verticalNumber + 100*horizontalNumber
}

func findReflection(pattern [][]string) int {
	options := make(map[int]int)
	for i := 0; i < len(pattern[0])-1; i++ {
		options[i] = 1
	}

	for _, row := range pattern {
		for i := 1; i < len(row)-2; i++ {
			for option, _ := range options {
				if option+i >= len(row) || option-i+1 < 0 {
					continue
				}
				left := row[option-i+1]
				right := row[option+i]
				if left != right {
					if len(options) == 1 {
						return 0
					}
					delete(options, option)
					continue
				}
			}
		}
	}
	for option := range options {
		return option + 1
	}
	return 0
}
