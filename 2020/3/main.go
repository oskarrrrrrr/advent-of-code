package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var trees Trees
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}

	tcount := trees.CountTreesOnSlope(3, 1)
	fmt.Println("[1] trees on path: ", tcount)

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
	fmt.Println("[2] mult result: ", mult)
}
