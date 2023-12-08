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
	var sum, sum2 int
	for _, line := range lines {
		sum += ExtractNumber(line)
		sum2 += ExtractNumber(SpellingToNumber(PrepareLine(line)))
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

type dict struct {
	key string
	value string
}

func PrepareLine(line string) string {
	patterns := []dict{
		{"twone", "twoone"},
		{"oneight", "oneeight"},
		{"threeight", "threeeight"},
		{"fiveight", "fiveeight"},
		{"nineight", "nineeight"},
		{"eightwo", "eighttwo"},
		{"eighthree", "eightthree"},
	}
	for _, pattern := range patterns {
		replaceRegex := regexp.MustCompile(pattern.key)
		line = replaceRegex.ReplaceAllString(line, pattern.value)
	}
	return line
}


func SpellingToNumber(line string) string {
	patterns := []string{"one","two","three","four","five","six","seven","eight","nine"}
	for index, pattern := range patterns {
		numberPattern := regexp.MustCompile(pattern)
		line = numberPattern.ReplaceAllString(line, strconv.Itoa(index + 1))
	}
	return line
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