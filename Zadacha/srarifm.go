package main

import (
	"fmt"
)

func сalculateAverage(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	var name string
	var age int
	fmt.Println("Привет человек! Меня создали для того, чтобы научиться создавать таких, как я и даже круче.")
	fmt.Println("Ты уж сильно на меня не сердись, я могу содержать ошибки, т.к. мой создатель еще только учится")
	fmt.Println("Я могу помочь рассчитать тебе среднее арифмитическое чисел")
	fmt.Println("Но сначала давай познакомимся. Как тебя зовут?")
	fmt.Scan(&name)
	fmt.Printf("Понял, принял %s. Теперь скажи сколько тебе лет?\n", name)
	fmt.Scan(&age)
	if age < 0 {
		fmt.Println("Ты еще не родился")
	} else if age > 0 && age < 18 {
		fmt.Println("Ты довольно молодой, разве ты уже забыл как считать? Но все равно я тебе помогу")
	} else if age > 18 && age < 30 {
		fmt.Println("Не так уж и много тебе лет, у тебя еще все впереди")
	} else if age > 30 && age < 50 {
		fmt.Printf("Сколько, сколько? %d? Ну пизда рулю, копи на гроб!\n", age)
	}
	fmt.Printf("Так ладно %s, давай начнем! Введи числа через пробел, а затем нажми Enter:\n", name)

	//var input string
	//fmt.Scanln(&input)
	//numbersStrings := strings.SplitN(input, " ", -10)
	////fmt.Println(numbers, input)
	//var num []float64
	//for _, nums := range numbersStrings {
	//	StrNum := strconv.ParseFloat()
}
