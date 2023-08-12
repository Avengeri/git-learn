package main

import (
	"fmt"
)

func MinMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	var abc string

	Андрей := []int{5, 10, 3, 8, 2, 7}
	Юля := []int{-2, 100, 6, 4, 2, 22}
	Полина := []int{333, -555, 15, 73, 4, 2458}

	fmt.Println("Андрей 5, 10, 3, 8, 2, 7\n")
	fmt.Println("Юля -2, 100, 6, 4, 2, 22\n")
	fmt.Println("Полина 333, -555, 15, 73, 4, 2458\n")
	fmt.Println("Введите имя для подсчета минимального и максимального значения (Андрей/Юля/Полина):")
	var answer string
	fmt.Scan(&answer)

	var numbers []int
	switch answer {
	case "Андрей":
		numbers = Андрей
	case "Юля":
		numbers = Юля
	case "Полина":
		numbers = Полина
	default:
		fmt.Println("Некорректное имя. Введите 'Андрей', 'Юля' или 'Полина'.")
		return

	}
	min, max := MinMax(numbers)
	fmt.Printf("Минимальное значение: %d\n", min)
	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Scan(&abc)

}
