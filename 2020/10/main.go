package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// naive recursive solution
func CountSolutions(adapters []int) int {
	if len(adapters) == 2 {
		return 1
	}
	sum := 0
	for i := 1; i < len(adapters) && adapters[i]-adapters[0] <= 3; i++ {
		sum += CountSolutions(adapters[i:])
	}
	return sum
}

// naive recursive solution with memoization, still too slow
func CountSolutions2(adapters []int) int {
	mem := make([]int, len(adapters) + 1)
	for i := 0; i < len(mem); i++ {
	    mem[i] = -1
	}
	mem[2] = 1
	mem[1] = 0
	mem[0] = 0
    return countSolutions2(adapters, mem)
}

func countSolutions2(adapters []int, mem []int) int {
	if x := mem[len(adapters)]; x > 0 {
		return x
	}
	sum := 0
	for i := 1; i < len(adapters) && adapters[i]-adapters[0] <= 3; i++ {
		sum += countSolutions2(adapters[i:], mem)
	}
	mem[len(adapters)] = sum
	return sum
}

// smart solution without recursion
func CountSolutions3(adapters []int) int {
	mem := make([]int, len(adapters))
	mem[len(adapters)-1] = 0
	mem[len(adapters)-2] = 1
	for i := len(adapters) - 3; i >= 0; i-- {
		mem[i] = 0
		for j := i; j < len(adapters) && adapters[j]-adapters[i] <= 3; j++ {
			mem[i] += mem[j]
		}
	}
	return mem[0]
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var adapters []int
	adapters = append(adapters, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		adapters = append(adapters, num)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	diff1, diff3 := 0, 0
	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
	}
	fmt.Println("[1]", diff1*diff3)

	// fmt.Println("[2]", count_solutions(adapters))
	fmt.Println("[2] v2", CountSolutions2(adapters))
	fmt.Println("[2] v3", CountSolutions3(adapters))
}
