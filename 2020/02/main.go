package main

import (
	"bufio"
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

func newPasswordFromInput(text string) Password {
	fields := strings.Fields(text)
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
	file, _ := os.Open("input.txt")
	defer file.Close()
	var passwords []Password
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwords = append(passwords, newPasswordFromInput(scanner.Text()))
	}
	valid1, valid2 := 0, 0
	for _, p := range passwords {
		if p.validate1() {
			valid1++
		}
		if p.validate2() {
			valid2++
		}
	}
	fmt.Printf("[1] valid passwords: %v\n", valid1)
	fmt.Printf("[2] valid passwords: %v\n", valid2)
}
