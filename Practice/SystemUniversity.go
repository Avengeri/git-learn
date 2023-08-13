package main

import "fmt"

type Subject struct {
	Object string
	Mark   int
}
type Student struct {
	Name  string
	Age   int
	Marks []int
	Faculty
}
type Faculty struct {
	title    string
	Students []string
}

func AverageMarks(Mark []int) int {
	sum := 0
	for i := 0; i < len(Mark); i++ {
		sum += Mark[i]
	}
	return sum / len(Mark)
}
func AverageMarksFaculty(Mark []int) int {
	sum := 0
	for i := 0; i < len(Mark); i++ {
		sum += Mark[i]
	}
	return sum / len(Mark)
}
func main() {
	Student1 := Student{Name: "Ваня", Age: 21, Marks: []int{3, 3, 4, 5, 1, 2}}
	Student2 := Student{Name: "Петя", Age: 19, Marks: []int{5, 5, 3, 4, 4, 1}}
	Student3 := Student{Name: "Коля", Age: 24, Marks: []int{5, 5, 4, 3, 5, 1}}
	fmt.Println(Student1, Student2, Student3)
}
