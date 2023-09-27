package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Constraint struct {
	letter       byte
	lower, upper int
}

type Password struct {
	text       string
	constraint Constraint
}

func ParsePassword(s string) Password {
	fields := strings.Fields(s)
	if len(fields) != 3 {
		panic("Unexpected input!")
	}
	constraint_bounds := strings.Split(fields[0], "-")
	lower, _ := strconv.Atoi(constraint_bounds[0])
	upper, _ := strconv.Atoi(constraint_bounds[1])
	return Password{
		text: fields[2],
		constraint: Constraint{
			letter: fields[1][0],
			lower:  lower,
			upper:  upper,
		},
	}
}

func ReadInput() []Password {
	inputBytes, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(inputBytes))
	var passwords []Password
	for _, line := range strings.Split(input, "\n") {
		passwords = append(passwords, ParsePassword(line))
	}
	return passwords
}

func (p Password) validate1() bool {
	c := 0
	for i := 0; i < len(p.text); i++ {
		if p.constraint.letter == p.text[i] {
			c++
		}
	}
	return p.constraint.lower <= c && c <= p.constraint.upper
}

func (p Password) validate2() bool {
	left := p.text[p.constraint.lower-1] == p.constraint.letter
	right := p.text[p.constraint.upper-1] == p.constraint.letter
	return left != right
}

func main() {
	passwords := ReadInput()
	valid1, valid2 := 0, 0
	for _, p := range passwords {
		if p.validate1() {
			valid1++
		}
		if p.validate2() {
			valid2++
		}
	}
	fmt.Println("[1]", valid1)
	fmt.Println("[2]", valid2)
}
