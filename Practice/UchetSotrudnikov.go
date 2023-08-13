package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int8
	Salary    float64
}

func main() {
	Employee1 := Employee{"Ваня", "Ваничкин", 35, 40000}
	Employee2 := Employee{"Петя", "Петичкин", 25, 38000}
	Employee3 := Employee{"Катя", "Катина", 28, 14788}
	Employee4 := Employee{"Саня", "Саничкин", 59, 105444}

	EmployeePtr := &Employee1
	EmployeePtr.Salary = 99999

	fmt.Println("Информация о сотрудниках")
	PrintEmployee(Employee1)
	PrintEmployee(Employee2)
	PrintEmployee(Employee3)
	PrintEmployee(Employee4)
}

func PrintEmployee(e Employee) {
	fmt.Printf("Имя: %s %s\n", e.FirstName, e.LastName)
	fmt.Printf("Возраст: %d \n", e.Age)
	fmt.Printf("Зарплата: %.0f \n", e.Salary)
	fmt.Println(".............")
}
