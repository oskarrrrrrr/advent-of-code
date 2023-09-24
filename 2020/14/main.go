package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	pos   uint64
	value uint64
}

func (a Assignment) String() string {
	return fmt.Sprintf("Assignment(pos=%v, value=%v)", a.pos, a.value)
}

type Mask struct {
	ones   uint64
	zeroes uint64
	xs     uint64
}

func (m Mask) String() string {
	return fmt.Sprintf(
		"Mask(ones=%s, zeroes=%s, xs=%s)",
		strconv.FormatUint(m.ones, 2),
		strconv.FormatUint(m.zeroes, 2),
		strconv.FormatUint(m.xs, 2),
	)
}

func ParseAssignment(s string) *Assignment {
	var pos, value uint64
	cb_pos := strings.Index(s, "]")
	_pos, _ := strconv.Atoi(s[4:cb_pos])
	pos = uint64(_pos)
	_value, _ := strconv.Atoi(s[cb_pos+4:])
	value = uint64(_value)
	return &Assignment{
		pos:   pos,
		value: value,
	}
}

func ParseMask(s string) *Mask {
	eq_idx := strings.Index(s, "=")
	s = strings.TrimSpace(s[eq_idx+1:])
	var ones, zeroes, xs uint64 = 0, math.MaxUint64, 0
	for i, m := range s {
		if m == '1' {
			ones |= 1 << (len(s) - i - 1)
		} else if m == '0' {
			zeroes &= math.MaxUint64 ^ (1 << (len(s) - i - 1))
		} else if m == 'X' {
			xs |= 1 << (len(s) - i - 1)
		} else {
			panic("Unexpected char!")
		}
	}
	return &Mask{
		ones:   ones,
		zeroes: zeroes,
		xs:     xs,
	}
}

func ParseInput() []interface{} {
	file, _ := os.Open("input.txt")
	defer file.Close()
	var input []interface{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			input = append(input, ParseMask(line))
		} else if strings.HasPrefix(line, "mem") {
			input = append(input, ParseAssignment(line))
		} else {
			panic("Unexpected line format!")
		}
	}
	return input
}

type Mem map[uint64]uint64
type OnAssigmentAction func(Mem, *Mask, *Assignment)

func Solve(input []interface{}, onAssigment OnAssigmentAction) uint64 {
	var mask *Mask
	mem := make(Mem)
	for _, stmt := range input {
		switch s := stmt.(type) {
		case *Mask:
			mask = s
		case *Assignment:
			onAssigment(mem, mask, s)
		default:
			panic("Ooops unexpected type!")
		}
	}
	var sum uint64
	for _, x := range mem {
		sum += x
	}
	return sum
}

func (m Mask) P1Apply(v uint64) uint64 {
	return (v & m.zeroes) | m.ones
}

func P1OnAssignment(mem Mem, mask *Mask, a *Assignment) {
	mem[a.pos-1] = mask.P1Apply(a.value)
}

func (m Mask) P2Apply(v uint64, ord uint64) (uint64, bool) {
	v = v | m.ones
	all_ones := true
	for xsI := 0; xsI < 64; xsI++ {
		if m.xs&(1<<xsI) > 0 {
			if ord&1 == 1 {
				v |= (1 << xsI)
			} else {
				v &= math.MaxUint64 ^ (1 << xsI)
				all_ones = false
			}
			ord >>= 1
		}
	}
	return v, all_ones
}

func P2OnAssignment(mem Mem, mask *Mask, a *Assignment) {
	var done bool
	var ord uint64
	for !done {
		var memIdx uint64
		memIdx, done = mask.P2Apply(a.pos, ord)
		mem[memIdx] = a.value
		ord++
	}
}

func main() {
	input := ParseInput()
	fmt.Println("[1]", Solve(input, P1OnAssignment))
	fmt.Println("[2]", Solve(input, P2OnAssignment))
}
