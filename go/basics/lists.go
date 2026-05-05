package main

import "fmt"

func main() {
	// массив (слайс)
	numbers := []int{1, 2, 3, 4, 5}

	// добавление элемента
	numbers = append(numbers, 6)

	// обход
	for _, n := range numbers {
		fmt.Println(n)
	}

	// создание нового слайса (квадраты)
	squares := make([]int, len(numbers))

	for i, n := range numbers {
		squares[i] = n * n
	}

	fmt.Println(squares)
}
