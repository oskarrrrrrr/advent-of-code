package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func ParseNums(str string) []int {
	var result []int
	var s scanner.Scanner
	s.Init(strings.NewReader(str))
	s.Mode = scanner.ScanInts
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Int {
			n, err := strconv.Atoi(s.TokenText())
			if err != nil {
				panic(err)
			}
			result = append(result, n)
		}
	}
	return result
}

type Race struct {
	Time     int
	Distance int
}

func GetInput() []Race {
	bytes, _ := os.ReadFile("input.txt")
	text := strings.TrimSpace(string(bytes))
	timesStr, distancesStr, _ := strings.Cut(text, "\n")
	times, distances := ParseNums(timesStr), ParseNums(distancesStr)
	var races []Race
	for i := 0; i < len(times); i++ {
		races = append(
			races,
			Race{
				Time:     times[i],
				Distance: distances[i],
			},
		)
	}
	return races
}

func Part1(races []Race) int {
	result := 1
	for _, race := range races {
		winCount := 0
		for k := 1; k < race.Time; k++ {
			if k*(race.Time-k) > race.Distance {
				winCount++
			}
		}
		result *= winCount
	}
	return result
}

func Part2(races []Race) int {
	var time, dist float64
	for _, race := range races {
		timeLen := len(strconv.Itoa(race.Time))
		time = time*math.Pow10(timeLen) + float64(race.Time)
		distLen := len(strconv.Itoa(race.Distance))
		dist = dist*math.Pow10(distLen) + float64(race.Distance)
	}

	delta := time*time - 4*dist
	deltaSqrt := math.Sqrt(delta)
	k1 := (time + deltaSqrt) / 2
	k2 := (time - deltaSqrt) / 2

	return int(math.Floor(k1)-math.Ceil(k2)) + 1
}

func main() {
	races := GetInput()
	fmt.Println(races)
	fmt.Println("part 1:", Part1(races), "(expected 500346)")
	fmt.Println("part 2:", Part2(races), "(expected 42515755)")
}
