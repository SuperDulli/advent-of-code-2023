package main

import (
	"aoc2023/util"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/go-camp/interval"
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
	fmt.Println(seeds)
	fmt.Println(stages)
	fmt.Println(slices.Min(destinations))

	part2(seeds, stages)
}

func part2(rawSeeds []int, rawStages [][]transform) {
	seeds := interval.OrderedSet{}
	stages := [][]transformRange{}

	for i := 0; i < len(rawSeeds)-1; i += 2 {
		seeds.Add(interval.Interval{
			Begin:    rawSeeds[i],
			IncBegin: true,
			End:      rawSeeds[i] + rawSeeds[i+1],
			IncEnd:   false,
		})
	}

	for _, rawStage := range rawStages {
		stage := []transformRange{}
		for _, instruction := range rawStage {
			stage = append(stage, transformRange{
				distance: instruction.dest_start - instruction.source_start,
				source: interval.Interval{
					Begin:    instruction.source_start,
					IncBegin: true,
					End:      instruction.source_start + instruction.range_len,
					IncEnd:   false,
				},
			})
		}
		stages = append(stages, stage)
	}

	fmt.Println(seeds.String())
	fmt.Println(stages)

	// transform the seed ranges
	for n, stage := range stages {
		fmt.Println(n, stage)
		transformedSeeds := interval.OrderedSet{}
		for _, seed := range seeds.Intervals() {
			for _, instruction := range stage {
				fmt.Println(seed)
				fmt.Println("instruction", instruction)
				if instruction.source.Contains(seed) {
					// move seed
					transformedSeeds.Add(seed.Move(instruction.distance))
					fmt.Println("contained!")
					seeds.Remove(seed)
					break
				}
				intersection := seed.Intersect(instruction.source)
				before, after := seed.Bisect(intersection)
				if intersection.IsEmpty() {
					continue
				}
				fmt.Println(intersection)
				fmt.Println(before, after)
				// move only intersection
				transformedSeeds.Add(intersection.Move(instruction.distance))
				if !before.IsEmpty() {
					transformedSeeds.Add(before)
				}
				if !after.IsEmpty() {
					transformedSeeds.Add(after)
				}
				fmt.Println(transformedSeeds)
				seeds.Remove(seed)
				break
			}
		}
		// add not transformed
		for _, seed := range seeds.Intervals() {
			transformedSeeds.Add(seed)
		}
		seeds = transformedSeeds
		fmt.Println(seeds)
	}
	fmt.Println(seeds.Bound().Begin)
}

type transform struct {
	dest_start   int
	source_start int
	range_len    int
}

type transformRange struct {
	distance int
	source   interval.Interval
}
