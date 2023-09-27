package main

import (
	"fmt"
	"os"
	"strings"
)

type Trees []string

func (t Trees) At(row, col int) bool {
	col = col % len(t[0])
	return t[row][col] == '#'
}

func (t Trees) CountTreesOnSlope(right, down int) int {
	tcount := 0
	row, col := 0, 0
	for row < len(t) {
		if t.At(row, col) {
			tcount++
		}
		row += down
		col += right
	}
	return tcount
}

func ReadInput() Trees {
	inputBytes, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(inputBytes))
	var trees Trees
	for _, line := range strings.Split(input, "\n") {
		trees = append(trees, line)
	}
	return trees
}

func Part1(trees Trees) int {
	return trees.CountTreesOnSlope(3, 1)
}

func Part2(trees Trees) int {
	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	mult := 1
	for _, slope := range slopes {
		mult *= trees.CountTreesOnSlope(slope[0], slope[1])
	}
	return mult
}

func main() {
	trees := ReadInput()
	fmt.Println("[1]", Part1(trees))
	fmt.Println("[2]", Part2(trees))
}
