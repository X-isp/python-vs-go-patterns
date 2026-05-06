package main

import "fmt"

func reverseArray(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(reverseArray(nums))
}
