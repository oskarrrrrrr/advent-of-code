package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DirectionType byte

const (
	DIR_N DirectionType = 'N'
	DIR_S               = 'S'
	DIR_E               = 'E'
	DIR_W               = 'W'
	DIR_L               = 'L'
	DIR_R               = 'R'
	DIR_F               = 'F'
)

type Direction struct {
	dir   DirectionType
	value int
}

func (dir Direction) String() string {
	return fmt.Sprintf("Direction(dir=%c, value=%v)", dir.dir, dir.value)
}

type Waypoint struct {
	x, y int
}

func (w Waypoint) String() string {
	return fmt.Sprintf("Waypoint(x=%v, y=%v)", w.x, w.y)
}

type Ship struct {
	x, y     int
	waypoint Waypoint
}

func (ship Ship) String() string {
	return fmt.Sprintf("Ship(x=%v, y=%v, waypoint=%v)", ship.x, ship.y, ship.waypoint)
}

func (ship *Ship) Move(move Direction) {
	switch move.dir {
	case DIR_N:
		ship.waypoint.y += move.value
	case DIR_S:
		ship.waypoint.y -= move.value
	case DIR_E:
		ship.waypoint.x -= move.value
	case DIR_W:
		ship.waypoint.x += move.value
	case DIR_F:
        ship.x += move.value * ship.waypoint.x
        ship.y += move.value * ship.waypoint.y
	}
}

func (ship *Ship) Rotate(rotation Direction) {
	switch rotation.dir {
	case DIR_L:
		ship.rotate(rotation)
	case DIR_R:
		ship.rotate(rotation)
	}
}

func (ship *Ship) rotate(rotation Direction) {
    var rot int
    if rotation.dir == DIR_L {
        rot = 360 - rotation.value
    } else {
        rot = rotation.value
    }
    wx, wy := &ship.waypoint.x, &ship.waypoint.y
    switch rot {
    case 90:
        *wx, *wy = -(*wy), *wx
    case 180:
        ship.waypoint.x *= -1
        ship.waypoint.y *= -1
    case 270:
        *wx, *wy = *wy, -(*wx)
    case 360:
    default:
        panic("Unexpected angle!")
    }
}

func ReadDirections() (directions []Direction) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.Atoi(line[1:])
		directions = append(directions, Direction{
			dir:   DirectionType(line[0]),
			value: value,
		})
	}
	return
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	directions := ReadDirections()
	ship := Ship{waypoint: Waypoint{x: -10, y: 1}}
	for _, dir := range directions {
		ship.Move(dir)
		ship.Rotate(dir)
		// fmt.Println(ship)
	}
	result := intAbs(ship.x) + intAbs(ship.y)
    fmt.Println("[1] broken after solving [2] :/")
	fmt.Println("[2]", result)
}
