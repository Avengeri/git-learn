package main

import "fmt"

func first() {
	fmt.Println("text first function")
}

func second(a int, b int) {
	fmt.Println(a + b)
}

func third(c string, d string) string {
	str := "first" + c + d
	return str
}

func math_math(q int, w int) (s1 int, s2 int, s3 int, s4 float32) {
	s1 = q + w
	s2 = q - w
	s3 = q + w
	s4 = float32(q) / float32(w)
	return
}

func main() {
	first()
	first()
	second(5, 10)
	stroka := third(" second", " third")
	fmt.Println(stroka)
	s11, s22, s33, s44 := math_math(10, 3)
	fmt.Println(s11, s22, s33, s44)
}
