package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/scanner"
)

type Almanac struct {
	Seeds []uint64
	Maps  [7]AlmanacMap
}

func (a Almanac) String() string {
	sb := strings.Builder{}
	sb.WriteString("seeds:")
	for _, num := range a.Seeds {
		sb.WriteByte(' ')
		sb.WriteString(strconv.FormatUint(num, 10))
	}
	for _, almanacMap := range a.Maps {
		sb.WriteByte('\n')
		sb.WriteByte('\n')
		sb.WriteString(fmt.Sprint(almanacMap))
	}
	return sb.String()
}

type AlmanacMap struct {
	Name    string
	Entries []AlmanacMapEntry
}

func (am AlmanacMap) Get(x uint64) uint64 {
	for _, entry := range am.Entries {
		if x >= entry.Start && x < entry.Start+entry.Len {
			return entry.Dest + x - entry.Start
		}
	}
	return x
}

func (am AlmanacMap) String() string {
	sb := strings.Builder{}
	sb.WriteString(am.Name)
	sb.WriteString(" map:")
	for _, entry := range am.Entries {
		sb.WriteByte('\n')
		sb.WriteString(fmt.Sprint(entry))
	}
	return sb.String()
}

type AlmanacMapEntry struct {
	Dest  uint64
	Start uint64
	Len   uint64
}

func (ame AlmanacMapEntry) String() string {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(ame.Dest, 10))
	sb.WriteByte(' ')
	sb.WriteString(strconv.FormatUint(ame.Start, 10))
	sb.WriteByte(' ')
	sb.WriteString(strconv.FormatUint(ame.Len, 10))
	return sb.String()
}

func ParseNums(str string) []uint64 {
	var result []uint64
	var s scanner.Scanner
	s.Init(strings.NewReader(str))
	s.Mode = scanner.ScanInts
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Int {
			n, err := strconv.ParseUint(s.TokenText(), 10, 64)
			if err != nil {
				panic(err)
			}
			result = append(result, n)
		}
	}
	return result
}

func GetInput() Almanac {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	input = strings.ReplaceAll(input, "\r\n", "\n")

	lines := bufio.NewScanner(strings.NewReader(input))
	lines.Split(bufio.ScanLines)

	var almanac Almanac

	lines.Scan()
	almanac.Seeds = ParseNums(lines.Text())
	lines.Scan()

	parseAlmanacMap := func() (almanacMap AlmanacMap) {
		lines.Scan()
		almanacMap.Name, _, _ = strings.Cut(lines.Text(), " ")
		for lines.Scan() {
			line := lines.Text()
			if line == "" {
				break
			}
			nums := ParseNums(line)
			newEntry := AlmanacMapEntry{
				Dest:  nums[0],
				Start: nums[1],
				Len:   nums[2],
			}
			almanacMap.Entries = append(almanacMap.Entries, newEntry)
		}
		slices.SortFunc(
			almanacMap.Entries,
			func(a, b AlmanacMapEntry) int {
				if a.Start > b.Start {
					return 1
				}
				if a.Start == b.Start {
					return 0
				}
				return -1
			},
		)

		return
	}

	for i := 0; i < 7; i++ {
		almanac.Maps[i] = parseAlmanacMap()
	}

	return almanac
}

func Part1(a Almanac) uint64 {
	var minLoc uint64 = math.MaxUint64

	var walk func(uint64, int)
	walk = func(value uint64, mapIdx int) {
		if mapIdx == len(a.Maps) {
			minLoc = min(minLoc, value)
		} else {
			new := a.Maps[mapIdx].Get(value)
			walk(new, mapIdx+1)
		}
	}

	for _, seed := range a.Seeds {
		walk(seed, 0)
	}

	return minLoc
}

func Part2(a Almanac) uint64 {
	var minLoc uint64 = math.MaxUint64

	var walk func(uint64, int)
	walk = func(value uint64, mapIdx int) {
		if mapIdx == len(a.Maps) {
			minLoc = min(minLoc, value)
		} else {
			new := a.Maps[mapIdx].Get(value)
			walk(new, mapIdx+1)
		}
	}

	for i := 0; i < len(a.Seeds)-1; i += 2 {
		fmt.Println(i/2+1, "/", len(a.Seeds)/2)
		var j uint64
		for j = 0; j < a.Seeds[i+1]; j++ {
			walk(a.Seeds[i]+j, 0)
		}
	}

	return minLoc
}

func main() {
	almanac := GetInput()
	fmt.Println("part 1:", Part1(almanac), "(expected 196167384)")
	fmt.Println("part 2:", Part2(almanac), "(expected 125742456)")
}
