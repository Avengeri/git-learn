package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	sum := 0
	a := 1
	for a <= 10 {
		sum += a
		a++
	}
	fmt.Println("Сумма", sum)
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
	t := 0
	for {
		fmt.Println(t)
		t++
		if t > 10 {
			break
		}
	}
}
