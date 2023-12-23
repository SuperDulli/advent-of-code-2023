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
		lineData := strings.Fields(line)
		riddle := lineData[0]
		hints := strings.FieldsFunc(lineData[1], func(r rune) bool {
			return r == ','
		})
		sum += analyze(riddle, util.ConvertToNumbers(hints), false)
	}
	fmt.Println(sum)
}

func analyze(s string, hints []int, groupEnded bool) int {
	if len(s) == 0 {
		if len(hints) > 0 {

			return 0
		}
		return 1
	}
	switch s[0] {
	case '.':
		return analyze(s[1:], hints, false)
	case '?':
		positive := 0
		if !groupEnded {
			positive = analyze("#"+s[1:], hints, groupEnded)

		}
		negative := analyze("."+s[1:], hints, groupEnded)
		return positive + negative
	case '#':
		// check first group
		if len(hints) == 0 {
			return 0
		}
		for i := 0; i < hints[0]; i++ {
			if i >= len(s) || s[i] == '.' {
				return 0
			}
		}
		if hints[0] < len(s) && s[hints[0]] == '#' {
			return 0 // group is longer than hint
		}
		return analyze(s[hints[0]:], hints[1:], true)
	default:
		panic("default case")
	}
}
