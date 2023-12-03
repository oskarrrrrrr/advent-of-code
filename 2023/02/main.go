package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

type Game []CubeSet

func GetInput() []Game {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	s := string(b)
	games := make([]Game, 0)
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		newGame := make(Game, 0)
		line = line[strings.Index(line, ":")+1:]
		sets := strings.Split(line, ";")
		for _, set := range sets {
			newSet := CubeSet{}
			colors := strings.Split(strings.TrimSpace(set), ",")
			for _, color := range colors {
				color := strings.TrimSpace(color)
				split := strings.Split(color, " ")
				num, c := split[0], split[1]
				n, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				switch c {
				case "red":
					newSet.Red = n
				case "green":
					newSet.Green = n
				case "blue":
					newSet.Blue = n
				}
			}
			newGame = append(newGame, newSet)
		}
		games = append(games, newGame)
	}
	return games
}

func Part1(games []Game) int {
	maxRed, maxGreen, maxBlue := 12, 13, 14
	sum := 0

	checkGame := func(game Game) bool {
		for _, set := range game {
			if set.Red > maxRed || set.Green > maxGreen || set.Blue > maxBlue {
				return false
			}
		}
		return true
	}

	for i, game := range games {
		if checkGame(game) {
			sum += i + 1
		}
	}
	return sum
}

func Part2(games []Game) uint64 {
	calcMinSet := func(game Game) CubeSet {
		minSet := CubeSet{}
		for _, set := range game {
			minSet.Blue = max(minSet.Blue, set.Blue)
			minSet.Red = max(minSet.Red, set.Red)
			minSet.Green = max(minSet.Green, set.Green)
		}
		return minSet
	}
	var sum uint64 = 0
	for _, game := range games {
		minSet := calcMinSet(game)
		sum += uint64(minSet.Red) * uint64(minSet.Green) * uint64(minSet.Blue)
	}
	return sum
}

func main() {
	input := GetInput()
	fmt.Println("part 1:", Part1(input), "(expected: 1853)")
	fmt.Println("part 2:", Part2(input), "(expected: 72706)")
}
