package main

import "fmt"

type subject struct {
	Name  string
	Grade int
}
type students struct {
	Name    string
	Age     int
	Grades  []subject
	Faculty string
}
type faculty struct {
	Name     string
	Students []students
}

type university struct {
	Name      string
	Faculties []faculty
}

func averageMarks(s []subject) int {
	sum := 0
	for _, subject := range s {
		sum += subject.Grade
	}
	return sum / len(s)
}
func averageFacultyMarks(s []subject) int {
	sum := 0
	for _, subject := range s {
		sum += subject.Grade
	}
	return sum / len(s)
}

func main() {

	subjects := []subject{
		{"Математика", 4},
		{"Информатика", 2},
		{"История", 3},
		{"Русский язык", 5},
	}

	Student1 := students{Name: "Ваня", Age: 21, Grades: subjects}
	Student2 := students{Name: "Петя", Age: 19, Grades: subjects}
	Student3 := students{Name: "Коля", Age: 24, Grades: subjects}

	faculties := faculty{Name: "программирования", Students: []students{Student1, Student2, Student3}}
	university := university{Name: "Лучший университет", Faculties: []faculty{faculties}}
	fmt.Println("Название университета: ", university.Name)
	fmt.Println("Факультеты: ")
	for _, faculty := range university.Faculties {
		fmt.Println(" Факультет", faculty.Name)
		fmt.Println(" Студенты:")
		for _, student := range faculty.Students {
			fmt.Println("    Имя:", student.Name)
			fmt.Println("    Возраст:", student.Age)
			fmt.Println("    Оценки:")
			for _, subject := range student.Grades {
				fmt.Println(" 		Предмет:", subject.Name)
				fmt.Println(" 		Оценка:", subject.Grade)
			}
			fmt.Println(" 	 Средний балл студента", averageMarks(student.Grades))
		}
		fmt.Println("  Средний балл факультета", averageFacultyMarks(faculty.Students[0].Grades))
	}
}
