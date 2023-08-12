package main

import (
	"fmt"
)

func main() {
	age := 26
	if age >= 18 {
		fmt.Println("Вы взрослый")
	} else {
		fmt.Println("ты мелкий пиздюк")

	}

	grade := 999
	if grade <= 10 {
		fmt.Println("ты лох ебаный")
	} else if grade >= 10 && grade < 50 {
		fmt.Println("ну такое")
	} else if grade >= 50 && grade < 80 {
		fmt.Println("Пойдет, на пиво хватит")
	} else if grade >= 80 && grade < 100 {
		fmt.Println("Ну неплохо,неплохо")
	} else {
		fmt.Println("Кто ты бля?")
	}
	x := 10
	y := 10
	if x == y {
		fmt.Println("Они равны")
	} else {
		fmt.Println("Они не равны")
	}
	pensia := 7000
	if pensia < 1000 || pensia > 10000 {
		fmt.Println("Че то не адекватное")
	} else {
		fmt.Println("Нормальненько")
	}

}
