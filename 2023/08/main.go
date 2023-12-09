package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Puzzle struct {
	Pattern      string
	Start, End   int
	Labels       []string
	Instructions []Instruction
}

type Instruction struct {
	Left, Right int
}

func ReadInput() (puzzle Puzzle) {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	puzzle.Pattern = lines[0]
	type InstructionLabels struct{ Left, Right string }
	var instructionsLabels []InstructionLabels
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		origin, leftRight, _ := strings.Cut(line, " = ")
		puzzle.Labels = append(puzzle.Labels, origin)
		instructionsLabels = append(
			instructionsLabels,
			InstructionLabels{Left: leftRight[1:4], Right: leftRight[6:9]},
		)
		puzzle.Start = slices.Index(puzzle.Labels, "AAA")
		puzzle.End = slices.Index(puzzle.Labels, "ZZZ")
	}
	for _, labels := range instructionsLabels {
		puzzle.Instructions = append(
			puzzle.Instructions,
			Instruction{
				Left:  slices.Index(puzzle.Labels, labels.Left),
				Right: slices.Index(puzzle.Labels, labels.Right),
			},
		)
	}
	return
}

func Part1(puzzle Puzzle) (steps int) {
	curr := puzzle.Start
	for curr != puzzle.End {
		for dirIdx := 0; dirIdx < len(puzzle.Pattern) && curr != puzzle.End; dirIdx++ {
			dir := puzzle.Pattern[dirIdx]
			if dir == 'L' {
				curr = puzzle.Instructions[curr].Left
			} else if dir == 'R' {
				curr = puzzle.Instructions[curr].Right
			} else {
				panic("Unexpected direction")
			}
			steps++
		}
	}
	return
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, rest ...int) int {
	result := a * b / GCD(a, b)
	for _, r := range rest {
		result = LCM(result, r)
	}
	return result
}

func Part2(puzzle Puzzle) int {
	var curr []int
	for i, label := range puzzle.Labels {
		if label[2] == 'A' {
			curr = append(curr, i)
		}
	}
	doneCount := 0
	pathLen := make([]int, len(curr))
	for doneCount < len(curr) {
		for dirIdx := 0; dirIdx < len(puzzle.Pattern) && doneCount < len(curr); dirIdx++ {
			dir := puzzle.Pattern[dirIdx]
			for i := 0; i < len(curr); i++ {
				if puzzle.Labels[curr[i]][2] == 'Z' {
					continue
				}
				if dir == 'L' {
					curr[i] = puzzle.Instructions[curr[i]].Left
				} else if dir == 'R' {
					curr[i] = puzzle.Instructions[curr[i]].Right
				} else {
					panic("Unexpected direction")
				}
				pathLen[i]++
				if puzzle.Labels[curr[i]][2] == 'Z' {
					doneCount++
				}
			}
		}
	}
	return LCM(pathLen[0], pathLen[1], pathLen[2:]...)
}

func main() {
	puzzle := ReadInput()
	fmt.Println("part 1:", Part1(puzzle), "(expected 21883)")
	fmt.Println("part 2:", Part2(puzzle), "(expected 12833235391111)")
}
