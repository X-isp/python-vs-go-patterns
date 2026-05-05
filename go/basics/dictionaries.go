package main

import "fmt"

func main() {
	user := map[string]interface{}{
		"name":   "Alex",
		"age":    25,
		"active": true,
	}

	fmt.Println(user["name"])

	// добавление значения
	user["city"] = "Vilnius"

	fmt.Println(user)
}
