package main

import (
	"fmt"
	"os"
)

type Point struct{ X, Y, Z, W int }
type Space map[Point]any

func ParseSpace(s string) Space {
	space := make(Space)
	var row, col int
	for _, c := range s {
		if c == '\n' {
			col = 0
			row++
			continue
		}
		if c == '#' {
			space.Add(Point{X: col, Y: row, Z: 0, W: 0})
		}
		col++
	}
	return space
}

func (space Space) Add(point Point) {
	space[point] = struct{}{}
}

func (space Space) GetNeighbours(point Point) (active []Point, inactive []Point) {
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if w == 0 && z == 0 && y == 0 && x == 0 {
						continue
					}
					p := Point{
						X: point.X + x,
						Y: point.Y + y,
						Z: point.Z + z,
						W: point.W + w,
					}
					if _, exists := space[p]; exists {
						active = append(active, p)
					} else {
						inactive = append(inactive, p)
					}
				}
			}
		}
	}
	return
}

func (space Space) conwayIter() Space {
	newSpace := make(Space)
	incativeNeighboursCount := make(map[Point]int)
	for p := range space {
		active, inactive := space.GetNeighbours(p)
		if len(active) == 2 || len(active) == 3 {
			newSpace.Add(p)
		}
		for _, n := range inactive {
			incativeNeighboursCount[n]++
		}
	}
	for p := range incativeNeighboursCount {
		if incativeNeighboursCount[p] == 3 {
			newSpace.Add(p)
		}
	}
	return newSpace
}

func (space Space) Conway(iters int, verbose bool) Space {
	if verbose {
		fmt.Print("Before any cycles:\n\n")
		fmt.Println(space)
	}
	for i := 0; i < iters; i++ {
		space = space.conwayIter()
		if verbose {
			if i == 0 {
				fmt.Print("After 1 cycle:\n\n")
			} else {
				fmt.Printf("After %v cycles:\n\n", i+1)
			}
			fmt.Println(space)
		}
	}
	return space
}

func ReadInput() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	input := ReadInput()
	space := ParseSpace(input)
	space = space.Conway(6, false)
	fmt.Println("[2]", len(space))
}
