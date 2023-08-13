package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	fmt.Println("Введите числа через пробел:")
	var input string
	fmt.Scanln(&input)

	// Разделяем введенную строку на числа
	inputNumbers := strings.Fields(input)
	var numbers []float64

	// Преобразуем строки в числа
	for _, strNum := range inputNumbers {
		num, err := strconv.ParseFloat(strNum, 64)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число:", err)
			return
		}
		numbers = append(numbers, num)
	}

	// Вызываем функцию CalculateAverage и получаем результат
	average := CalculateAverage(numbers)

	fmt.Printf("Среднее арифметическое: %.2f\n", average)
}
