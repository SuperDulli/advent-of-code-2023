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
		sum += countPossible(riddle, util.ConvertToNumbers(hints))
	}
	fmt.Println(sum)
}

func countPossible(pattern string, hints []int) int {
	states := "."
	for _, hint := range hints {
		for i := 0; i < hint; i++ {
			states += "#"
		}
		states += "."
	}

	stateMap := make(map[int]int)
	stateMap[0] = 1

	// NFA
	newStateMap := make(map[int]int)
	for _, char := range pattern {
		for state, _ := range stateMap {
			switch char {
			case '.':
				// move to neighbor dot
				if state+1 < len(states) && states[state+1] == '.' {
					newStateMap[state+1] = newStateMap[state+1] + stateMap[state]
				}
				// stay on the dot
				if states[state] == '.' {
					newStateMap[state] = newStateMap[state] + stateMap[state]
				}
			case '#':
				// move to neighbor hash
				if state+1 < len(states) && states[state+1] == '#' {
					newStateMap[state+1] = newStateMap[state+1] + stateMap[state]
				}
			case '?':
				if state+1 < len(states) {
					newStateMap[state+1] = newStateMap[state+1] + stateMap[state]
				}
				if states[state] == '.' {
					newStateMap[state] = newStateMap[state] + stateMap[state]
				}
			}
		}
		for k := range stateMap {
			delete(stateMap, k)
		}
		for k, v := range newStateMap {
			stateMap[k] = v
			delete(newStateMap, k)
		}
	}

	// result is the combined count in the last two states
	return stateMap[len(states)-1] + stateMap[len(states)-2]
}
