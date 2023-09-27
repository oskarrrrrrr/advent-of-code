package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func TwoSum(nums []int, target int) (int, int, bool) {
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

func ThreeSum(nums []int, target int) (int, int, int, bool) {
	for i := 0; i < len(nums)-2; i++ {
		a, b, err := TwoSum(nums[i+1:], target-nums[i])
		if !err {
			return nums[i], a, b, false
		}
	}
	return 0, 0, 0, true
}

func ReadInput() []int {
	inputBytes, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(inputBytes))
	var nums []int
	for _, numStr := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(numStr)
		nums = append(nums, n)
	}
	return nums
}

func main() {
	nums := ReadInput()
	sort.Ints(nums[:])
	a, b, _ := TwoSum(nums, 2020)
	fmt.Println("[1]", a*b)
	a, b, c, _ := ThreeSum(nums, 2020)
	fmt.Println("[2]", a*b*c)
}
