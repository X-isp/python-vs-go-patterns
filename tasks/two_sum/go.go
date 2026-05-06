package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if index, ok := seen[diff]; ok {
			return []int{index, i}
		}

		seen[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}
