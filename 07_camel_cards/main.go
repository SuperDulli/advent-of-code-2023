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

const (
	FIVE_OF_A_KIND = iota
	FOUR_OF_A_KIND
	FULL_HOUSE
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

func main() {
	lines := util.ReadLines(os.Args[1])
	var players []player
	for _, line := range lines {
		cards := strings.Fields(line)[0]
		bid := util.ConvertToNumber(strings.Fields(line)[1])
		players = append(players, player{cards, bid})
	}

	slices.SortFunc(players, func(a, b player) int {
		return compareHands(a.hand, b.hand, "AKQJT98765432", false)
	})

	sum := 0
	for rank, player := range players {
		sum += (rank + 1) * player.bid
	}
	fmt.Println(sum)

	// part 2

	slices.SortFunc(players, func(a, b player) int {
		return compareHands(a.hand, b.hand, "AKQT98765432J", true)
	})

	sum = 0
	for rank, player := range players {
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
func compareHands(a, b string, order string, useJoker bool) int {
	if a == b {
		return 0
	}
	aCount := countCards(a)
	bCount := countCards(b)
	rankDifference := getRank(aCount) - getRank(bCount)
	if useJoker {
		rankDifference = getRankWithJoker(aCount) - getRankWithJoker(bCount)
	}
	if rankDifference > 0 {
		return -1
	} else if rankDifference < 0 {
		return 1
	}
	i := 0
	for a[i] == b[i] {
		i++
	}
	return strings.Index(order, b[i:i+1]) - strings.Index(order, a[i:i+1])
}

func getRank(hand map[rune]int) int {
	len := len(hand)
	switch len {
	case 1:
		// five of a kind
		return FIVE_OF_A_KIND
	case 2:
		for _, card := range hand {
			if card == 4 {
				// four of a kind
				return FOUR_OF_A_KIND
			}
		}
		// full house
		return FULL_HOUSE
	case 3:
		for _, card := range hand {
			if card == 3 {
				// three of a kind
				return THREE_OF_A_KIND
			}
		}
		// two pair
		return TWO_PAIR
	case 4:
		// one pair
		return ONE_PAIR
	case 5:
		// high card
		return HIGH_CARD
	default:
		log.Fatal("failed to identify hand rank")
	}
	return -1
}

func getRankWithJoker(hand map[rune]int) int {
	len := len(hand)
	jokerCount := hand['J']
	switch jokerCount {
	case 5:
		fallthrough
	case 4:
		// five of a kind
		return FIVE_OF_A_KIND
	case 3:
		if len == 2 {
			// five of a kind
			return FIVE_OF_A_KIND
		}
		return FOUR_OF_A_KIND
	case 2:
		// JJAAA -> five of a kind
		// JJAAX -> (Fullhouse)/ four of a kind
		// JJABC -> Three of a kind
		if len == 4 {
			return THREE_OF_A_KIND
		}
		if len == 3 {
			return FOUR_OF_A_KIND
		}
		return FIVE_OF_A_KIND
	case 1:
		switch len {
		case 5:
			return ONE_PAIR
		case 4:
			return THREE_OF_A_KIND
		case 3:
			for _, card := range hand {
				if card == 3 {
					// JAAAB 4
					return FOUR_OF_A_KIND
				}
			}
			// JAABB FH
			return FULL_HOUSE
		case 2:
			return FIVE_OF_A_KIND
		}
		// JAABC 3
		// JABCD 1
	case 0:
		return getRank(hand)
	default:
		log.Fatal("failed to identify hand rank")
	}
	return -1
}
