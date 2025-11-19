package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println("res:", res)

	nums = []int{0, 4, 3, 0}
	target = 0
	res = twoSum(nums, target)
	fmt.Println("res:", res)
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	foundSolutions := map[int]int{}

	for i, n := range nums {
		c := target - n
		missing := target - c
		i2, found := foundSolutions[missing]
		if found {
			return []int{i, i2}
		} else {
			foundSolutions[c] = i
		}
	}
	return []int{}
}
