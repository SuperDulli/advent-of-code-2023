package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	FLIP_FLOP = iota
	CONJUNCTION
	BROADCAST
)

type Module struct {
	kind         int
	name         string
	state        int
	inputs       map[string]int
	destinations []string
}

type Event struct {
	source string
	signal int
	dest   string
}

func main() {
	lines := util.ReadLines(os.Args[1])

	modules := make(map[string]Module, 0)
	for _, line := range lines {
		splitLine := strings.Split(line, " -> ")
		module := Module{
			destinations: regexp.MustCompile(`,\W`).Split(splitLine[1], -1),
			state:        -1, // low
		}
		module.inputs = make(map[string]int)
		switch line[0] {
		case '%':
			module.name = splitLine[0][1:]
			module.kind = FLIP_FLOP
		case '&':
			module.name = splitLine[0][1:]
			module.kind = CONJUNCTION
		default:
			module.name = splitLine[0]
			module.kind = BROADCAST
		}
		modules[module.name] = module
	}
	// find all inputs
	for moduleName, module := range modules {
		for _, destination := range module.destinations {
			_, ok := modules[destination]
			if !ok {
				continue
			}
			modules[destination].inputs[moduleName] = -1
		}
	}

	low := 0
	high := 0

	var queue []Event
	for i := 0; i < 1000; i++ {

		queue = []Event{{
			"button",
			-1,
			"broadcaster",
		}}
		for len(queue) > 0 {
			event := queue[0]
			queue = queue[1:]

			module := modules[event.dest]
			signal := module.process(event.signal, event.source)
			modules[event.dest] = module // save changed state
			switch event.signal {
			case -1:
				low++
			case 1:
				high++
			}
			if signal == 0 {
				continue
			}
			for _, dest := range modules[event.dest].destinations {
				queue = append(queue, Event{event.dest, signal, dest})
			}
		}
	}
	fmt.Println(low * high)
}

// 1 high, -1 low
func (module *Module) process(signal int, source string) int {
	switch module.kind {
	case FLIP_FLOP:
		if signal == -1 {
			module.state *= -1 // toggle
			return module.state
		}
	case CONJUNCTION:
		module.inputs[source] = signal
		for _, value := range module.inputs {
			if value != 1 {
				return 1
			}
		}
		return -1
	case BROADCAST:
		return signal
	}
	return 0 // nothing happens
}
