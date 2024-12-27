package main

import "core:fmt"
import "core:os"
import "core:bytes"

Antena :: struct { name: rune, x: int, y: int}

part1 :: proc(antenas: [128][dynamic]Antena, w: int, h: int) -> int {
    check_freq :: proc(antenas: []Antena, x: int, y: int) -> bool {
        for antena, antena_idx in antenas {
            vec: [2]int = { x - antena.x, y - antena.y }
            for next_antena in antenas[antena_idx+1:] {
                next_vec: [2]int = { x- next_antena.x, y - next_antena.y }
                if (vec[0] == 2*next_vec[0] && vec[1] == 2*next_vec[1]) ||
                   (2*vec[0] == next_vec[0] && 2*vec[1] == next_vec[1]) {
                    return true
                }
            }
        }
        return false
    }
    count := 0
    for y := 0; y < h; y += 1 {
        for x := 0; x < w; x += 1 {
            for ants in antenas {
                if check_freq(ants[:], x, y) {
                    count += 1
                    break
                }
            }
        }
    }
    return count
}

part2 :: proc(antenas: [128][dynamic]Antena, w: int, h: int) -> int {
    check_freq :: proc(antenas: []Antena, x: int, y: int) -> bool {
        for antena, antena_idx in antenas {
            vec: [2]int = { x - antena.x, y - antena.y }
            for next_antena in antenas[antena_idx+1:] {
                next_vec: [2]int = { x- next_antena.x, y - next_antena.y }
                if (vec[0] * next_vec[1] == vec[1] * next_vec[0]) do return true
            }
        }
        return false
    }
    count := 0
    for y := 0; y < h; y += 1 {
        for x := 0; x < w; x += 1 {
            for ants in antenas {
                if check_freq(ants[:], x, y) {
                    count += 1
                    break
                }
            }
        }
    }
    return count
}

main :: proc() {
    data, ok := os.read_entire_file_from_filename("input.txt")
    defer delete(data)
    assert(ok, "Failed to read input.txt")

    antenas := [128][dynamic]Antena{}
    defer {
        for arr in antenas do if arr != nil do delete(arr)
    }

    w, h: int
    for line in bytes.split_iterator(&data, {'\n'}) {
        w = 0
        for ch in line {
            if ch != '.' {
                if antenas[ch] == nil do antenas[ch] = make([dynamic]Antena)
                append(&antenas[ch], Antena{ name = rune(ch), x = w, y = h })
            }
            w += 1
        }
        h += 1
    }

    fmt.println("part1: ", part1(antenas, w, h))
    fmt.println("part2: ", part2(antenas, w, h))
}
