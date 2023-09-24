package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	lower, upper uint
}

func (r Range) in(x uint) bool {
	return r.lower <= x && x <= r.upper
}

type Field struct {
	name   string
	ranges []Range
}

func (f Field) ValidFieldValue(v uint) bool {
	for _, r := range f.ranges {
		if r.in(v) {
			return true
		}
	}
	return false
}

type Ticket struct {
	values []uint
}

func parseUint(s string) uint {
	value, _ := strconv.ParseUint(s, 10, 32)
	return uint(value)
}

func ParseRange(s string) Range {
	parts := strings.Split(s, "-")
	lower, upper := parseUint(parts[0]), parseUint(parts[1])
	return Range{lower: lower, upper: upper}
}

func ParseField(s string) Field {
	nameSepIdx := strings.Index(s, ":")
	name := s[:nameSepIdx]
	var ranges []Range
	for _, rangeStr := range strings.Fields(s[nameSepIdx+1:]) {
		if rangeStr == "or" {
			continue
		}
		ranges = append(ranges, ParseRange(rangeStr))
	}
	return Field{
		name:   name,
		ranges: ranges,
	}
}

func ParseTicket(s string) Ticket {
	var values []uint
	for _, n := range strings.Split(s, ",") {
		values = append(values, parseUint(n))
	}
	return Ticket{values}
}

func ParseInput() (fields []Field, yourTicket Ticket, nearbyTickets []Ticket) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var fieldsDone, yourTicketDone bool
	for scanner.Scan() {
		line := scanner.Text()
		if !fieldsDone {
			if line == "" {
				fieldsDone = true
				continue
			}
			fields = append(fields, ParseField(line))
		} else if !yourTicketDone {
			if line == "your ticket:" {
				continue
			}
			if line == "" {
				yourTicketDone = true
				continue
			}
			yourTicket = ParseTicket(line)
		} else {
			if line == "nearby tickets:" {
				continue
			}
			nearbyTickets = append(nearbyTickets, ParseTicket(line))
		}
	}
	return
}

func (t Ticket) IsValid(fields []Field) bool {
	for _, v := range t.values {
		for _, f := range fields {
			for _, r := range f.ranges {
				if r.in(v) {
					goto checkNext
				}
			}
		}
		return false
	checkNext:
	}
	return true
}

func Part1(fields []Field, tickets []Ticket) uint {
	var sum uint
	for _, ticket := range tickets {
		for _, v := range ticket.values {
			allInvalid := true
			for _, field := range fields {
				for _, r := range field.ranges {
					if r.in(v) {
						allInvalid = false
						goto endValueChecks
					}
				}
			}
		endValueChecks:
			if allInvalid {
				sum += v
			}
		}
	}
	return sum
}

func Part2(fields []Field, yourTicket Ticket, nearbyTickets []Ticket) uint {
    var validTickets []*Ticket
    for i := range nearbyTickets {
        if nearbyTickets[i].IsValid(fields) {
            validTickets = append(validTickets, &nearbyTickets[i])
        }
    }

	fieldsValues := make([][]uint, len(fields))
	for ti := range validTickets {
		for i := range fields {
			fieldsValues[i] = append(fieldsValues[i], validTickets[ti].values[i])
		}
	}
	for _, values := range fieldsValues {
		sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	}

	var determined []string
	candidates := make([][]string, len(fields))
	for fieldIdx, values := range fieldsValues {
		for _, field := range fields {
			allValid := true
			for _, v := range values {
				if !field.ValidFieldValue(v) {
					allValid = false
					goto endValueChecks
				}
			}
		endValueChecks:
			if allValid {
				candidates[fieldIdx] = append(candidates[fieldIdx], field.name)
			}
		}
		if len(candidates[fieldIdx]) == 1 {
			determined = append(determined, candidates[fieldIdx][0])
		}
	}

	for len(determined) < len(fields) {
		for i := range candidates {
			if len(candidates[i]) == 1 {
				continue
			}
			for _, d := range determined {
				if dIdx := slices.Index(candidates[i], d); dIdx != -1 {
					candidates[i] = slices.Delete(candidates[i], dIdx, dIdx+1)
				}
			}
			if len(candidates[i]) == 1 {
				determined = append(determined, candidates[i][0])
			}
		}
	}

    var result uint = 1
    for i := range candidates {
        name := candidates[i][0]
        if strings.HasPrefix(name, "departure") {
            result *= yourTicket.values[i]
            fmt.Println(i, name)
        }
    }

	return result
}

func main() {
	fields, yourTicket, nearbyTickets := ParseInput()
	fmt.Println("[1]", Part1(fields, nearbyTickets))
	fmt.Println("[2]", Part2(fields, yourTicket, nearbyTickets))
}
