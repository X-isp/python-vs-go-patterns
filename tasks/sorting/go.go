package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func main() {
	data := []int{5, 2, 9, 1, 5, 6}

	fmt.Println("Before:", data)
	fmt.Println("After:", bubbleSort(data))
}
