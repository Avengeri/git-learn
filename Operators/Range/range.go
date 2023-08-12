package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	for index, value := range slice1 {
		fmt.Println(index, value)
	}

	map1 := map[string]int{
		"Андрей": 29,
		"Юля":    28,
		"Поля":   18,
	}
	for index1, value1 := range map1 {
		fmt.Println(index1, value1)
	}
	text := "Hello"
	for index2, value2 := range text {
		fmt.Println(index2, value2)
	}
	var a string
	fmt.Scan(&a)
}
