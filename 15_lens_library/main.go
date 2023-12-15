package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

type lens struct {
	label string
	value int
}

func main() {
	initSequence := util.ReadLines(os.Args[1])[0]
	steps := strings.FieldsFunc(initSequence, func(r rune) bool {
		return r == ','
	})

	labels := make(map[string]int, 0)
	boxes := make([][]lens, 256)

	sum := 0
	for _, step := range steps {
		label := regexp.MustCompile("[a-z]+").FindString(step)
		_, ok := labels[label]
		if !ok {
			labels[label] = hash(label)
		}
		n := labels[label]
		switch regexp.MustCompile("[^a-z]+").FindString(step)[0] {
		case '=':
			index := slices.IndexFunc(boxes[n], func(l lens) bool {
				return l.label == label
			})
			number := util.ConvertToNumber(regexp.MustCompile("[0-9]+").FindString(step))
			if index >= 0 {
				boxes[n][index].value = number
			} else {
				boxes[n] = append(boxes[n], lens{label, number})
			}
		case '-':
			indexToRemove := slices.IndexFunc(boxes[n], func(l lens) bool {
				return l.label == label
			})
			if indexToRemove >= 0 {
				boxes[n] = util.RemoveStable(boxes[n], indexToRemove)
			}
		}
		sum += hash(step)
	}
	fmt.Println(sum)

	focusingPower := 0
	for n, box := range boxes {
		for i, slot := range box {
			focusingPower += (n + 1) * (i + 1) * slot.value
		}
	}
	fmt.Println(focusingPower)
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
