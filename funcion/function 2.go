package main

import "fmt"

func test(somefunct func(int) int) {
	fmt.Println(somefunct(5))
}
func test2(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	a := func() {
		fmt.Println("Hello")
	}
	b := func(q int, w int) (q1 float64, q2 float64) {
		q1 = float64(q) * float64(w)
		q2 = float64(q) / float64(w)
		return
	}
	a()
	q11, q22 := b(19, 28)
	fmt.Println(q11, q22)
	square := func(x int) int {
		return x * x
	}
	test(square)
	test2("Hello маза фака")()
}
