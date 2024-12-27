package main

import "core:fmt"
import "core:os"
import "core:bytes"

Pos :: [2]int
Empty :: struct{}

Direction :: enum { UP, RIGHT, DOWN, LEFT }
DirectionVector := [Direction]Pos {
    .UP = { -1, 0 }, .RIGHT = { 0, 1 }, .DOWN = { 1, 0 }, .LEFT = { 0, -1 }
}
DirectionRotateClockwise := [Direction]Direction {
    .UP = .RIGHT, .RIGHT = .DOWN, .DOWN = .LEFT, .LEFT = .UP
}

Cell :: enum u8 { EMPTY, OBSTACLE }
CellChar := [Cell]u8 { .EMPTY = '.', .OBSTACLE = '#' }

Map :: struct {
    w: int,
    h: int,
    start: Pos,
    cells: [dynamic]Cell,
}

map_at :: proc(_map: Map, pos: Pos) -> Cell {
    return _map.cells[_map.w * pos[0] + pos[1]]
}

map_set :: proc(_map: Map, pos: Pos, new_value: Cell) {
    _map.cells[_map.w * pos[0] + pos[1]] = new_value
}

map_parse :: proc(data: []u8) -> Map {
    data := bytes.trim_right_space(data)
    _map: Map
    col := 0
    for ch in data {
        switch ch {
        case '.':
            append(&_map.cells, Cell.EMPTY)
            col += 1
        case '#':
            append(&_map.cells, Cell.OBSTACLE)
            col += 1
        case '^':
            append(&_map.cells, Cell.EMPTY)
            _map.start[0] = _map.h
            _map.start[1] = col
        case '\n':
            _map.h += 1
            if _map.w == 0 do _map.w = col
            col = 0
        }
    }
    _map.h += 1
    return _map
}

map_print :: proc(_map: Map) {
    for row := 0; row < _map.h; row += 1 {
        for col := 0; col < _map.w; col += 1 {
            fmt.print(rune(CellChar[map_at(_map, {row, col})]))
        }
        fmt.println()
    }
}

pos_is_valid :: proc(_map: Map, pos: Pos) -> bool {
    return 0 <= pos[0] && pos[0] < _map.h &&
           0 <= pos[1] && pos[1] < _map.w
}

pos_move :: proc(pos: Pos, dir: Direction) -> Pos {
    dir_vec := DirectionVector[dir]
    return { pos[0] + dir_vec[0], pos[1] + dir_vec[1] }
}

part1 :: proc(_map: Map) -> int {
    pos := _map.start
    dir := Direction.UP
    visited := map[Pos]Empty{}
    defer delete(visited)
    for {
        visited[pos] = Empty{}
        next_pos := pos_move(pos, dir)
        pos_is_valid(_map, next_pos) or_break
        switch map_at(_map, next_pos) {
        case .OBSTACLE:
            dir = DirectionRotateClockwise[dir]
        case .EMPTY:
            pos = next_pos
        }
    }
    return len(visited)
}

part2 :: proc(_map: Map) -> int {

    Entry :: struct { pos: Pos, dir: Direction }
    visited := map[Entry]Empty{}
    defer delete(visited)
    has_loop :: proc(visited: ^map[Entry]Empty, _map: Map) -> bool {
        clear(visited)
        pos := _map.start
        dir := Direction.UP
        for {
            entry := Entry{ pos = pos, dir = dir }
            if entry in visited do return true
            visited[entry] = Empty{}
            next_pos := pos_move(pos, dir)
            pos_is_valid(_map, next_pos) or_break
            switch map_at(_map, next_pos) {
            case .OBSTACLE:
                dir = DirectionRotateClockwise[dir]
            case .EMPTY:
                pos = next_pos
            }
        }
        return false
    }

    pos := _map.start
    dir := Direction.UP
    obstacles := map[Pos]Empty{}
    defer delete(obstacles)
    for {
        next_pos := pos_move(pos, dir)
        pos_is_valid(_map, next_pos) or_break
        switch map_at(_map, next_pos) {
        case .OBSTACLE:
            dir = DirectionRotateClockwise[dir]
        case .EMPTY:
            map_set(_map, next_pos, .OBSTACLE)
            if has_loop(&visited, _map) do obstacles[next_pos] = Empty{}
            map_set(_map, next_pos, .EMPTY)
            pos = next_pos
        }
    }
    return len(obstacles)
}

main :: proc() {
    data, err := os.read_entire_file_or_err("input.txt")
    assert(len(data) > 0, "Failed to read input.")
    _map := map_parse(data)
    fmt.println("part1: ", part1(_map))
    fmt.println("part2: ", part2(_map))
}
