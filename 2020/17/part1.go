package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Point struct{ X, Y, Z int }
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
			space.Add(Point{X: col, Y: row, Z: 0})
		}
		col++
	}
	return space
}

func (space Space) Add(point Point) {
	space[point] = struct{}{}
}

func (space Space) String() string {
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	var layers [][]Point
	for p := range space {
		minX, maxX = min(minX, p.X), max(maxX, p.X)
		minY, maxY = min(minY, p.Y), max(maxY, p.Y)
		for layerIdx, layer := range layers {
			if len(layer) > 0 && (layer)[0].Z == p.Z {
				layers[layerIdx] = append(layer, p)
				goto nextPoint
			}
		}
		layers = append(layers, []Point{p})
	nextPoint:
	}

	sort.Slice(layers, func(i, j int) bool {
		return layers[i][0].Z < layers[j][0].Z
	})
	for layerIdx := range layers {
		sort.Slice(layers[layerIdx], func(i, j int) bool {
			iy, jy := layers[layerIdx][i].Y, layers[layerIdx][j].Y
			if iy < jy {
				return true
			}
			if iy == jy {
				ix, jx := layers[layerIdx][i].X, layers[layerIdx][j].X
				return ix < jx
			}
			return false
		})
	}

	sb := strings.Builder{}
	for layerIdx, layer := range layers {
		if layerIdx != 0 {
			sb.WriteByte('\n')
		}
		header := fmt.Sprintf("z=%v\n", layer[0].Z)
		sb.WriteString(header)
		pointIdx := 0
		for row := minY; row <= maxY; row++ {
			for col := minX; col <= maxX; col++ {
				if pointIdx >= len(layer) || row < layer[pointIdx].Y || col < layer[pointIdx].X {
					sb.WriteByte('.')
				} else if col == layer[pointIdx].X {

					sb.WriteByte('#')
					pointIdx++
				} else {
					panic("aaaaaa")
				}
			}
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func (space Space) GetNeighbours(point Point) (active []Point, inactive []Point) {
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if z == 0 && y == 0 && x == 0 {
					continue
				}
				p := Point{
					X: point.X + x,
					Y: point.Y + y,
					Z: point.Z + z,
				}
				if _, exists := space[p]; exists {
					active = append(active, p)
				} else {
					inactive = append(inactive, p)
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
	fmt.Println("[1]", len(space))
}
