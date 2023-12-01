package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func part1(input string) uint64 {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var sum uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		firstFound := false
		var prev uint64 = 0
		for _, c := range line {
			if '0' <= c && c <= '9' {
				prev = uint64(c - '0')
				if !firstFound {
					sum += prev * 10
					firstFound = true
				}
			}
		}
		sum += prev
	}
	return sum
}

var DIGITS = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func part2(input string) uint64 {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var sum uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		firstFound := false
		var prev uint64 = 0
		for i, c := range line {
			if '0' <= c && c <= '9' {
				prev = uint64(c - '0')
				if !firstFound {
					sum += prev * 10
					firstFound = true
				}
			} else {
				for d, ds := range DIGITS {
					if strings.HasPrefix(line[i:], ds) {
						prev = uint64(d + 1)
						if !firstFound {
							sum += prev * 10
							firstFound = true
						}
					}
				}
			}
		}
		sum += prev
	}
	return sum
}

func main() {
	input := getInput()
	fmt.Println("part1:", part1(input), "(expected: 56042)")
	fmt.Println("part2:", part2(input), "(expected: 55358)")
}
