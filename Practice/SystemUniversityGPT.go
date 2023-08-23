package main

import "fmt"

type Subject struct {
	Name  string
	Grade int
}

type Student struct {
	Name    string
	Age     int
	Grades  []Subject
	Faculty string
}

type Faculty struct {
	Name     string
	Students []Student
}

type University struct {
	Name      string
	Faculties []Faculty
}

func CalculateAverageGrade(subjects []Subject) float64 {
	total := 0
	for _, subject := range subjects {
		total += subject.Grade
	}
	return float64(total) / float64(len(subjects))
}

func CalculateFacultyAverage(subjects []Subject) float64 {
	total := 0
	for _, subject := range subjects {
		total += subject.Grade
	}
	return float64(total) / float64(len(subjects))
}

func main() {
	subjects := []Subject{
		{Name: "Математика", Grade: 4},
		{Name: "Физика", Grade: 5},
		{Name: "Информатика", Grade: 3},
	}

	student := Student{
		Name:    "Иван",
		Age:     20,
		Grades:  subjects,
		Faculty: "Информатики",
	}

	faculty := Faculty{
		Name:     "Информатики",
		Students: []Student{student},
	}

	university := University{
		Name:      "Лучший университет",
		Faculties: []Faculty{faculty},
	}

	fmt.Println("Имя университета:", university.Name)
	fmt.Println("Факультеты:")
	for _, faculty := range university.Faculties {
		fmt.Println("  Факультет:", faculty.Name)
		fmt.Println("  Студенты:")
		for _, student := range faculty.Students {
			fmt.Println("    Имя:", student.Name)
			fmt.Println("    Возраст:", student.Age)
			fmt.Println("    Оценки:")
			for _, subject := range student.Grades {
				fmt.Println("      Предмет:", subject.Name)
				fmt.Println("      Оценка:", subject.Grade)
			}
			fmt.Println("    Средний балл студента:", CalculateAverageGrade(student.Grades))
		}
		fmt.Println("  Средний балл факультета:", CalculateFacultyAverage(faculty.Students[0].Grades))
	}
}
