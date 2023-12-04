package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
    WinningNums []int
    DrawnNums []int
}

func GetInput() []Card {
    raw_text, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    text := string(raw_text)

    parseNums := func(s string) []int {
        nums := make([]int, 0)
        for _, numStr := range strings.Split(s, " ") {
            if numStr == "" {
                continue
            }
            num, err := strconv.Atoi(numStr)
            if err != nil {
                panic(err)
            }
            nums = append(nums, num)
        }
        return nums
    }

    lines := strings.Split(text, "\n")
    cards := make([]Card, 0)
    for _, line := range lines {
        if line == "" {
            continue
        }
        _, nums, _ := strings.Cut(line, ": ")
        winningNums, drawnNums, _ := strings.Cut(nums, " | ")
        newCard := Card {
            WinningNums: parseNums(winningNums),
            DrawnNums: parseNums(drawnNums),
        }
        cards = append(cards, newCard)
    }

    return cards
}

func Part1(cards []Card) int {
    sum := 0

    for _, card := range cards {
        points := 0
        for _, num := range card.DrawnNums {
            if slices.Contains(card.WinningNums, num) {
                if points == 0 {
                    points = 1
                } else {
                    points *= 2
                }
            }
        }
        sum += points
    }

    return sum
}

type stack[T any] struct {
    slice []T
}

func Stack[T any]() stack[T] {
    return stack[T] {
        slice: make([]T, 0),
    }
}

func (s *stack[T]) Push(el T) {
    s.slice = append(s.slice, el)
}

func (s *stack[T]) Pop() T {
    last := s.slice[len(s.slice)-1]
    s.slice = s.slice[:len(s.slice)-1]
    return last
}

func (s stack[T]) Len() int {
    return len(s.slice)
}

func Part2(cards []Card) int {
    matches := make([]int, len(cards))
    for cardIdx, card := range cards {
        for _, num := range card.DrawnNums {
            if slices.Contains(card.WinningNums, num) {
                matches[cardIdx]++
            }
        }
    }

    sum := 0
    stack := Stack[int]()

    for i := 0; i < len(cards); i++ {
        stack.Push(i)
    }

    for stack.Len() > 0 {
        curr := stack.Pop()
        sum += 1
        if curr > len(cards) {
            continue
        }
        m := matches[curr]
        for i := 1; i <= m; i++ {
            stack.Push(curr + i)
        }
    }

    return sum
}

func main() {
    cards := GetInput()
    fmt.Println("part 1:", Part1(cards), "(expected: 27454)")
    fmt.Println("part 2:", Part2(cards), "(expected: 6857330)")
}
