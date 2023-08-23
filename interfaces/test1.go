package main

import "fmt"

type Runner interface {
	Run() string
}
type Swimmer interface {
	Swim() string
}
type Flyer interface {
	Fly() string
}
type Ducker interface {
	Runner
	Swimmer
	Flyer
}
type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Printf("Человек %s бегает", h.Name)
}

func interfaceValues() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("runner is nil")
	} else {
		fmt.Println("runner is not nil")
	}
}
func main() {
	interfaceValues()
}
