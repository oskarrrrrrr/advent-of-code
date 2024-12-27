package main

import "core:os"
import "core:fmt"
import "core:bytes"
import "core:strconv"

part1 :: proc(target: u64, nums: []u64) -> bool {
    _part1 :: proc(target: u64, nums: []u64, curr: u64) -> bool {
        if len(nums) == 0 do return target == curr
        return _part1(target, nums[1:], curr + nums[0]) ||
               _part1(target, nums[1:], curr * nums[0])
    }
    return _part1(target, nums[1:], nums[0])
}

part2 :: proc(target: u64, nums: []u64) -> bool {
    concat :: proc(n1: u64, n2: u64) -> u64 {
        digits: [64]u8
        n2 := n2
        idx := 0
        for n2 > 0 {
            digits[idx] = u8(n2 % 10)
            n2 /= 10
            idx += 1
        }
        idx -= 1
        res := n1
        for idx >= 0 {
            res = (res * 10) + u64(digits[idx])
            idx -= 1
        }
        return res
    }

    _part1 :: proc(target: u64, nums: []u64, curr: u64) -> bool {
        if len(nums) == 0 do return target == curr
        return _part1(target, nums[1:], curr + nums[0]) ||
               _part1(target, nums[1:], curr * nums[0]) ||
               _part1(target, nums[1:], concat(curr, nums[0]))
    }
    return _part1(target, nums[1:], nums[0])
}

main :: proc() {
    data, err := os.read_entire_file_from_filename_or_err("input.txt")
    defer delete(data)
    assert(err == nil, os.error_string(err))
    p1, p2: u64
    nums: [dynamic]u64
    for line in bytes.split_iterator(&data, {'\n'}) {
        parts := bytes.split(line, {':'})
        target, _ := strconv.parse_u64(string(parts[0]))
        nums_str := bytes.trim_space(parts[1])
        clear(&nums)
        for num_str in bytes.split_iterator(&nums_str, {' '}) {
            new_num, _ := strconv.parse_u64(string(num_str))
            append(&nums, new_num)
        }
        if part1(target, nums[:]) do p1 += target
        if part2(target, nums[:]) do p2 += target
    }
    fmt.println("part1: ", p1)
    fmt.println("part2: ", p2)
}
