package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Engine struct {
	Schema []byte
	Rows   int
	Cols   int
}

func (e Engine) String() string {
	sb := strings.Builder{}
	for r := 0; r < e.Rows; r++ {
		beg, end := r*e.Cols, (r+1)*e.Cols
		newSlice := e.Schema[beg:end]
		sb.Write(newSlice)
		if r < e.Rows-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func (e Engine) Neighbors(i int) []int {
	ns := make([]int, 0)

	left := i%e.Cols != 0
	right := (i+1)%e.Cols != 0
	down := i >= e.Cols
	up := i < (e.Rows-1)*e.Cols

	if left {
		ns = append(ns, i-1)
		if up {
			ns = append(ns, i-1+e.Cols)
		}
		if down {
			ns = append(ns, i-1-e.Cols)
		}
	}
	if right {
		ns = append(ns, i+1)
		if up {
			ns = append(ns, i+1+e.Cols)
		}
		if down {
			ns = append(ns, i+1-e.Cols)
		}
	}
	if up {
		ns = append(ns, i+e.Cols)
	}
	if down {
		ns = append(ns, i-e.Cols)
	}

	return ns
}

func (e Engine) LastInLine(i int) bool {
	return (i+1)%e.Cols == 0
}

func isSymbol(char byte) bool {
	symbols := []byte{
		'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '\\', '+', '/',
	}
	return slices.Contains(symbols, char)
}

func isGear(char byte) bool {
	return char == '*'
}

type Part struct {
	Num        int
	Len        int
	Pos        int
	SymbolsPos []int
}

func (e Engine) Parts(symbolFilter func(byte) bool) []Part {
	parts := make([]Part, 0)
	curr, currLen := 0, 0
	symbols := make([]int, 0)

	addCurr := func(pos int) {
		if curr > 0 && len(symbols) > 0 {
			newPart := Part{
				Num:        curr,
				Len:        currLen,
				Pos:        pos,
				SymbolsPos: symbols,
			}
			parts = append(parts, newPart)
		}
		curr, currLen = 0, 0
		symbols = make([]int, 0)
	}

	findNeighboringSymbols := func(i int) {
		ns := e.Neighbors(i)
		for _, n := range ns {
			if symbolFilter(e.Schema[n]) {
				alreadyAdded := false
				for i := 0; !alreadyAdded && i < len(symbols); i++ {
					alreadyAdded = symbols[i] == n
				}
				if !alreadyAdded {
					symbols = append(symbols, n)
				}
			}
		}
	}

	for i, c := range e.Schema {
		if '0' <= c && c <= '9' {
			curr = (curr * 10) + int(c-'0')
			currLen++
			findNeighboringSymbols(i)
			if e.LastInLine(i) {
				addCurr(i)
			}
		} else {
			addCurr(i)
		}
	}
	addCurr(len(e.Schema) - 1)

	return parts
}

func GetInput() Engine {
	raw_schema, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	schema := string(raw_schema)
	schema = strings.TrimSpace(schema)
	rows := strings.Count(schema, "\n") + 1
	cols := strings.Index(schema, "\n")
	schema = strings.Join(strings.Split(schema, "\n"), "")
	return Engine{
		Schema: []byte(schema),
		Rows:   rows,
		Cols:   cols,
	}
}

func Part1(e Engine) int {
	sum := 0
	for _, part := range e.Parts(isSymbol) {
		sum += part.Num
	}
	return sum
}

func Part2(e Engine) uint64 {
	haveCommonSymbols := func(p1, p2 Part) bool {
		for _, s1 := range p1.SymbolsPos {
			for _, s2 := range p2.SymbolsPos {
				if s1 == s2 {
					return true
				}
			}
		}
		return false
	}

	partsWithGears := e.Parts(isGear)
	processed := make([]bool, len(partsWithGears))
	var sum uint64

	for i := 0; i < len(partsWithGears)-1; i++ {
		if processed[i] {
			continue
		}
		processed[i] = true
		curr := partsWithGears[i]
		friendsCount := 0
		var friend Part
		for j := i + 1; j < len(partsWithGears); j++ {
			other := partsWithGears[j]
			if !processed[j] && haveCommonSymbols(curr, other) {
				processed[j] = true
				friendsCount++
				friend = other
			}
		}
		if friendsCount == 1 {
			sum += uint64(curr.Num * friend.Num)
		}
	}

	return sum
}

func main() {
	engine := GetInput()
	fmt.Println("part 1:", Part1(engine), "(expected 556367)")
	fmt.Println("part 2:", Part2(engine), "(expected 89471771)")
}
