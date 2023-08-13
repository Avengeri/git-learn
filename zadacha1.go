package main

//aaa
import "fmt"

import (
	"math"
	"strings"
)

func calculateAverage(marks [5]float64) float64 {
	sum := float64(0)
	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}
	return sum / float64(len(marks))
}
func main() {
	fmt.Println("Нам даны оценки учеников\n")
	fmt.Println("Оценки Андрея 5,5,5,4,5\n")
	fmt.Println("Оценки Юли 4,4,4,4,5\n")
	fmt.Println("Оценки Полины 1,1,2,2,3\n")
	fmt.Println("Введите exit, чтобы выйти из программы\n")

	var abs string

	Андрей := [5]float64{5, 5, 5, 4, 5}
	Юля := [5]float64{4, 4, 4, 4, 5}
	Полина := [5]float64{1, 1, 2, 2, 3}
outerLoop:
	for {
		var answer string
		fmt.Println("Введите Андрей/Юля/Полина, чтобы подсчитать среднее арифмитическое оценок.")
		fmt.Scan(&answer)
		if answer == "exit" {
			break
		}
		answer = strings.ToLower(answer)
		var result float64
		switch answer {
		case "андрей":
			result = calculateAverage(Андрей)
		case "юля":
			result = calculateAverage(Юля)
		case "полина":
			average := calculateAverage(Полина)
			roundedAverage := int(math.Round(average))
			if roundedAverage == 2 {
				fmt.Printf("%d, а Полина оказывается двоечница\n", roundedAverage)
			} else {
				fmt.Printf("%d, а Полина оказывается хорошим студентом\n", roundedAverage)
			}
		default:
			fmt.Println("Некорректный ввод. Пожалуйста, выберите Андрей, Юля или Полина.")

		}
		fmt.Println(math.Round(result))

		var continueAnswer string
		for {
			fmt.Print("Будем считать еще? да - повторный запуск, нет - выход из программы: ")
			fmt.Scan(&continueAnswer)
			continueAnswer = strings.ToLower(continueAnswer)
			if continueAnswer == "нет" {
				break outerLoop
			} else if continueAnswer == "да" {
				continue outerLoop
			} else {
				fmt.Println("Некорректный ввод. Введите 'да' или 'нет'.")
			}
		}
		fmt.Scan(&abs)
	}
}
