package main

import (
    "slices"
	"fmt"
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
    sign := 1
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Int {
			n, err := strconv.Atoi(s.TokenText())
			if err != nil {
				panic(err)
			}
            n *= sign
            sign = 1
			result = append(result, n)
		} else if tok == '-' {
            sign = -1
        }
	}
	return result
}

type History []int

func GetInput() (histories []History) {
    bytes, _ := os.ReadFile("input.txt")
    lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
    for _, line := range lines {
        histories = append(histories, ParseNums(line))
    }
    return
}

func Predict(history History) int {
    allZero := true
    for _, x := range history {
        if x != 0 {
            allZero = false
            break
        }
    }
    if allZero {
        return 0
    }

    var newHist History
    for i := 1; i < len(history); i++ {
        newHist = append(newHist, history[i] - history[i-1])
    }

    return history[len(history) - 1] + Predict(newHist)
}

func Part1(histories []History) (sum int) {
    for _, hist := range histories {
        sum += Predict(hist)
    }
    return
}

func Part2(histories []History) (sum int) {
    for _, hist := range histories {
        slices.Reverse(hist)
        sum += Predict(hist)
    }
    return
}

func main() {
    histories := GetInput()
    fmt.Println("part 1:", Part1(histories), "(expected 1819125966)")
    fmt.Println("part 2:", Part2(histories), "(expected ?)")
}
