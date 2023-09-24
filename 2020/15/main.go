package main

import "fmt"

func Part1(initial []uint32, rounds uint32) uint32 {
	initialLen := uint32(len(initial))
	var spoken [2020]uint32
	var i uint32
	for ; i < rounds; i++ {
		if i < initialLen {
			spoken[i] = initial[i]
		} else {
			lastUnique := true
			var idx uint32
			for j := i - 1; j > 0; j-- {
				if spoken[j-1] == spoken[i-1] {
					lastUnique = false
					idx = j
					break
				}
			}
			if lastUnique {
				spoken[i] = 0
			} else {
				spoken[i] = i - idx
			}
		}
	}
	return spoken[rounds-1]
}

func Part2(initial []uint32, rounds uint32) uint32 {
	initialLen := uint32(len(initial))
	mem := make(map[uint32]uint32)
	var last, i uint32
	for ; i < rounds; i++ {
		if i < initialLen {
			last = initial[i]
			mem[last] = i
		} else {
			lastIdx, lastNonUnique := mem[last]
			mem[last] = i - 1
			if lastNonUnique {
				last = i - 1 - lastIdx
			} else {
				last = 0
			}
		}
	}
	return last
}

func main() {
	initial := []uint32{9, 19, 1, 6, 0, 5, 4}
	var rounds uint32 = 2020
	fmt.Println("[1]", Part1(initial, rounds))
	rounds = 30000000
	fmt.Println("[2]", Part2(initial, rounds))
}
