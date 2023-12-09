package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	Hand string
	Bid  int
}

type HandStrength int

const (
	HighCard HandStrength = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

func (hs HandStrength) String() string {
	switch hs {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfKind:
		return "ThreeOfKind"
	case FullHouse:
		return "FullHouse"
	case FourOfKind:
		return "FourOfKind"
	case FiveOfKind:
		return "FiveOfKind"
	default:
		return fmt.Sprintf("Unknown HandStrength: %d", hs)
	}
}

func countDistinctValues(s string) int {
	distinctValues := 0
	var seen [5]bool
	for i := 0; i < len(s); i++ {
		if seen[i] {
			continue
		}
		seen[i] = true
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}
			if s[i] == s[j] {
				seen[j] = true
			}
		}
		distinctValues++
	}
	return distinctValues
}

func HandStrength1(hand string) HandStrength {
	switch countDistinctValues(hand) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		for i := 0; i < 3; i++ {
			count := 0
			for j := 0; j < len(hand); j++ {
				if hand[i] == hand[j] {
					count++
				}
			}
			if count == 2 {
				return TwoPair
			} else if count == 3 {
				return ThreeOfKind
			}
		}
	case 2:
		hasFour := false
		for i := 0; i < 2 && !hasFour; i++ {
			count := 0
			for j := 0; j < len(hand); j++ {
				if hand[i] == hand[j] {
					count++
				}
			}
			if count == 4 {
				hasFour = true
			}
		}
		if hasFour {
			return FourOfKind
		} else {
			return FullHouse
		}
	case 1:
		return FiveOfKind
	}

	panic("unexpected number of distinct cards!")
}

func HandStrength2(hand string) HandStrength {
	jokers := 0
	for _, c := range hand {
		if c == 'J' {
			jokers++
		}
	}
	switch countDistinctValues(hand) {
	case 5:
		if jokers == 1 {
			return OnePair
		}
		return HighCard
	case 4:
		if jokers > 0 {
			return ThreeOfKind
		}
		return OnePair
	case 3:
		for i := 0; i < 3; i++ {
			count := 0
			for j := 0; j < len(hand); j++ {
				if hand[i] == hand[j] {
					count++
				}
			}
			if count == 2 {
				if jokers == 1 {
					return FullHouse
				}
				if jokers == 2 {
					return FourOfKind
				}
				return TwoPair
			} else if count == 3 {
				if jokers > 0 {
					return FourOfKind
				}
				return ThreeOfKind
			}
		}
	case 2:
		hasFour := false
		for i := 0; i < 2 && !hasFour; i++ {
			count := 0
			for j := 0; j < len(hand); j++ {
				if hand[i] == hand[j] {
					count++
				}
			}
			if count == 4 {
				hasFour = true
			}
		}
		if hasFour {
			if jokers > 0 {
				return FiveOfKind
			}
			return FourOfKind
		} else {
			if jokers > 0 {
				return FiveOfKind
			}
			return FullHouse
		}
	case 1:
		return FiveOfKind
	}

	panic("unexpected number of distinct cards!")
}

func CardToNum(card byte, useJoker bool) int {
	if '2' <= card && card <= '9' {
		return int(card - '0')
	}
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if useJoker {
			return 1
		}
		return 11
	case 'T':
		return 10
	}
	panic("unexpected card value")
}

func getGameCmp(ver int) func(Game, Game) int {
	var handStrength func(string) HandStrength
	var useJoker bool
	if ver == 1 {
		handStrength = HandStrength1
		useJoker = false
	} else if ver == 2 {
		handStrength = HandStrength2
		useJoker = true
	} else {
		panic("unexpected ver")
	}

	return func(a Game, b Game) int {
		aStrength := handStrength(a.Hand)
		bStrength := handStrength(b.Hand)

		if aStrength == bStrength {
			for i := 0; i < len(a.Hand); i++ {
				cardA := CardToNum(a.Hand[i], useJoker)
				cardB := CardToNum(b.Hand[i], useJoker)
				if cardA < cardB {
					return -1
				} else if cardA > cardB {
					return 1
				}
			}
			return 0
		}
		if aStrength < bStrength {
			return -1
		}
		return 1
	}
}

func GetInput() []Game {
	bytes, _ := os.ReadFile("input.txt")
	text := string(bytes)
	lines := strings.Split(text, "\n")
	var games []Game
	for _, line := range lines {
		if line == "" {
			continue
		}
		hand, bidStr, _ := strings.Cut(line, " ")
		bid, _ := strconv.Atoi(bidStr)
		newGame := Game{
			Hand: hand,
			Bid:  bid,
		}
		games = append(games, newGame)
	}
	return games
}

func SumBids(games []Game) int {
	result := 0
	for i, game := range games {
		result += (i + 1) * game.Bid
	}
	return result
}

func Part1(games []Game) int {
	slices.SortFunc(games, getGameCmp(1))
	return SumBids(games)
}

func Part2(games []Game) int {
	slices.SortFunc(games, getGameCmp(2))
	return SumBids(games)
}

func main() {
	games := GetInput()
	fmt.Println("part 1:", Part1(games), "(expected 250453939)")
	fmt.Println("part 2:", Part2(games), "(expected 248652697)")
}
