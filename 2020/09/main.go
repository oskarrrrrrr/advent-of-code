package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 25

func get_part2_result(window []int) int {
	min := window[0]
	max := window[0]
	for _, x := range window {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min + max
}

func part2(input []int, sum int) int {
	for wlen := 2; wlen <= 20; wlen++ {
		var window []int
		var s int
		for i := 0; i < wlen; i++ {
			window = append(window, input[i])
			s += input[i]
		}
		if s == sum {
			return get_part2_result(window)
		}
		for i := wlen; i < len(input); i++ {
			j := i % wlen
			s -= window[j]
			s += input[i]
			window[j] = input[i]
			if s == sum {
				return get_part2_result(window)
			}
		}
	}
	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	values := make(map[int]int)
	var window [N]int
	var input []int
	var part1_result int
	i := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
		if values[window[i]] != 0 {
			values[window[i]]--
			any_ok := false
			for _, x := range window {
				if values[num-x] != 0 {
					any_ok = true
				}
			}
			if !any_ok {
				fmt.Println("[1] Failed at", num)
				part1_result = num
				break
			}
		}
		window[i] = num
		i = (i + 1) % N
		values[num]++
	}

	fmt.Println("[2]", part2(input, part1_result))
}
