package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Seat byte

const (
	SEAT_OUT_OF_BOUNDS Seat = '-'
	SEAT_FLOOR              = '.'
	SEAT_EMPTY              = 'L'
	SEAT_OCCUPIED           = '#'
)

type ToggleSeat struct {
	row, col int
}

type Seats [][]Seat

func (seats Seats) String() string {
	sb := strings.Builder{}
	for _, row := range seats {
		for _, seat := range row {
			sb.WriteByte(byte(seat))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (seats Seats) Get(row, col int) Seat {
	if row < 0 ||
		row >= len(seats) ||
		col < 0 ||
		len(seats) == 0 ||
		col >= len(seats[0]) {
		return SEAT_OUT_OF_BOUNDS
	}
	return seats[row][col]
}

func (seats Seats) OccupiedNeighboursCount(row, col int) int {
	count := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			if seats.Get(row+r, col+c) == SEAT_OCCUPIED {
				count++
			}
		}
	}
	return count
}

func (seats Seats) OccupiedDistantNeighboursCount(row, col int) int {
	count := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			for i := 1; ; i++ {
				rr := row + (r * i)
				cc := col + (c * i)
				switch seats.Get(rr, cc) {
				case SEAT_OCCUPIED:
					count++
					goto out
				case SEAT_EMPTY:
					goto out
				case SEAT_OUT_OF_BOUNDS:
					goto out
				}
			}
		out:
		}
	}
	return count
}

func ReadSeats(fileName string) Seats {
	file, _ := os.Open(fileName)
	defer file.Close()
	var seats Seats
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Seat, len(line))
		for i, b := range line {
			row[i] = Seat(b)
		}
		seats = append(seats, row)
	}
	return seats
}

type ToggleSeatStrategy func(Seats, int, int) bool

func StabilizeSeats(seats Seats, toggleSeatStrategy ToggleSeatStrategy) int {
	iter := func() []ToggleSeat { return seatsIter(seats, toggleSeatStrategy) }
	for ops := iter(); len(ops) > 0; ops = iter() {
		for _, op := range ops {
			seat := &seats[op.row][op.col]
			switch *seat {
			case SEAT_OCCUPIED:
				*seat = SEAT_EMPTY
			case SEAT_EMPTY:
				*seat = SEAT_OCCUPIED
			default:
				panic("Unexpecte seat type!")
			}
		}
	}
	count := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == SEAT_OCCUPIED {
				count++
			}
		}
	}
	return count
}

func seatsIter(seats Seats, toggleSeatStrategy ToggleSeatStrategy) []ToggleSeat {
	var ops []ToggleSeat
	for row := 0; row < len(seats); row++ {
		for col := 0; col < len(seats[0]); col++ {
			if toggleSeatStrategy(seats, row, col) {
				ops = append(ops, ToggleSeat{row, col})
			}
		}
	}
	return ops
}

func Part1ToggleSeatStrategy(seats Seats, row, col int) bool {
	occupied := seats.OccupiedNeighboursCount(row, col)
	if (seats[row][col] == SEAT_EMPTY && occupied == 0) ||
		(seats[row][col] == SEAT_OCCUPIED && occupied >= 4) {
		return true
	}
	return false
}

func Part2ToggleSeatStrategy(seats Seats, row, col int) bool {
	occupied := seats.OccupiedDistantNeighboursCount(row, col)
	if (seats[row][col] == SEAT_EMPTY && occupied == 0) ||
		(seats[row][col] == SEAT_OCCUPIED && occupied >= 5) {
		return true
	}
	return false
}

func main() {
	seats := ReadSeats("input.txt")
	fmt.Println("[1]", StabilizeSeats(seats, Part1ToggleSeatStrategy))
	seats = ReadSeats("input.txt")
	fmt.Println("[2]", StabilizeSeats(seats, Part2ToggleSeatStrategy))
}
