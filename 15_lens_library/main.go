package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"strings"
)

func main() {
	initSequence := util.ReadLines(os.Args[1])[0]
	steps := strings.FieldsFunc(initSequence, func(r rune) bool {
		return r == ','
	})

	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	fmt.Println(sum)
}

func hash(s string) int {
	currentValue := 0
	for _, value := range s {
		currentValue += int(value)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}
