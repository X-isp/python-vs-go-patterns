package main

import "fmt"

// функция приветствия
func greet(name string) string {
	return "Привет, " + name
}

// функция сложения
func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(greet("Alex"))
	fmt.Println(add(5, 3))
}
