package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := util.ReadLines(os.Args[1])
	var sum int
	for _, line := range lines {
		sum += ExtractNumber(line)
	}
	fmt.Println(sum)
}

func ExtractNumber(line string) int {
	nonNumberRegex := regexp.MustCompile(`[^0-9]+`)
	numbers := nonNumberRegex.ReplaceAllString(line, "")
	numberStr := numbers[0:1] + numbers[len(numbers)-1:]
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		panic(err)
	}
	return number
}