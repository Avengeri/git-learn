package main

import (
	"fmt"
	"math"
)

func calculaAverage(marks [5]float64) float64 {
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
	fmt.Println("Оценки Веталя 1,1,2,2,3\n")

	var abs string
	Андрей := [5]float64{5, 5, 5, 4, 5}
	Юля := [5]float64{4, 4, 4, 4, 5}
	Веталь := [5]float64{1, 1, 2, 2, 3}
	for {
		var answer string
		fmt.Println("Введите Андрей/Юля/Веталь, чтобы подсчитать среднее арифмитическое оценок.")
		fmt.Scan(&answer)
		if answer == "exit" {
			break
		}
		var result float64
		switch answer {
		case "Андрей":
			result = calculaAverage(Андрей)
		case "Юля":
			result = calculaAverage(Юля)
		case "Веталь":
			average := calculaAverage(Веталь)
			roundedAverage := int(math.Round(average))
			if roundedAverage == 2 {
				fmt.Printf("%d, а Веталь оказывается двоечник", roundedAverage)
			} else {
				fmt.Printf("%d, а Веталь оказывается хорошим студентом\n", roundedAverage)
			}
		default:
			fmt.Println("Неккоректное значение, введите Андрей, Юля или Веталь")
			continue
		}
		fmt.Println(math.Round(result))
	}
	fmt.Scan(&abs)
}
