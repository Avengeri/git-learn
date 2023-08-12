package main

import "fmt"

func main() {
	day := "Четверг"
	switch day {
	case "Понедельник":
		fmt.Println("первый день")
	case "Вторник":
		fmt.Println("второй день")
	case "Среда":
		fmt.Println("третий день")
	default:
		fmt.Println("Какой то другой день")
	}
}
