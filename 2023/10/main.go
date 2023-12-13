package main

import (
	"fmt"
	"os"
	"strings"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (item T, ok bool) {
	if len(q.items) == 0 {
		return
	}
	ok = true
	item = q.items[0]
	q.items = q.items[1:]
	return
}

func (q Queue[T]) Size() int {
	return len(q.items)
}

func GetInput() []string {
	bytes, _ := os.ReadFile("input.txt")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

type Point struct{ Row, Col int }

func getDist(field []string) (dist map[Point]int) {
	var s Point
	for rowIdx, row := range field {
		colIdx := strings.Index(row, "S")
		if colIdx != -1 {
			s.Row = rowIdx
			s.Col = colIdx
			break
		}
	}

	height, width := len(field), len(field[0])
	neighbors := func(p Point) (ns []Point) {
		curr := field[p.Row][p.Col]
		if p.Row > 0 {
			next := field[p.Row-1][p.Col]
			if (curr == 'S' || curr == '|' || curr == 'J' || curr == 'L') &&
				(next == '|' || next == '7' || next == 'F') {
				ns = append(ns, Point{p.Row - 1, p.Col})
			}
		}
		if p.Row < height-1 {
			next := field[p.Row+1][p.Col]
			if (curr == 'S' || curr == '|' || curr == '7' || curr == 'F') &&
				(next == '|' || next == 'J' || next == 'L') {
				ns = append(ns, Point{p.Row + 1, p.Col})
			}
		}
		if p.Col > 0 {
			next := field[p.Row][p.Col-1]
			if (curr == 'S' || curr == '-' || curr == 'J' || curr == '7') &&
				(next == '-' || next == 'F' || next == 'L') {
				ns = append(ns, Point{p.Row, p.Col - 1})
			}
		}
		if p.Col < width-1 {
			next := field[p.Row][p.Col+1]
			if (curr == 'S' || curr == '-' || curr == 'F' || curr == 'L') &&
				(next == '-' || next == 'J' || next == '7') {
				ns = append(ns, Point{p.Row, p.Col + 1})
			}
		}
		return
	}

	var _max int
	dist = make(map[Point]int)
	dist[s] = 0
	var todo Queue[Point]
	todo.Enqueue(s)
	for todo.Size() > 0 {
		curr, _ := todo.Dequeue()
		currDist, ok := dist[curr]
		if !ok {
			panic("no distance for curr")
		}
		for _, n := range neighbors(curr) {
			nDist, seen := dist[n]
			if seen {
				dist[n] = min(nDist, currDist+1)
			} else {
				dist[n] = currDist + 1
				todo.Enqueue(n)
			}
			_max = max(_max, dist[n])
		}
	}
	return
}

func printDist(dist map[Point]int, height, width int) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if d, seen := dist[Point{row, col}]; seen {
				fmt.Printf("[%2d] ", d)
			} else {
				fmt.Print("[  ] ")
			}
		}
		fmt.Println()
	}
}

func Part1(dist map[Point]int) (_max int) {
	for _, v := range dist {
		_max = max(_max, v)
	}
	return
}

func Part2(dist map[Point]int, field []string) int {
	insideTilesCount := 0
	for rowIdx, row := range field {
		inside := false
		for colIdx, tile := range row {
			if _, inDist := dist[Point{rowIdx, colIdx}]; inDist {
				// check manually that S works as J
				if tile == '|' || tile == 'S' || tile == 'J' || tile == 'L' {
					inside = !inside
				}
			} else {
				if inside {
					insideTilesCount += 1
				}
			}
		}
	}

	return insideTilesCount
}

func main() {
	field := GetInput()
	dist := getDist(field)
	fmt.Println("part 1:", Part1(dist), "(expected 6757)")
	fmt.Println("part 2:", Part2(dist, field), "(expected 523)")
}
