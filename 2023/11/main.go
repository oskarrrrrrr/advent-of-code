package main

import (
	"fmt"
	"os"
	"strings"
)

func GetInput() (image []string) {
    bytes, _ := os.ReadFile("input.txt")
    image = strings.Split(strings.TrimSpace(string(bytes)), "\n")
    return
}

func Sol(image []string, expansion int) int {
    extraCols := make([]int, len(image[0]))
    for col := 0; col < len(image[0]); col++ {
        allEmpty := true
        for _, row := range image {
            if row[col] != '.' {
                allEmpty = false
                break
            }
        }
        prev := 0
        if col != 0 {
            prev = extraCols[col-1]
        }
        if allEmpty {
            extraCols[col] = prev + expansion - 1
        } else {
            extraCols[col] = prev
        }
    }

    extraRows := make([]int, len(image))
    for rowIdx, row := range image {
        allEmpty := true
        for _, c := range row {
            if c != '.' {
                allEmpty = false
                break
            }
        }
        prev := 0
        if rowIdx != 0 {
            prev = extraRows[rowIdx-1]
        }
        if allEmpty {
            extraRows[rowIdx] = prev + expansion - 1
        } else {
            extraRows[rowIdx] = prev
        }

    }

    type Point struct { Row, Col int }
    var points []Point
    for rowIdx, row := range image {
        for colIdx, c := range row {
            if c == '#' {
                points = append(points, Point {rowIdx, colIdx})
            }
        }
    }

    abs := func(a int) int {
        if a < 0 { return -a }
        return a
    }

    sumDist := 0
    for p1Idx, p1 := range points {
        for p2Idx := p1Idx + 1; p2Idx < len(points); p2Idx++ {
            p2 := points[p2Idx]
            if p1Idx == p2Idx { continue }
            sumDist += abs(extraRows[p1.Row] - extraRows[p2.Row]) +
                abs(extraCols[p1.Col] - extraCols[p2.Col]) +
                abs(p1.Row - p2.Row) + abs(p1.Col - p2.Col)
        }
    }
    return sumDist
}

func Part1(image []string) int {
    return Sol(image, 2)
}

func Part2(image []string) int {
    return Sol(image, 1000000)
}

func main() {
    image := GetInput()
    fmt.Println("part 1:", Part1(image), "(expected 9693756)")
    fmt.Println("part 2:", Part2(image), "(expected 717878258016)")
}
