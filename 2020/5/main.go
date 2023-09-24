package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	// "os"
	// "bufio"
)

const (
	LOWER_SELECTOR = iota
	UPPER_SELECTOR = iota
)

type BoardingPassCode struct {
	row_selectors    [7]int
	column_selectors [3]int
}

func newBoardingPassCode(s string) BoardingPassCode {
	var row [7]int
	for i := 0; i < 7; i++ {
		if s[i] == 'F' {
			row[i] = LOWER_SELECTOR
		} else if s[i] == 'B' {
			row[i] = UPPER_SELECTOR
		} else {
			fmt.Println("Unexpected selector:", s[i])
			os.Exit(1)
		}
	}
	var column [3]int
	for i := 7; i < len(s); i++ {
		if s[i] == 'L' {
			column[i-7] = LOWER_SELECTOR
		} else if s[i] == 'R' {
			column[i-7] = UPPER_SELECTOR
		} else {
			fmt.Println("Unexpected selector:", s[i])
			os.Exit(1)
		}
	}
	return BoardingPassCode{row, column}
}

func (bpc BoardingPassCode) GetRow() int {
	l, r := 0, 128
	for _, s := range bpc.row_selectors {
		pivot := l + ((r - l) / 2)
		if s == LOWER_SELECTOR {
			r = pivot
		} else if s == UPPER_SELECTOR {
			l = pivot
		} else {
			panic("whaaat")
		}
	}
	return l
}

func (bpc BoardingPassCode) GetColumn() int {
	l, r := 0, 8
	for _, s := range bpc.column_selectors {
		pivot := l + ((r - l) / 2)
		if s == LOWER_SELECTOR {
			r = pivot
		} else if s == UPPER_SELECTOR {
			l = pivot
		} else {
			panic("whaaat")
		}
	}
	return l
}

func (bpc BoardingPassCode) GetId() int {
	return (bpc.GetRow() * 8) + bpc.GetColumn()
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	var bp_ids []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bpc := newBoardingPassCode(scanner.Text())
		bp_ids = append(bp_ids, bpc.GetId())
	}
	sort.Ints(bp_ids)
	fmt.Println("[1] max id found:", bp_ids[len(bp_ids)-1])

	for i := 0; i < len(bp_ids)-1; i++ {
		if bp_ids[i]+2 == bp_ids[i+1] {
			fmt.Println("[2] missing seat:", bp_ids[i]+1)
		}
	}
}
