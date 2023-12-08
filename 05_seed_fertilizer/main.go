package main

import (
	"aoc2023/util"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	lines := util.ReadLines(os.Args[1])
	var seeds []int
	var stages [][]transform
	var maps []transform
	for n, line := range lines {
		if n == 0 {
			seeds = util.ConvertToNumbers(strings.Split(strings.Split(line, ": ")[1], " "))
			continue
		}
		if line == "" || line[len(line)-1] == ':' {
			if len(maps) > 0 {
				slices.SortFunc(maps, func(a, b transform) int {
					return cmp.Compare(a.source_start, b.source_start)
				})
				stages = append(stages, maps)
			}
			maps = make([]transform, 0)
			continue
		}
		numbers := util.ConvertToNumbers(strings.Split(line, " "))
		maps = append(maps, transform{numbers[0], numbers[1], numbers[2]})
	}
	stages = append(stages, maps)
	var destinations []int
	for _, seed := range seeds {
		for _, stage := range stages {
			for _, m := range stage {
				if seed >= m.source_start && seed < m.source_start+m.range_len {
					seed = m.dest_start + (seed - m.source_start)
					break
				}
			}
		}
		destinations = append(destinations, seed)
	}
	fmt.Println(slices.Min(destinations))
}

type transform struct {
	dest_start   int
	source_start int
	range_len    int
}
