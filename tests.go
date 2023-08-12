package main

import (
	"fmt"
	"strings"
)

func calculaAverage(marks [5]float64) float64 {
	sum := float64(0)
	for _, mark := range marks {
		sum += mark
	}
	return sum / float64(len(marks))
}

func main() {
	fmt.Println("Введите имя ученика:")
	var name string
	fmt.Scan(&name)
	name = strings.TrimSpace(name)

	var marks [5]float64
	fmt.Println("Введите 5 оценок через пробел:")
	for i := 0; i < len(marks); i++ {
		_, err := fmt.Scan(&marks[i])
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
	}

	fmt.Printf("Оценки %s: %v\n", name, marks)

	var result float64
	result = calculaAverage(marks)
	fmt.Printf("Среднее арифметическое оценок %s: %f\n", name, result)

	var continueAnswer string
	for {
		fmt.Print("Будем считать еще? да - повторный запуск, нет - выход из программы: ")
		fmt.Scan(&continueAnswer)
		continueAnswer = strings.ToLower(continueAnswer)
		if continueAnswer == "нет" {
			break
		} else if continueAnswer != "да" {
			fmt.Println("Некорректный ввод. Введите 'да' или 'нет'.")
		}
	}

	fmt.Println("Программа завершена.")
}
