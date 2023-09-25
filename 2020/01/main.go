package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func twoSum(nums []int, target int) (int, int, bool) {
	l, r := 0, len(nums)-1
	for l < r {
		s := nums[l] + nums[r]
		if s < target {
			l += 1
		} else if s > target {
			r -= 1
		} else {
			return nums[l], nums[r], false
		}
	}
	return 0, 0, true
}

func threeSum(nums []int, target int) (int, int, int, bool) {
	for i := 0; i < len(nums)-2; i++ {
		a, b, err := twoSum(nums[i+1:], target-nums[i])
		if !err {
			return nums[i], a, b, false
		}
	}
	return 0, 0, 0, true
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}
	sort.Ints(nums[:])
	a, b, err := twoSum(nums, 2020)
	if err {
		log.Fatal("Couldn't find two sum!")
	}
	fmt.Printf("[two sum] a: %d, b: %d, result: %d\n", a, b, a*b)
	a, b, c, err := threeSum(nums, 2020)
	if err {
		log.Fatal("Couldn't find three sum!")
	}
	fmt.Printf("[three sum] a: %d, b: %d, c: %d, result: %d\n", a, b, c, a*b*c)
}
