package main

import (
	"fmt"
	"math"
	"slices"
)

type Point struct{ x, y, z int }
type Space map[Point]any

func (s Space) String() string {
	minX, maxX, minY, maxY := math.MinInt, math.MaxInt, math.MinInt, math.MaxInt
	var zs []int
	for p := range s {
		zs = append(zs, p.z)
	}
	slices.Sort(zs)

	for _, z := range zs {
		fmt.Printf("z=%v\n", z)

	}
}

func main() {

}
