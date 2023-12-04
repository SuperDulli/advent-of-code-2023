package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	lines := util.ReadLines(os.Args[1])
	// lines := util.ReadLines("example.txt")
	var sum int
	cards := getCards(lines)
	for _, card := range cards {
		sum += calcPoints(card)
	}
	fmt.Println(sum)
}

type card struct {
	win     []int
	numbers map[int]int
}

func getCards(lines []string) []card {
	var cards []card
	for _, line := range lines {
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		win := strings.Split(strings.Trim(numbers[0], " "), " ")
		actual := util.ConvertToNumber(strings.Split(strings.Trim(numbers[1], " "), " "))
		actualMap := make(map[int]int)
		for _, n := range actual {
			actualMap[n] = 1
		}
		cards = append(cards, card{win: util.ConvertToNumber(win), numbers: actualMap})
	}
	return cards
}

func calcPoints(card card) int {
	count := 0
	for _, n := range card.win {
		_, ok := card.numbers[n]
		if ok {
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return int(math.Pow(2, float64(count-1)))
}
