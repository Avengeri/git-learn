package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Выберите программу для запуска:")
	fmt.Println("1. Подсчет среднего арифмитического оценок")
	fmt.Println("2. Поиск min и max в массиве")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	switch choice {
	case 1:
		runZadacha1()
	case 2:
		runZadacha2()
	default:
		fmt.Println("Некорректный выбор.")
	}
}

func runZadacha1() {
	cmd := exec.Command("./zadacha1.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске zadacha1:", err)
	}
}

func runZadacha2() {
	cmd := exec.Command("./zadacha2.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске zadacha2:", err)
	}
}
