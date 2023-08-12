package main

import (
	"fmt"
)

func main() {
	var (
		x1 = [2]string{"hello", "world"}
		x2 = len(x1)
		x3 = [...]int64{1, 2, 3, 4, 5, 9}

		slice1 = make([]int64, 0, 5)
		slice2 = make([]int64, 5, 5)
		slice3 = make([]string, 0, 5)
		slice4 = make([]string, 5, 5)

		ages = map[string]int{
			"Андрей": 29,
			"Юля":    28,
			"Поля":   18,
			"Мама":   49,
			"Котик":  34,
		}
	)

	slice1 = []int64{1, 2, 3, 4}
	slice1 = append(slice1, 7, 9, 14, 24)
	slice2 = []int64{1, 2, 3, 4, 5}
	slice3 = []string{"hello", "bye"}

	ages["Папа"] = 49
	ages["Бабушка"]++

	Andrey := ages["Андрей"]
	Julia := ages["Юля"]
	Polina := ages["Поля"]
	Mama := ages["Мама"]
	Kotik := ages["Котик"]
	Papa := ages["Папа"]

	fmt.Println(x1, x2, x3, slice1, slice2, slice3, slice4, cap(slice1), len(slice1))
	fmt.Println(Andrey, Julia, Polina, Mama, Kotik, ages, Papa)
	fmt.Println(ages["Поля"])

	type point struct {
		X int
		Y string
	}
	p := point{
		X: 5,
		Y: "bye bye",
	}
	fmt.Println(p, p.X, p.Y)

	type point2 struct {
		A int
		B string
		C float32
	}
	p2 := point2{
		A: 32,
		B: "Лошара",
		C: 1.234,
	}
	fmt.Println(p2, p2.A, p2.B, p2.C)
	p2.A = 55
	fmt.Println(p2.A)

}
