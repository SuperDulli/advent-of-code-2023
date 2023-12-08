package main

import (
	"aoc2023/util"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type player struct {
	hand string
	bid  int
}

func main() {
	lines := util.ReadLines(os.Args[1])
	var players []player
	for _, line := range lines {
		cards := strings.Fields(line)[0]
		bid := util.ConvertToNumber(strings.Fields(line)[1])
		players = append(players, player{cards, bid})
	}

	slices.SortFunc(players, func(a, b player) int {
		return compareHands(a.hand, b.hand)
	})
	fmt.Println(players)

	sum := 0
	for rank, player := range players {
		fmt.Println(rank, player.hand, player.bid, getRank(countCards(player.hand)))
		sum += (rank + 1) * player.bid
	}
	fmt.Println(sum)
}

func countCards(cards string) map[rune]int {
	counts := make(map[rune]int)
	for _, card := range cards {
		counts[card] += 1
	}
	return counts
}

// 0 if equal, -1 a < b, 1 if a > b
func compareHands(a, b string) int {
	if a == b {
		return 0
	}
	aCount := countCards(a)
	bCount := countCards(b)
	rankDifference := getRank(aCount) - getRank(bCount)
	if rankDifference > 0 {
		return -1
	} else if rankDifference < 0 {
		return 1
	}
	i := 0
	for a[i] == b[i] {
		i++
	}
	order := "AKQJT98765432"
	return strings.Index(order, b[i:i+1]) - strings.Index(order, a[i:i+1])
}

func getRank(hand map[rune]int) int {
	len := len(hand)
	switch len {
	case 1:
		// five of a kind
		return 1
	case 2:
		for _, card := range hand {
			if card == 4 {
				// four of a kind
				return 2
			}
		}
		// full house
		return 3
	case 3:
		for _, card := range hand {
			if card == 3 {
				// three of a kind
				return 4
			}
		}
		// two pair
		return 5
	case 4:
		// one pair
		return 6
	case 5:
		// high card
		return 7
	default:
		log.Fatal("failed to identify hand rank")
	}
	return -1
}
