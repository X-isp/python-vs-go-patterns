package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	data := make([]int, 2000)
	for i := range data {
		data[i] = rand.Intn(10000)
	}

	start := time.Now()
	bubbleSort(data)
	duration := time.Since(start)

	fmt.Printf("Go: %.6f\n", duration.Seconds())
}
