package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"strings"
)

type WorkflowStep struct {
	feature     string
	greater     bool
	threshold   int
	destination string
}

type Part struct {
	features map[string]int
}

func main() {
	lines := util.ReadLines(os.Args[1])

	workflows := make(map[string][]WorkflowStep, 0)
	parts := []Part{}

	parseWorkflows := true
	for _, line := range lines {
		if line == "" {
			parseWorkflows = false
			continue
		}
		if parseWorkflows {
			split1 := strings.FieldsFunc(line, func(r rune) bool {
				return r == '{'
			})
			name := split1[0]
			stepsData := strings.FieldsFunc(split1[1][:len(split1[1])-1], func(r rune) bool {
				return r == ','
			})
			var steps []WorkflowStep
			for _, step := range stepsData {
				if !strings.Contains(step, ":") {
					steps = append(steps, WorkflowStep{destination: step})
					break
				}
				stepData := strings.FieldsFunc(step, func(r rune) bool {
					return r == ':'
				})
				symbol := '<'
				if strings.Contains(stepData[0], ">") {
					symbol = '>'
				}
				condition := strings.FieldsFunc(stepData[0], func(r rune) bool {
					return r == symbol
				})
				steps = append(steps, WorkflowStep{
					feature:     condition[0],
					greater:     symbol == '>',
					threshold:   util.ConvertToNumber(condition[1]),
					destination: stepData[1],
				})
			}
			workflows[name] = steps
			continue
		}
		// parse parts
		partData := strings.FieldsFunc(line[1:len(line)-1], func(r rune) bool {
			return r == ','
		})
		var part Part
		part.features = make(map[string]int)
		for _, component := range partData {
			feature := strings.FieldsFunc(component, func(r rune) bool {
				return r == '='
			})
			part.features[feature[0]] = util.ConvertToNumber(feature[1])
		}
		parts = append(parts, part)
	}

	sum := 0
	for _, part := range parts {
		pos := processWorkflow(part, workflows["in"])
		for pos != "A" && pos != "R" {
			pos = processWorkflow(part, workflows[pos])
		}
		if pos == "A" {
			sum += partSum(part)
		}
	}
	fmt.Println(sum)
}

func processWorkflow(part Part, workflow []WorkflowStep) string {
	for _, step := range workflow {
		if step.feature == "" {
			return step.destination
		}
		if step.greater && part.features[step.feature] > step.threshold {
			return step.destination
		}
		if !step.greater && part.features[step.feature] < step.threshold {
			return step.destination
		}
	}
	panic("undefined")
}

func partSum(part Part) int {
	sum := 0
	for _, value := range part.features {
		sum += value
	}
	return sum
}
