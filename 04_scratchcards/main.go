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
	var sum, sum2 int
	cardMap := make(map[int]int)
	cards := getCards(lines)
	for i, card := range cards {
		sum += calcPoints(card)
		cardMap[i] += 1
		for j := 0; j < cardMap[i]; j++ {
			for _, id := range card.tickets {
				cardMap[id] += 1
			}
		}
	}
	for _, value := range cardMap {
		sum2 += value
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

type card struct {
	id      int
	win     []int
	numbers map[int]int
	tickets []int
}

func getCards(lines []string) []card {
	var cards []card
	for _, line := range lines {
		splitLine := strings.Split(line, ":")
		game := strings.Split(splitLine[0], " ")
		id := util.ConvertToNumber(game[len(game)-1])
		numbers := strings.Split(splitLine[1], "|")
		win := strings.Split(strings.Trim(numbers[0], " "), " ")
		actual := util.ConvertToNumbers(strings.Split(strings.Trim(numbers[1], " "), " "))
		actualMap := make(map[int]int)
		for _, n := range actual {
			actualMap[n] = 1
		}
		card := card{id: id, win: util.ConvertToNumbers(win), numbers: actualMap}
		card.calcTickets()
		cards = append(cards, card)
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

func (card *card) calcTickets() {
	count := 0
	for _, n := range card.win {
		_, ok := card.numbers[n]
		if ok {
			count++
		}
	}
	for i := card.id; i < card.id+count; i++ {
		card.tickets = append(card.tickets, i)
	}
}
