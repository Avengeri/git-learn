package main

import "fmt"

func main() {
	var answer string
	var a string
	fmt.Print("Ты жирная свинья? (да/нет): ")
	fmt.Scan(&answer)
	if answer == "да" {
		fmt.Println("Хватит жрать!")
	} else if answer == "нет" {
		fmt.Println("Не пизди мне тут")
	} else {
		fmt.Println("Повыебываться решил? Сказано было введи да или нет!")
	}
	fmt.Scan(&a)
}
