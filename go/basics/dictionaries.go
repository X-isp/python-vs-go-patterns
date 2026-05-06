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
	user["city"] = "Saint Petersburg"

	fmt.Println(user)
}
